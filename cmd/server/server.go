package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	serviceStorage "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/storage"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/service"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service"

	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/bot"

	serviceApiPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/api/service"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

type iService interface {
	AddTask(ctx context.Context, data *models.AddTaskData) (uint64, error)
	DeleteTask(ctx context.Context, data *models.DeleteTaskData) error
	TasksList(ctx context.Context, data *models.ListTaskData) ([]*models.Task, error)
	UpdateTask(ctx context.Context, data *models.UpdateTaskData) error
	GetTask(ctx context.Context, data *models.GetTaskData) (*models.DetailedTask, error)
}

// Run запуск GRPC, REST and Tg Bot
func Run() {
	var wg sync.WaitGroup
	wg.Add(3)

	ctx, cl := context.WithCancel(context.Background())
	defer cl()

	storage, err := serviceStorage.NewGRPC(storageAddress)
	if err != nil {
		log.Fatal(err)
	}

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

func RunREST(ctx context.Context) {
	ctx, cl := context.WithCancel(ctx)
	defer cl()

	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterServiceHandlerFromEndpoint(ctx, gwmux, addressGRPC, opts); err != nil {
		log.Fatal(err)
	}

	log.Println("rest started")

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	fs := http.FileServer(http.Dir(swaggerDir))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	log.Println("swagger started")

	if err := http.ListenAndServe(addressHTTP, mux); err != nil {
		log.Fatal(err)
	}
}
