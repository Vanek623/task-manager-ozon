package storage

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/storage/config"
	storageApiPkg "gitlab.ozon.dev/Vanek623/task-manager-system/external/api"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/counters"

	storagePkg "gitlab.ozon.dev/Vanek623/task-manager-system/external/task/storage"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/storage"
	"google.golang.org/grpc"
)

// Run запуск сервиса хранилища
func Run(ctx context.Context) error {
	ctx, cl := context.WithCancel(ctx)
	defer cl()

	cs := counters.New("storage_service")

	var wg sync.WaitGroup
	storage, err := RunDB(ctx, &wg, cs)
	if err != nil {
		return err
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		RunGRPC(ctx, &wg, storage, cs)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		RunKafka(ctx, storage, cs)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		RunHTTP(ctx, &wg)
	}()

	wg.Wait()

	return nil
}

// RunGRPC запуск GRPC
func RunGRPC(ctx context.Context, wg *sync.WaitGroup, s *storagePkg.Storage, cs *counters.Counters) {
	ctx, cl := context.WithCancel(ctx)
	defer cl()

	cfg := config.GetConfigGRPC()
	listener, err := net.Listen(cfg.ConnectionType, cfg.Host)
	if err != nil {
		log.Error(err)
	}

	server := grpc.NewServer()

	pb.RegisterStorageServer(server, storageApiPkg.NewProtobufAPI(s, cs))

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		server.Stop()
		log.Info("Kafka down")
	}()

	if err := server.Serve(listener); err != nil {
		log.Error(err)
	}
}

// RunKafka запуск Kafka
func RunKafka(ctx context.Context, s *storagePkg.Storage, cs *counters.Counters) {
	cfg := config.GetConfigKafka()

	saramaCfg := sarama.NewConfig()
	client, err := sarama.NewConsumerGroup(cfg.Brokers, cfg.Group, saramaCfg)
	if err != nil {
		log.Error(err)
	}

	handler := storageApiPkg.NewKafkaAPI(s, cs)

	log.WithFields(log.Fields{
		"brokers": cfg.Brokers,
		"topics":  cfg.Topics,
	}).Info("Kafka run")

	for {
		select {
		case <-ctx.Done():
			log.Info("Stop kafka")
			if err := client.Close(); err != nil {
				log.Error(err)
			}
			return
		default:
		}

		if err := client.Consume(ctx, cfg.Topics, handler); err != nil {
			log.WithField("error", err).Error("Consuming")
			time.Sleep(time.Second * 5)
		}
	}
}

func connectToDB(ctx context.Context, wg *sync.WaitGroup) (*pgxpool.Pool, error) {
	cfg := config.GetConfigDB()
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s "+
		"dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.UserName, cfg.Password, cfg.Name)

	pool, err := pgxpool.Connect(ctx, psqlConn)
	if err != nil {
		return nil, errors.Wrap(err, "can't connect to database")
	}

	if err = pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	poolConfig := pool.Config()
	poolConfig.MaxConnIdleTime = cfg.MaxConnIdleTime
	poolConfig.MaxConnLifetime = cfg.MaxConnLifetime
	poolConfig.MinConns = cfg.MinConnections
	poolConfig.MaxConns = cfg.MaxConnections

	log.WithFields(log.Fields{
		"host": cfg.Host,
		"port": cfg.Port,
	}).Info("Connected to storage")

	wg.Add(1)
	go func() {
		<-ctx.Done()
		pool.Close()
		log.Info("DB connection down")
		wg.Done()
	}()

	return pool, nil
}

// RunDB поднятие БД
func RunDB(ctx context.Context, wg *sync.WaitGroup, cs *counters.Counters) (*storagePkg.Storage, error) {
	pool, err := connectToDB(ctx, wg)
	if err != nil {
		return nil, err
	}

	pg := storagePkg.NewPostgres(pool, cs, false)
	storage, err := storagePkg.NewMemcached(pg, cs, config.MemcachedHost)
	if err != nil {
		return nil, err
	}

	return storage, nil
}

// RunHTTP запуск HTTP сервера (счетчики)
func RunHTTP(ctx context.Context, wg *sync.WaitGroup) {
	serv := http.Server{
		Addr:              config.HTTPHost,
		ReadHeaderTimeout: time.Second,
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		if err := serv.Shutdown(context.Background()); err != nil {
			log.Error(err)
		}
	}()

	if err := serv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Error(err)
		}
	}
}
