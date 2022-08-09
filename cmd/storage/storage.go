package storage

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	storageApiPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/api/storage"
	storagePkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/storage"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/storage"
	"google.golang.org/grpc"
)

func RunStorage(ctx context.Context, password string) {
	listener, err := net.Listen(connectionType, serviceAddress)
	if err != nil {
		log.Fatal(err)
	}

	pool, err := connectToDB(ctx, password)
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	storage := storagePkg.NewPostgres(pool, true)
	pb.RegisterStorageServer(server, storageApiPkg.NewApi(storage))

	log.Println("grpc started")

	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func connectToDB(ctx context.Context, password string) (*pgxpool.Pool, error) {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s "+
		"dbname=%s sslmode=disable", hostDB, portDB, userName, password, nameDB)

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
