package storage

import (
	"context"
	"log"
	"time"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"

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
}

func (s *grpc) Add(ctx context.Context, data *models.AddTaskData) (uint64, error) {
	request := &pb.TaskAddRequest{Title: data.Title()}

	if tmp := data.Description(); tmp != "" {
		request.Description = &tmp
	}

	resp, err := s.client.TaskAdd(ctx, request)
	if err != nil {
		return 0, err
	}

	return resp.GetID(), nil
}

func (s *grpc) Delete(ctx context.Context, data *models.DeleteTaskData) error {
	_, err := s.client.TaskDelete(ctx, &pb.TaskDeleteRequest{ID: data.ID()})

	return err
}

func (s *grpc) List(ctx context.Context, data *models.ListTaskData) ([]*models.Task, error) {
	resp, err := s.client.TaskList(ctx, &pb.TaskListRequest{
		Limit:  data.Limit(),
		Offset: data.Offset(),
	})

	if err != nil {
		return nil, err
	}

	res := make([]*models.Task, len(resp.GetTasks()))
	for i, task := range resp.GetTasks() {
		res[i] = models.NewTask(task.GetID(), task.GetTitle())
	}

	return res, nil
}

func (s *grpc) Update(ctx context.Context, data *models.UpdateTaskData) error {
	request := &pb.TaskUpdateRequest{
		ID:    data.ID(),
		Title: data.Title(),
	}

	if tmp := data.Description(); tmp != "" {
		request.Description = &tmp
	}

	_, err := s.client.TaskUpdate(ctx, request)
	if err != nil {
		return err
	}

	return nil
}

func (s *grpc) Get(ctx context.Context, data *models.GetTaskData) (*models.DetailedTask, error) {
	resp, err := s.client.TaskGet(ctx, &pb.TaskGetRequest{ID: data.ID()})
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
	}, nil
}
