package server

import (
	"context"
	"net"
	"net/http"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	serviceApiPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/api"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/counters"

	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/bot"

	"github.com/google/uuid"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service"
	serviceStorage "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/storage"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/service"

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

// Run запуск Kafka&GRPC, REST and Tg Bot
func Run() error {
	ctx, cl := context.WithCancel(context.Background())
	defer cl()

	cs := counters.New("task_service")
	syncStorage, err := serviceStorage.NewGRPC(ctx, storageAddress, cs)
	if err != nil {
		return err
	}

	asyncStorage, err := serviceStorage.NewKafka(ctx, brokers, syncStorage, cs)
	if err != nil {
		return err
	}

	s := service.New(asyncStorage)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		RunHTTP(ctx)
	}()

	wg.Add(1)
	go func() {
		bot.Run(ctx, s, cs)
	}()

	wg.Add(1)
	go func() {
		RunGRPC(ctx, s, cs)
	}()

	wg.Wait()

	return nil
}

// RunGRPC запускает GRPC
func RunGRPC(ctx context.Context, service iService, cs *counters.Counters) {
	ctx, cl := context.WithCancel(ctx)
	defer cl()

	listener, err := net.Listen(connectionType, addressGRPC)
	if err != nil {
		log.Error(err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterServiceServer(s, serviceApiPkg.NewAPI(service, cs))

	log.Printf("GRPC up with address %s", addressGRPC)

	go func() {
		<-ctx.Done()

		s.Stop()
		log.Info("GRPC server down")
	}()

	if err = s.Serve(listener); err != nil {
		log.Error(err)
	}
}

// RunHTTP запускает REST и swagger
func RunHTTP(ctx context.Context) {
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterServiceHandlerFromEndpoint(ctx, gwmux, addressGRPC, opts); err != nil {
		log.Error(err)
		return
	}

	log.Infof("REST up with address %s", addressHTTP)

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	fs := http.FileServer(http.Dir(swaggerDir))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	log.Infof("Swagger up with address %s", addressHTTP)

	serv := http.Server{
		Addr:              addressHTTP,
		Handler:           mux,
		ReadHeaderTimeout: time.Second,
	}
	go func() {
		<-ctx.Done()
		if err := serv.Shutdown(context.Background()); err != nil {
			log.Error(err)
		} else {
			log.Info("HTTP server down")
		}
	}()

	if err := serv.ListenAndServe(); err != nil {
		log.Error(err)
	}
}
