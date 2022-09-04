package service

import (
	"context"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/bot"
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/service/config"
	serviceApiPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/api"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/counters"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/client"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/tracer"
	"go.opentelemetry.io/otel"

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

type Server struct {
	cs  *counters.Counters
	wg  *sync.WaitGroup
	ctx context.Context

	s iService
}

func NewServer(ctx context.Context) (*Server, error) {
	srv := &Server{
		cs:  counters.New("task_service"),
		wg:  &sync.WaitGroup{},
		ctx: ctx,
	}

	if err := srv.makeService(ctx, false); err != nil {
		return nil, err
	}

	return srv, nil
}

// Run запуск Kafka&GRPC, REST and Tg Bot
func (s *Server) Run() {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.RunHTTP()
		log.Info("HTTP down")
	}()

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		pbClient, err := client.New(config.GetServiceGRPCConfig().Host, 1)
		if err != nil {
			log.Error(err)
			return
		}

		bot.Run(s.ctx, pbClient, s.cs)
		log.Info("Bot down")
	}()

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.RunGRPC()
		log.Info("GRPC down")
	}()

	s.wg.Wait()

	log.Info("Server down")
}

func (s *Server) makeService(ctx context.Context, isStorageSync bool) error {
	if isStorageSync {
		grpcCfg := config.GetStorageGRPCConfig()
		syncStorage, err := serviceSyncStorage.NewGRPC(ctx, grpcCfg.Host, s.cs)
		if err != nil {
			return err
		}

		log.WithField("host", grpcCfg.Host).Debug("Connected to storage over GRPC")
		s.s = servicePkg.NewServiceWithSyncStorage(syncStorage)
		return nil
	}

	kafkaCfg := config.GetKafkaConfig()
	asyncStorageWriter, err := serviceAsyncStorage.NewKafkaWriter(ctx, kafkaCfg.Brokers, s.cs)
	if err != nil {
		return err
	}

	log.WithField("brokers", kafkaCfg.Brokers).Debug("Connected to storage over Kafka")

	redisCfg := config.GetRedisConfig()
	asyncStorageReader, err := serviceAsyncStorage.NewRedisReader(ctx, &redisCfg)
	if err != nil {
		return err
	}

	log.Debug("Connected to redis")

	s.s = servicePkg.NewServiceWithAsyncStorage(asyncStorageWriter, asyncStorageReader)

	return nil
}

func (s *Server) MonitoringInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	s.cs.Inc(counters.Incoming)

	newCtx, span := otel.Tracer(tracer.Name).Start(ctx, tracer.MakeSpanName(info.FullMethod))
	defer span.End()

	res, err := handler(newCtx, req)
	if err != nil {
		s.cs.Inc(counters.Fail)
		return nil, err
	}

	s.cs.Inc(counters.Success)
	return res, nil
}

// RunGRPC запускает GRPC
func (s *Server) RunGRPC() {
	cfg := config.GetServiceGRPCConfig()
	listener, err := net.Listen(cfg.ConnectionType, cfg.Host)
	if err != nil {
		log.Error(err)
		return
	}
	grpcS := grpc.NewServer(grpc.UnaryInterceptor(s.MonitoringInterceptor))
	pb.RegisterServiceServer(grpcS, serviceApiPkg.NewAPI(s.s, s.cs))

	log.WithField("host", cfg.Host).Info("GRPC server up")

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		<-s.ctx.Done()
		grpcS.Stop()
	}()

	if err = grpcS.Serve(listener); err != nil {
		if errors.Is(err, grpc.ErrServerStopped) {
			log.Info(err)
		} else {
			log.Error(err)
		}
	}
}

// RunHTTP запускает REST и swagger
func (s *Server) RunHTTP() {
	gwmux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	cfg := config.GetServiceGRPCConfig()
	if err := pb.RegisterServiceHandlerFromEndpoint(s.ctx, gwmux, cfg.Host, opts); err != nil {
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

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		<-s.ctx.Done()
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
