package storage

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	storageAsyncApi "gitlab.ozon.dev/Vanek623/task-manager-system/external/api/async"
	storageSyncApi "gitlab.ozon.dev/Vanek623/task-manager-system/external/api/sync"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/cache"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/task/models"

	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/storage/config"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/counters"

	storagePkg "gitlab.ozon.dev/Vanek623/task-manager-system/external/task/storage"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/storage"
	"google.golang.org/grpc"
)

type iTaskStorage interface {
	Add(ctx context.Context, t *models.Task) error
	Delete(ctx context.Context, ID *uuid.UUID) error
	List(ctx context.Context, limit, offset uint64) ([]*models.Task, error)
	Update(ctx context.Context, t *models.Task) error
	Get(ctx context.Context, ID *uuid.UUID) (*models.Task, error)
}

type Server struct {
	wg  *sync.WaitGroup
	cs  *counters.Counters
	ctx context.Context

	s iTaskStorage
}

func NewServer(ctx context.Context) (*Server, error) {
	s := &Server{
		wg:  &sync.WaitGroup{},
		cs:  counters.New("storage_service"),
		ctx: ctx,
	}

	if err := s.runDB(); err != nil {
		return nil, err
	}

	return s, nil
}

// Run запуск сервиса хранилища
func (s *Server) Run() {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.RunGRPC()
	}()

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.RunKafka()
	}()

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.RunHTTP()
	}()

	s.wg.Wait()
}

// RunGRPC запуск GRPC
func (s *Server) RunGRPC() {
	ctx, cl := context.WithCancel(s.ctx)
	defer cl()

	cfg := config.GetConfigGRPC()
	listener, err := net.Listen(cfg.ConnectionType, cfg.Host)
	if err != nil {
		log.Error(err)
	}

	server := grpc.NewServer()

	pb.RegisterStorageServer(server, storageSyncApi.NewProtobufAPI(s.s, s.cs))

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		<-ctx.Done()
		server.Stop()
		log.Info("Kafka down")
	}()

	if err := server.Serve(listener); err != nil {
		log.Error(err)
	}
}

// RunKafka запуск Kafka
func (s *Server) RunKafka() {
	cfg := config.GetConfigKafka()

	saramaCfg := sarama.NewConfig()
	client, err := sarama.NewConsumerGroup(cfg.Brokers, cfg.Group, saramaCfg)
	if err != nil {
		log.Error(err)
		return
	}

	redisCfg := config.GetRedisConfig()
	cw, err := cache.NewRedisWriter(s.ctx, &redisCfg)
	if err != nil {
		log.Error(err)
		return
	}

	handler := storageAsyncApi.NewKafkaAPI(s.s, s.cs, cw)

	log.WithFields(log.Fields{
		"brokers": cfg.Brokers,
		"topics":  cfg.Topics,
	}).Info("Kafka run")

	for {
		select {
		case <-s.ctx.Done():
			log.Info("Stop kafka")
			if err := client.Close(); err != nil {
				log.Error(err)
			}
			return
		default:
		}

		if err := client.Consume(s.ctx, cfg.Topics, handler); err != nil {
			log.WithField("error", err).Error("Consuming")
			time.Sleep(time.Second * 5)
		}
	}
}

func (s *Server) connectToDB() (*pgxpool.Pool, error) {
	cfg := config.GetConfigDB()
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s "+
		"dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.UserName, cfg.Password, cfg.Name)

	pool, err := pgxpool.Connect(s.ctx, psqlConn)
	if err != nil {
		return nil, errors.Wrap(err, "can't connect to database")
	}

	if err = pool.Ping(s.ctx); err != nil {
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

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		<-s.ctx.Done()
		pool.Close()
		log.Info("DB connection down")
	}()

	return pool, nil
}

// runDB поднятие БД
func (s *Server) runDB() error {
	pool, err := s.connectToDB()
	if err != nil {
		return err
	}

	pg := storagePkg.NewPostgres(pool, s.cs, false)
	s.s, err = storagePkg.NewMemcached(pg, s.cs, config.MemcachedHost)
	if err != nil {
		return err
	}

	return nil
}

// RunHTTP запуск HTTP сервера (счетчики)
func (s *Server) RunHTTP() {
	serv := http.Server{
		Addr:              config.HTTPHost,
		ReadHeaderTimeout: time.Second,
	}

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		<-s.ctx.Done()

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
