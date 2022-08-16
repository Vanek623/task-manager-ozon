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
	AddTask(ctx context.Context, data models.AddTaskData) (uint64, error)
	DeleteTask(ctx context.Context, data models.DeleteTaskData) error
	TasksList(ctx context.Context, data models.ListTaskData) ([]*models.Task, error)
	UpdateTask(ctx context.Context, data models.UpdateTaskData) error
	GetTask(ctx context.Context, data models.GetTaskData) (*models.DetailedTask, error)
}

type iTaskStorage interface {
	Add(ctx context.Context, t *models.Task) (uint64, error)
	Delete(ctx context.Context, ID uint64) error
	List(ctx context.Context, limit, offset uint64) ([]*models.Task, error)
	Update(ctx context.Context, t *models.Task) error
	Get(ctx context.Context, ID uint64) (*models.Task, error)
}

// Run запуск GRPC, REST and Tg Bot
func Run(storage iTaskStorage) {
	var wg sync.WaitGroup
	wg.Add(3)

	ctx, cl := context.WithCancel(context.Background())
	defer cl()

	s, err := service.New(storage)
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
