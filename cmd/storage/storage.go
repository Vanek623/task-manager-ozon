package storage

import (
	"context"
	"fmt"
	"net"
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
func Run(ctx context.Context) {
	ctx, cl := context.WithCancel(ctx)
	defer cl()

	pool, err := connectToDB(ctx)
	defer func() {
		pool.Close()
	}()

	if err != nil {
		log.Error(err)
	}

	cs := counters.New("storage_service")
	storage := storagePkg.NewPostgres(pool, cs)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		RunGRPC(ctx, storage, cs)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		RunKafka(ctx, storage, cs)
		wg.Done()
	}()

	wg.Wait()
}

// RunGRPC запуск GRPC
func RunGRPC(ctx context.Context, s *storagePkg.Storage, cs *counters.Counters) {
	ctx, cl := context.WithCancel(ctx)
	defer cl()

	cfg := config.GetConfigGRPC()
	listener, err := net.Listen(cfg.ConnectionType, cfg.Host)
	if err != nil {
		log.Error(err)
	}

	server := grpc.NewServer()

	pb.RegisterStorageServer(server, storageApiPkg.NewProtobufAPI(s, cs))

	go func() {
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

func connectToDB(ctx context.Context) (*pgxpool.Pool, error) {
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

	return pool, nil
}
