package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service"

	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/bot"

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
	TasksList(ctx context.Context, data models.ListTaskData) ([]models.TaskBrief, error)
	UpdateTask(ctx context.Context, data models.UpdateTaskData) error
	GetTask(ctx context.Context, data models.GetTaskData) (*models.DetailedTask, error)
}

// Run запуск GRPC, REST and Tg Bot
func Run() {
	var wg sync.WaitGroup
	wg.Add(3)

	ctx, cl := context.WithCancel(context.Background())
	defer cl()

	s, err := service.New()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer wg.Done()
		RunGRPC(s)
	}()

	go func() {
		defer wg.Done()
		RunREST(ctx)
	}()

	go func() {
		defer wg.Done()
		bot.Run(s)
	}()

	wg.Wait()
}

// RunGRPC запускает GRPC
func RunGRPC(service iService) {
	listener, err := net.Listen(connectionType, addressGRPC)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterServiceServer(s, serviceApiPkg.NewAPI(service))

	log.Println("grpc started")

	if err = s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

// @title       TaskBrief manager API
// @version     1.0
// @description API Server for TaskBrief manager application

// @host     localhost:8080
// @BasePath /

// RunREST запускает REST
func RunREST(ctx context.Context) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterServiceHandlerFromEndpoint(ctx, mux, addressGRPC, opts); err != nil {
		log.Fatal(err)
	}

	log.Println("rest started")

	if err := http.ListenAndServe(addressHTTP, mux); err != nil {
		log.Fatal(err)
	}
}
