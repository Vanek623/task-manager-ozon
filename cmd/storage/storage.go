package storage

import (
	"context"
	"fmt"
	"log"
	"net"

	storagePkg "gitlab.ozon.dev/Vanek623/task-manager-system/external/task/storage"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	storageApiPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/api/storage"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/storage"
	"google.golang.org/grpc"
)

// Run запуск сервиса хранилища
func Run(conf *ConfigDB) {
	listener, err := net.Listen(connectionType, serviceAddress)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cl := context.WithCancel(context.Background())
	defer cl()

	pool, err := connectToDB(ctx, conf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("connected to storage with address %s:%s", conf.Host, conf.Port)

	server := grpc.NewServer()
	storage := storagePkg.NewPostgres(pool, true)
	pb.RegisterStorageServer(server, storageApiPkg.NewAPI(storage))

	log.Printf("GRPC up with address %s", serviceAddress)

	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func connectToDB(ctx context.Context, conf *ConfigDB) (*pgxpool.Pool, error) {
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s "+
		"dbname=%s sslmode=disable", conf.Host, conf.Port, conf.UserName, conf.Password, conf.Name)

	pool, err := pgxpool.Connect(ctx, psqlConn)
	if err != nil {
		return nil, errors.Wrap(err, "can't connect to database")
	}

	if err = pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	config := pool.Config()
	config.MaxConnIdleTime = maxConnIdleTime
	config.MaxConnLifetime = maxConnLifetime
	config.MinConns = minConnections
	config.MaxConns = maxConnections

	return pool, nil
}
