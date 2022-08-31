package sync

import (
	"context"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/counters"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
	storageModelsPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/storage/models"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/tracer"
	"go.opentelemetry.io/otel"

	"github.com/pkg/errors"

	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/storage"
	grpcPkg "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	reconnectMaxCount = 5
	reconnectTimeout  = time.Second
)

type grpc struct {
	client pb.StorageClient
	cs     *counters.Counters
}

func (s *grpc) Add(ctx context.Context, data *storageModelsPkg.AddTaskData) error {
	id := data.ID()
	request := &pb.TaskAddRequest{
		ID:    uuidToBytes(&id),
		Title: data.Title(),
	}

	if tmp := data.Description(); tmp != "" {
		request.Description = &tmp
	}

	log.WithFields(log.Fields{
		"id":          request.GetID(),
		"title":       request.GetTitle(),
		"description": request.GetDescription(),
	}).Info("Outbound add request")

	_, span := otel.Tracer(tracer.Name).Start(ctx, tracer.MakeSpanName("Add GRPC"))
	defer span.End()

	s.cs.Inc(counters.Outbound)
	_, err := s.client.TaskAdd(ctx, request)
	if err != nil {
		return err
	}

	return nil
}

func (s *grpc) Delete(ctx context.Context, data *models.DeleteTaskData) error {
	request := &pb.TaskDeleteRequest{ID: uuidToBytes(data.ID())}

	log.WithFields(log.Fields{
		"id": request.GetID(),
	}).Info("Outbound delete request")

	_, span := otel.Tracer(tracer.Name).Start(ctx, tracer.MakeSpanName("Delete GRPC"))
	defer span.End()

	s.cs.Inc(counters.Outbound)
	_, err := s.client.TaskDelete(ctx, request)
	if err != nil {
		return err
	}

	return nil
}

func (s *grpc) List(ctx context.Context, data *models.ListTaskData) ([]*models.Task, error) {
	request := &pb.TaskListRequest{
		Limit:  data.Limit(),
		Offset: data.Offset(),
	}

	log.WithFields(log.Fields{
		"offset": request.GetOffset(),
		"limit":  request.GetLimit(),
	}).Info("Outbound list request")

	s.cs.Inc(counters.Outbound)
	resp, err := s.client.TaskList(ctx, request)

	if err != nil {
		return nil, err
	}

	res := make([]*models.Task, len(resp.GetTasks()))
	for i, task := range resp.GetTasks() {
		id, err := uuid.FromBytes(task.GetID())
		if err != nil {
			return nil, err
		}
		res[i] = models.NewTask(&id, task.GetTitle())
	}

	return res, nil
}

func (s *grpc) Update(ctx context.Context, data *models.UpdateTaskData) error {
	request := &pb.TaskUpdateRequest{
		ID:    uuidToBytes(data.ID()),
		Title: data.Title(),
	}

	if tmp := data.Description(); tmp != "" {
		request.Description = &tmp
	}

	log.WithFields(log.Fields{
		"id":          request.GetID(),
		"title":       request.GetTitle(),
		"description": request.GetDescription(),
	}).Info("Outbound update request")

	s.cs.Inc(counters.Outbound)
	_, err := s.client.TaskUpdate(ctx, request)
	if err != nil {
		return err
	}

	return nil
}

func (s *grpc) Get(ctx context.Context, data *models.GetTaskData) (*models.DetailedTask, error) {

	request := &pb.TaskGetRequest{ID: uuidToBytes(data.ID())}

	log.WithFields(log.Fields{
		"id": request.GetID(),
	}).Info("Outbound get request")

	s.cs.Inc(counters.Outbound)
	resp, err := s.client.TaskGet(ctx, request)
	if err != nil {
		return nil, err
	}

	return models.NewDetailedTask(resp.GetTitle(), resp.GetDescription(), time.Unix(resp.GetEdited(), 0)), nil
}

func newGRPC(ctx context.Context, address string, cs *counters.Counters) (*grpc, error) {
	time.Sleep(reconnectTimeout)

	con, err := grpcPkg.Dial(address, grpcPkg.WithTransportCredentials(insecure.NewCredentials()))
	for count := 1; err != nil || con == nil; count++ {
		if count > reconnectMaxCount {
			return nil, errors.Wrap(err, "service_storage: cannot connect to storage")
		}

		log.Errorf("cannot connect to storage, try to connect #%d of %d in %.1f s", count, reconnectMaxCount, reconnectTimeout.Seconds())
		time.Sleep(reconnectTimeout)
		con, err = grpcPkg.Dial(address, grpcPkg.WithTransportCredentials(insecure.NewCredentials()))
	}

	go func() {
		<-ctx.Done()
		if err := con.Close(); err != nil {
			log.Error(err)
		} else {
			log.Info("GRPC connection closed")
		}
	}()

	return &grpc{
		client: pb.NewStorageClient(con),
		cs:     cs,
	}, nil
}

func uuidToBytes(ID *uuid.UUID) []byte {
	bytes := make([]byte, 16)
	copy(bytes, ID[:])

	return bytes
}
