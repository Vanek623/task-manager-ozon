package service

import (
	"context"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/service"
)

type iService interface {
	AddTask(ctx context.Context, data models.AddTaskData) (uint, error)
	DeleteTask(ctx context.Context, data models.DeleteTaskData) error
	TasksList(ctx context.Context, data models.ListTaskData) ([]models.TaskBrief, error)
	UpdateTask(ctx context.Context, data models.UpdateTaskData) error
	GetTask(ctx context.Context, data models.GetTaskData) (*models.DetailedTask, error)
}

type implementation struct {
	pb.UnimplementedServiceServer
	s iService
}

//NewAPI создание обработчика сервиса
func NewAPI(s iService) pb.ServiceServer {
	return &implementation{s: s}
}

func (i *implementation) TaskCreate(ctx context.Context, in *pb.TaskCreateRequest) (*pb.TaskCreateResponse, error) {
	data := models.AddTaskData{
		Title:       in.GetTitle(),
		Description: in.GetDescription(),
	}

	ID, err := i.s.AddTask(ctx, data)
	if err != nil {
		return nil, err
	}

	return &pb.TaskCreateResponse{ID: uint64(ID)}, nil
}

func (i *implementation) TaskList(ctx context.Context, in *pb.TaskListRequest) (*pb.TaskListResponse, error) {
	data := models.ListTaskData{
		Limit:  uint(in.MaxTasksCount),
		Offset: uint(in.Offset),
	}

	tasks, err := i.s.TasksList(ctx, data)
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

func (i *implementation) TaskUpdate(ctx context.Context, in *pb.TaskUpdateRequest) (*pb.TaskUpdateResponse, error) {
	data := models.UpdateTaskData{
		ID:          uint(in.GetID()),
		Title:       in.GetTitle(),
		Description: in.GetDescription(),
	}

	if err := i.s.UpdateTask(ctx, data); err != nil {
		return nil, err
	}

	return &pb.TaskUpdateResponse{}, nil
}

func (i *implementation) TaskDelete(ctx context.Context, in *pb.TaskDeleteRequest) (*pb.TaskDeleteResponse, error) {
	data := models.DeleteTaskData{ID: uint(in.GetID())}

	if err := i.s.DeleteTask(ctx, data); err != nil {
		return nil, err
	}

	return &pb.TaskDeleteResponse{}, nil
}

func (i *implementation) TaskGet(ctx context.Context, in *pb.TaskGetRequest) (*pb.TaskGetResponse, error) {
	data := models.GetTaskData{ID: uint(in.GetID())}

	task, err := i.s.GetTask(ctx, data)
	if err != nil {
		return nil, err
	}

	result := pb.TaskGetResponse{
		Title:  task.Title,
		Edited: task.Edited.Unix(),
	}

	if task.Description != "" {
		result.Description = &task.Description
	}

	return &result, nil
}
