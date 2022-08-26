package storage

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/counters"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
	storageModelsPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/storage/models"

	"github.com/pkg/errors"

	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/storage"
	grpcPkg "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	reconnectMaxCount = 5
	reconnectTimeout  = time.Second
	grpcCountersName  = "storage_protobuf"
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

	s.cs.Inc(counters.Outbound)
	_, err := s.client.TaskAdd(ctx, request)
	if err != nil {
		return err
	}

	return nil
}

func (s *grpc) Delete(ctx context.Context, data *models.DeleteTaskData) error {
	s.cs.Inc(counters.Outbound)
	_, err := s.client.TaskDelete(ctx, &pb.TaskDeleteRequest{ID: uuidToBytes(data.ID())})
	if err != nil {
		return err
	}

	return nil
}

func (s *grpc) List(ctx context.Context, data *models.ListTaskData) ([]*models.Task, error) {
	s.cs.Inc(counters.Outbound)
	resp, err := s.client.TaskList(ctx, &pb.TaskListRequest{
		Limit:  data.Limit(),
		Offset: data.Offset(),
	})

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

	s.cs.Inc(counters.Outbound)
	_, err := s.client.TaskUpdate(ctx, request)
	if err != nil {
		return err
	}

	return nil
}

func (s *grpc) Get(ctx context.Context, data *models.GetTaskData) (*models.DetailedTask, error) {
	s.cs.Inc(counters.Outbound)
	resp, err := s.client.TaskGet(ctx, &pb.TaskGetRequest{ID: uuidToBytes(data.ID())})
	if err != nil {
		return nil, err
	}

	return models.NewDetailedTask(resp.GetTitle(), resp.GetDescription(), time.Unix(resp.GetEdited(), 0)), nil
}

func newGRPC(address string) (*grpc, error) {
	time.Sleep(reconnectTimeout)

	con, err := grpcPkg.Dial(address, grpcPkg.WithTransportCredentials(insecure.NewCredentials()))
	for count := 1; err != nil || con == nil; count++ {
		if count > reconnectMaxCount {
			return nil, errors.Wrap(err, "service_storage: cannot connect to storage server")
		}

		log.Printf("cannot connect to server, try to connect #%d of %d in %d\n", count, reconnectMaxCount, reconnectTimeout)
		time.Sleep(reconnectTimeout)
		con, err = grpcPkg.Dial(address, grpcPkg.WithTransportCredentials(insecure.NewCredentials()))
	}

	return &grpc{
		client: pb.NewStorageClient(con),
		cs:     counters.New(grpcCountersName),
	}, nil
}

func uuidToBytes(ID *uuid.UUID) []byte {
	bytes := make([]byte, 16)
	copy(bytes, ID[:])

	return bytes
}
