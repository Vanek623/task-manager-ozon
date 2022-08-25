package api

import (
	"context"

	"github.com/google/uuid"

	"gitlab.ozon.dev/Vanek623/task-manager-system/external/task/models"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/storage"
)

type iTaskStorage interface {
	Add(ctx context.Context, t *models.Task) error
	Delete(ctx context.Context, ID *uuid.UUID) error
	List(ctx context.Context, limit, offset uint64) ([]*models.Task, error)
	Update(ctx context.Context, t *models.Task) error
	Get(ctx context.Context, ID *uuid.UUID) (*models.Task, error)
}

type implementation struct {
	pb.UnimplementedStorageServer
	s iTaskStorage
}

// NewAPI создание обработчика хранилища
func NewAPI(s iTaskStorage) pb.StorageServer {
	return &implementation{s: s}
}

func (i *implementation) TaskAdd(ctx context.Context, in *pb.TaskAddRequest) (*pb.TaskAddResponse, error) {
	task := models.Task{
		Title:       in.GetTitle(),
		Description: in.GetDescription(),
	}

	var err error
	task.ID, err = uuid.FromBytes(in.GetID())
	if err != nil {
		return nil, err
	}

	if err := i.s.Add(ctx, &task); err != nil {
		return nil, err
	}

	return &pb.TaskAddResponse{}, nil
}

func (i *implementation) TaskList(ctx context.Context, in *pb.TaskListRequest) (*pb.TaskListResponse, error) {
	tasks, err := i.s.List(ctx, in.GetLimit(), in.GetOffset())
	if err != nil {
		return nil, err
	}

	result := make([]*pb.TaskListResponse_Task, len(tasks))
	for i, t := range tasks {
		result[i] = &pb.TaskListResponse_Task{
			ID:    uuidToBytes(&t.ID),
			Title: t.Title,
		}
	}

	return &pb.TaskListResponse{Tasks: result}, nil
}

func (i *implementation) TaskUpdate(ctx context.Context, in *pb.TaskUpdateRequest) (*pb.TaskUpdateResponse, error) {
	task := models.Task{
		Title:       in.GetTitle(),
		Description: in.GetDescription(),
	}

	var err error
	task.ID, err = uuid.FromBytes(in.GetID())
	if err != nil {
		return nil, err
	}

	if err = i.s.Update(ctx, &task); err != nil {
		return nil, err
	}

	return &pb.TaskUpdateResponse{}, nil
}

func (i *implementation) TaskDelete(ctx context.Context, in *pb.TaskDeleteRequest) (*pb.TaskDeleteResponse, error) {
	id, err := uuid.FromBytes(in.GetID())
	if err != nil {
		return nil, err
	}

	if err := i.s.Delete(ctx, &id); err != nil {
		return nil, err
	}

	return &pb.TaskDeleteResponse{}, nil
}

func (i *implementation) TaskGet(ctx context.Context, in *pb.TaskGetRequest) (*pb.TaskGetResponse, error) {
	id, err := uuid.FromBytes(in.GetID())
	if err != nil {
		return nil, err
	}

	task, err := i.s.Get(ctx, &id)
	if err != nil {
		return nil, err
	}

	res := &pb.TaskGetResponse{
		Title:   task.Title,
		Edited:  task.Edited.Unix(),
		Created: task.Created.Unix(),
	}

	if task.Description != "" {
		res.Description = &task.Description
	}

	return res, nil
}

func uuidToBytes(ID *uuid.UUID) []byte {
	bytes := make([]byte, 16)
	copy(bytes, ID[0:])

	return bytes
}
