package storage

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/storage"
)

type iTaskStorage interface {
	Add(ctx context.Context, t models.Task) (uint, error)
	Delete(ctx context.Context, ID uint) error
	List(ctx context.Context, limit, offset uint) ([]models.Task, error)
	Update(ctx context.Context, t models.Task) error
	Get(ctx context.Context, ID uint) (*models.Task, error)
}

type implementation struct {
	pb.UnimplementedStorageServer
	s iTaskStorage
}

func NewApi(s iTaskStorage) pb.StorageServer {
	return &implementation{s: s}
}

func codeTask(in *models.Task) *pb.Task {
	task := pb.Task{
		ID:      uint64(in.ID),
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
		ID:          uint(in.GetID()),
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

	ID, err := i.s.Add(ctx, *decoded)
	if err != nil {
		return nil, err
	}

	return &pb.TaskAddResponse{ID: uint64(ID)}, nil
}

func (i *implementation) TaskList(ctx context.Context, in *pb.TaskListRequest) (*pb.TaskListResponse, error) {
	tasks, err := i.s.List(ctx, uint(in.GetLimit()), uint(in.GetOffset()))
	if err != nil {
		return nil, err
	}

	result := make([]*pb.Task, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, codeTask(&task))
	}

	return &pb.TaskListResponse{Tasks: result}, nil
}

func (i *implementation) TaskUpdate(ctx context.Context, in *pb.TaskUpdateRequest) (*pb.TaskUpdateResponse, error) {
	task, err := decodeTask(in.GetTask())
	if err != nil {
		return nil, err
	}

	if err := i.s.Update(ctx, *task); err != nil {
		return nil, err
	}

	return &pb.TaskUpdateResponse{}, nil
}

func (i *implementation) TaskDelete(ctx context.Context, in *pb.TaskDeleteRequest) (*pb.TaskDeleteResponse, error) {
	if err := i.s.Delete(ctx, uint(in.GetID())); err != nil {
		return nil, err
	}

	return &pb.TaskDeleteResponse{}, nil
}

func (i *implementation) TaskGet(ctx context.Context, in *pb.TaskGetRequest) (*pb.TaskGetResponse, error) {
	task, err := i.s.Get(ctx, uint(in.ID))
	if err != nil {
		return nil, err
	}

	return &pb.TaskGetResponse{Task: codeTask(task)}, nil
}
