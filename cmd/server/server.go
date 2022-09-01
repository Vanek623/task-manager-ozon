package server

import (
	"context"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/bot"
	serviceApiPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/api"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/counters"

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
func Run(ctx context.Context) error {
	ctx, cl := context.WithCancel(ctx)
	defer cl()

	cs := counters.New("task_service")
	syncStorage, err := serviceStorage.NewGRPC(ctx, storageAddress, cs)
	if err != nil {
		return err
	}

	log.WithField("host", storageAddress).Debug("Connected to storage over GRPC")

	asyncStorage, err := serviceStorage.NewKafka(ctx, brokers, syncStorage, cs)
	if err != nil {
		return err
	}

	log.WithField("brokers", brokers).Debug("Connected to storage over Kafka")

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
	listener, err := net.Listen(connectionType, addressGRPC)
	if err != nil {
		log.Error(err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterServiceServer(s, serviceApiPkg.NewAPI(service, cs))

	log.WithField("host", addressGRPC).Info("GRPC server up")

	done := make(chan struct{})
	go func() {
		<-ctx.Done()

		s.Stop()
		done <- struct{}{}
	}()

	if err = s.Serve(listener); err != nil {
		if errors.Is(err, grpc.ErrServerStopped) {
			log.Info(err)
		} else {
			log.Error(err)
		}
	}

	<-done
}

// RunHTTP запускает REST и swagger
func RunHTTP(ctx context.Context) {
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterServiceHandlerFromEndpoint(ctx, gwmux, addressGRPC, opts); err != nil {
		log.Error(err)
		return
	}

	log.WithField("host", addressHTTP).Info("REST up")

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	fs := http.FileServer(http.Dir(swaggerDir))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	log.WithField("host", addressHTTP).Info("Swagger up")

	serv := http.Server{
		Addr:              addressHTTP,
		Handler:           mux,
		ReadHeaderTimeout: time.Second,
	}

	done := make(chan struct{})
	go func() {
		<-ctx.Done()
		if err := serv.Shutdown(context.Background()); err != nil {
			log.Error(err)
		}
		done <- struct{}{}
	}()

	if err := serv.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Info(err)
		} else {
			log.Error(err)
		}
	}

	<-done
}
