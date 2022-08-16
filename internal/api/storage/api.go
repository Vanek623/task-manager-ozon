package storage

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/storage"
)

type iTaskStorage interface {
	Add(ctx context.Context, t *models.Task) (uint64, error)
	Delete(ctx context.Context, ID uint64) error
	List(ctx context.Context, limit, offset uint64) ([]*models.Task, error)
	Update(ctx context.Context, t *models.Task) error
	Get(ctx context.Context, ID uint64) (*models.Task, error)
}

type implementation struct {
	pb.UnimplementedStorageServer
	s iTaskStorage
}

// NewAPI создание обработчика хранилища
func NewAPI(s iTaskStorage) pb.StorageServer {
	return &implementation{s: s}
}

func encodeTask(in *models.Task) *pb.Task {
	task := pb.Task{
		ID:      in.ID,
		Title:   in.Title,
		Created: in.Created.Unix(),
		Updated: in.Edited.Unix(),
	}

	if in.Description != "" {
		task.Description = &in.Description
	}

	return &task
}

func decodeTask(in *pb.Task) (*models.Task, error) {
	if in == nil {
		return nil, errors.New("task_decoding: empty data")
	}

	return &models.Task{
		ID:          in.GetID(),
		Title:       in.GetTitle(),
		Description: in.GetDescription(),
		Created:     time.Unix(in.GetCreated(), 0),
		Edited:      time.Unix(in.GetUpdated(), 0),
	}, nil
}

func (i *implementation) TaskAdd(ctx context.Context, in *pb.TaskAddRequest) (*pb.TaskAddResponse, error) {
	decoded, err := decodeTask(in.GetTask())
	if err != nil {
		return nil, err
	}

	id, err := i.s.Add(ctx, decoded)
	if err != nil {
		return nil, err
	}

	return &pb.TaskAddResponse{ID: id}, nil
}

func (i *implementation) TaskList(ctx context.Context, in *pb.TaskListRequest) (*pb.TaskListResponse, error) {
	tasks, err := i.s.List(ctx, in.GetLimit(), in.GetOffset())
	if err != nil {
		return nil, err
	}

	result := make([]*pb.Task, len(tasks))
	for i, t := range tasks {
		result[i] = encodeTask(t)
	}

	return &pb.TaskListResponse{Tasks: result}, nil
}

func (i *implementation) TaskUpdate(ctx context.Context, in *pb.TaskUpdateRequest) (*pb.TaskUpdateResponse, error) {
	decoded, err := decodeTask(in.GetTask())
	if err != nil {
		return nil, err
	}

	if err := i.s.Update(ctx, decoded); err != nil {
		return nil, err
	}

	return &pb.TaskUpdateResponse{}, nil
}

func (i *implementation) TaskDelete(ctx context.Context, in *pb.TaskDeleteRequest) (*pb.TaskDeleteResponse, error) {
	if err := i.s.Delete(ctx, in.GetID()); err != nil {
		return nil, err
	}

	return &pb.TaskDeleteResponse{}, nil
}

func (i *implementation) TaskGet(ctx context.Context, in *pb.TaskGetRequest) (*pb.TaskGetResponse, error) {
	task, err := i.s.Get(ctx, in.ID)
	if err != nil {
		return nil, err
	}

	return &pb.TaskGetResponse{Task: encodeTask(task)}, nil
}
