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
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/server/config"
	serviceApiPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/api"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/counters"

	"github.com/google/uuid"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"

	servicePkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service"

	serviceAsyncStorage "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/storage/async"
	serviceSyncStorage "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/storage/sync"
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

	s, err := makeService(ctx, cs, true)
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		RunHTTP(ctx, &wg)
		log.Info("HTTP down")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		bot.Run(ctx, s, cs)
		log.Info("Bot down")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		RunGRPC(ctx, &wg, s, cs)
		log.Info("GRPC down")
	}()

	wg.Wait()

	log.Info("Server down")

	return nil
}

func makeService(ctx context.Context, cs *counters.Counters, isStorageSync bool) (iService, error) {
	if isStorageSync {
		grpcCfg := config.GetStorageGRPCConfig()
		syncStorage, err := serviceSyncStorage.NewGRPC(ctx, grpcCfg.Host, cs)
		if err != nil {
			return nil, err
		}

		log.WithField("host", grpcCfg.Host).Debug("Connected to storage over GRPC")

		return servicePkg.NewServiceWithSyncStorage(syncStorage), nil
	}

	kafkaCfg := config.GetKafkaConfig()
	asyncStorageWriter, err := serviceAsyncStorage.NewKafkaWriter(ctx, kafkaCfg.Brokers, cs)
	if err != nil {
		return nil, err
	}

	log.WithField("brokers", kafkaCfg.Brokers).Debug("Connected to storage over Kafka")

	redisCfg := config.GetRedisConfig()
	asyncStorageReader, err := serviceAsyncStorage.NewRedisReader(ctx, &redisCfg)
	if err != nil {
		return nil, err
	}

	log.Debug("Connected to redis")

	return servicePkg.NewServiceWithAsyncStorage(asyncStorageWriter, asyncStorageReader), nil
}

// RunGRPC запускает GRPC
func RunGRPC(ctx context.Context, wg *sync.WaitGroup, service iService, cs *counters.Counters) {
	cfg := config.GetServiceGRPCConfig()
	listener, err := net.Listen(cfg.ConnectionType, cfg.Host)
	if err != nil {
		log.Error(err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterServiceServer(s, serviceApiPkg.NewAPI(service, cs))

	log.WithField("host", cfg.Host).Info("GRPC server up")

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		s.Stop()
	}()

	if err = s.Serve(listener); err != nil {
		if errors.Is(err, grpc.ErrServerStopped) {
			log.Info(err)
		} else {
			log.Error(err)
		}
	}
}

// RunHTTP запускает REST и swagger
func RunHTTP(ctx context.Context, wg *sync.WaitGroup) {
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	cfg := config.GetServiceGRPCConfig()
	if err := pb.RegisterServiceHandlerFromEndpoint(ctx, gwmux, cfg.Host, opts); err != nil {
		log.Error(err)
		return
	}

	log.WithField("host", config.HTTPAddress).Info("REST up")

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	fs := http.FileServer(http.Dir(config.SwaggerDir))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	log.WithField("host", config.HTTPAddress).Info("Swagger up")

	serv := http.Server{
		Addr:              config.HTTPAddress,
		Handler:           mux,
		ReadHeaderTimeout: time.Second,
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		if err := serv.Shutdown(context.Background()); err != nil {
			log.Error(err)
		}
	}()

	if err := serv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Error(err)
		}
	}
}
