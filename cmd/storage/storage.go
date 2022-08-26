package storage

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/storage/config"
	storageApiPkg "gitlab.ozon.dev/Vanek623/task-manager-system/external/api"

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
		log.Println(err)
	}

	storage := storagePkg.NewPostgres(pool, true)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		RunGRPC(ctx, storage)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		RunKafka(ctx, storage)
		wg.Done()
	}()

	wg.Wait()
}

// RunGRPC запуск GRPC
func RunGRPC(ctx context.Context, s *storagePkg.Storage) {
	cfg := config.GetConfigGRPC()
	listener, err := net.Listen(cfg.ConnectionType, cfg.Host)
	if err != nil {
		log.Println(err)
	}

	server := grpc.NewServer()

	pb.RegisterStorageServer(server, storageApiPkg.NewProtobufAPI(s))

	ch := make(chan error)
	go func() {
		err := server.Serve(listener)
		ch <- err
	}()

	select {
	case err := <-ch:
		if err != nil {
			log.Println(err)
		}
		server.Stop()
	case <-ctx.Done():
	}

	server.Stop()
	log.Println("Kafka stop")
}

// RunKafka запуск Kafka
func RunKafka(ctx context.Context, s *storagePkg.Storage) {
	cfg := config.GetConfigKafka()

	saramaCfg := sarama.NewConfig()
	client, err := sarama.NewConsumerGroup(cfg.Brokers, cfg.Group, saramaCfg)
	if err != nil {
		log.Println(err)
	}

	handler := storageApiPkg.NewKafkaAPI(s)

	log.Printf("Kafka run, working with brokers %v and topics %v", cfg.Brokers, cfg.Topics)

	for {
		select {
		case <-ctx.Done():
			log.Println("Stop kafka")
		default:
		}

		if err := client.Consume(ctx, cfg.Topics, handler); err != nil {
			log.Printf("on consume: %v", err)
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

	log.Printf("connected to storage with address %s:%s", cfg.Host, cfg.Port)

	return pool, nil
}
