package api

import (
	"context"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api"
)

type iTaskStorage interface {
	Add(ctx context.Context, t models.Task) error
	Delete(ctx context.Context, ID uint) error
	List(ctx context.Context) ([]models.Task, error)
	Update(ctx context.Context, t models.Task) error
	Get(ctx context.Context, ID uint) (*models.Task, error)
}

type implementation struct {
	pb.UnimplementedAdminServer
	tm iTaskStorage
}

//New создание обработчика
func New(tm iTaskStorage) pb.AdminServer {
	return &implementation{tm: tm}
}

func (i implementation) TaskCreate(ctx context.Context, in *pb.TaskCreateRequest) (*pb.TaskCreateResponse, error) {
	task := models.Task{
		Title:       in.GetTitle(),
		Description: in.GetDescription(),
	}

	if err := i.tm.Add(ctx, task); err != nil {
		return nil, err
	}

	return &pb.TaskCreateResponse{}, nil
}

func (i implementation) TaskList(ctx context.Context, _ *pb.TaskListRequest) (*pb.TaskListResponse, error) {
	tasks, err := i.tm.List(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*pb.TaskListResponse_Task, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, &pb.TaskListResponse_Task{
			ID:    uint64(task.ID),
			Title: task.Title,
		})
	}

	return &pb.TaskListResponse{Tasks: result}, nil
}

func (i implementation) TaskUpdate(ctx context.Context, in *pb.TaskUpdateRequest) (*pb.TaskUpdateResponse, error) {
	task, err := i.tm.Get(ctx, uint(in.GetID()))
	if err != nil {
		return nil, err
	}

	task.Title = in.GetTitle()
	task.Description = in.GetDescription()

	if err = i.tm.Update(ctx, *task); err != nil {
		return nil, err
	}

	return &pb.TaskUpdateResponse{}, nil
}

func (i implementation) TaskDelete(ctx context.Context, in *pb.TaskDeleteRequest) (*pb.TaskDeleteResponse, error) {
	if err := i.tm.Delete(ctx, uint(in.GetID())); err != nil {
		return nil, err
	}

	return &pb.TaskDeleteResponse{}, nil
}

func (i implementation) TaskGet(ctx context.Context, in *pb.TaskGetRequest) (*pb.TaskGetResponse, error) {
	task, err := i.tm.Get(ctx, uint(in.GetID()))
	if err != nil {
		return nil, err
	}

	result := pb.TaskGetResponse{
		Title: task.Title,
		Tm:    task.Created.Unix(),
	}

	if task.Description != "" {
		result.Description = &task.Description
	}

	return &result, nil
}
