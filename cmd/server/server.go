package server

import (
	"context"
	"log"
	"net"
	"net/http"

	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/bot"

	"github.com/google/uuid"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service"
	serviceStorage "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/storage"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/service"

	serviceApiPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/api/service"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

type iService interface {
	AddTask(ctx context.Context, data *models.AddTaskData) (*uuid.UUID, error)
	DeleteTask(ctx context.Context, data *models.DeleteTaskData) error
	TasksList(ctx context.Context, data *models.ListTaskData) ([]*models.Task, error)
	UpdateTask(ctx context.Context, data *models.UpdateTaskData) error
	GetTask(ctx context.Context, data *models.GetTaskData) (*models.DetailedTask, error)
}

// Run запуск GRPC, REST and Tg Bot
func Run() error {
	ctx, cl := context.WithCancel(context.Background())
	defer cl()

	storage, err := serviceStorage.NewGRPC(storageAddress)
	if err != nil {
		log.Fatal(err)
	}

	s := service.New(storage)

	go RunREST(ctx)
	go bot.Run(ctx, s)

	return RunGRPC(s)
}

// RunGRPC запускает GRPC
func RunGRPC(service iService) error {
	listener, err := net.Listen(connectionType, addressGRPC)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterServiceServer(s, serviceApiPkg.NewAPI(service))

	log.Printf("GRPC up with address %s", addressGRPC)

	if err = s.Serve(listener); err != nil {
		return err
	}

	return nil
}

// RunREST запускает REST
func RunREST(ctx context.Context) {
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterServiceHandlerFromEndpoint(ctx, gwmux, addressGRPC, opts); err != nil {
		log.Fatal(err)
	}

	log.Printf("REST up with address %s", addressHTTP)

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	fs := http.FileServer(http.Dir(swaggerDir))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	log.Printf("swagger up with address %s", addressHTTP)

	ch := make(chan struct{})
	go func() {
		if err := http.ListenAndServe(addressHTTP, mux); err != nil {
			log.Println(err)
		}
		ch <- struct{}{}
	}()

	select {
	case <-ctx.Done():
	case <-ch:
	}
}
