package server

import (
	"context"
	"log"
	"net"
	"net/http"

	serviceApiPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/api/service"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials/insecure"

	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/service"
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
	pb.RegisterServiceServer(s, serviceApiPkg.NewApi(service))

	log.Println("grpc started")

	if err = s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

// RunREST запускает REST
func RunREST(ctx context.Context) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterServiceHandlerFromEndpoint(ctx, mux, FullAddress, opts); err != nil {
		log.Fatal(err)
	}

	log.Println("rest started")

	if err := http.ListenAndServe(FullHTTPAddress, mux); err != nil {
		log.Fatal(err)
	}
}
