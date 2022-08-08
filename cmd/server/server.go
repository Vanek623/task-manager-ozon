package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v4/pgxpool"

	"google.golang.org/grpc/credentials/insecure"

	apiPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/api"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api"
	"google.golang.org/grpc"
)

type iService interface {
	AddTask(ctx context.Context, data models.AddTaskData) (uint, error)
	DeleteTask(ctx context.Context, data models.DeleteTaskData) error
	TasksList(ctx context.Context, data models.ListTaskData) ([]models.Task, error)
	UpdateTask(ctx context.Context, data models.UpdateTaskData) error
	GetTask(ctx context.Context, data models.GetTaskData) (*models.DetailedTask, error)
}

// RunGRPC запускает GRPC
func RunGRPC(service iService) {
	listener, err := net.Listen(ConnectionType, FullAddress)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterAdminServer(s, apiPkg.New(service))

	log.Println("grpc started")

	if err = s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

// RunREST запускает REST
func RunREST(ctx context.Context) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterAdminHandlerFromEndpoint(ctx, mux, FullAddress, opts); err != nil {
		log.Fatal(err)
	}

	log.Println("rest started")

	if err := http.ListenAndServe(FullHTTPAddress, mux); err != nil {
		log.Fatal(err)
	}
}

// ConnectToDB подключение к БД
func ConnectToDB(ctx context.Context, password string) *pgxpool.Pool {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s "+
		"dbname=%s sslmode=disable", hostDB, portDB, userName, password, nameDB)

	pool, err := pgxpool.Connect(ctx, psqlConn)
	if err != nil {
		log.Fatal("can't connect to database", err)
	}

	if err = pool.Ping(ctx); err != nil {
		pool.Close()
		log.Fatal(err)
	}

	config := pool.Config()
	config.MaxConnIdleTime = maxConnIdleTime
	config.MaxConnLifetime = maxConnLifetime
	config.MinConns = minConnections
	config.MaxConns = maxConnections

	return pool
}
