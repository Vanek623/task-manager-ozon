package service

import (
	"context"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/service"
)

type iService interface {
	AddTask(ctx context.Context, data *models.AddTaskData) (uint64, error)
	DeleteTask(ctx context.Context, data *models.DeleteTaskData) error
	TasksList(ctx context.Context, data *models.ListTaskData) ([]*models.Task, error)
	UpdateTask(ctx context.Context, data *models.UpdateTaskData) error
	GetTask(ctx context.Context, data *models.GetTaskData) (*models.DetailedTask, error)
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
	data := models.NewAddTaskData(in.GetTitle(), in.GetDescription())

	ID, err := i.s.AddTask(ctx, data)
	if err != nil {
		return nil, err
	}

	return &pb.TaskCreateResponse{ID: ID}, nil
}

func (i *implementation) TaskList(ctx context.Context, in *pb.TaskListRequest) (*pb.TaskListResponse, error) {
	data := models.NewListTaskData(in.GetLimit(), in.GetOffset())

	tasks, err := i.s.TasksList(ctx, data)
	if err != nil {
		return nil, err
	}

	result := make([]*pb.TaskListResponse_Task, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, &pb.TaskListResponse_Task{
			ID:    task.ID(),
			Title: task.Title(),
		})
	}

	return &pb.TaskListResponse{Tasks: result}, nil
}

func (i *implementation) TaskUpdate(ctx context.Context, in *pb.TaskUpdateRequest) (*pb.TaskUpdateResponse, error) {
	data := models.NewUpdateTaskData(in.GetID(), in.GetTitle(), in.GetDescription())

	if err := i.s.UpdateTask(ctx, data); err != nil {
		return nil, err
	}

	return &pb.TaskUpdateResponse{}, nil
}

func (i *implementation) TaskDelete(ctx context.Context, in *pb.TaskDeleteRequest) (*pb.TaskDeleteResponse, error) {
	data := models.NewDeleteTaskData(in.GetID())

	if err := i.s.DeleteTask(ctx, data); err != nil {
		return nil, err
	}

	return &pb.TaskDeleteResponse{}, nil
}

func (i *implementation) TaskGet(ctx context.Context, in *pb.TaskGetRequest) (*pb.TaskGetResponse, error) {
	data := models.NewGetTaskData(in.GetID())

	task, err := i.s.GetTask(ctx, data)
	if err != nil {
		return nil, err
	}

	result := pb.TaskGetResponse{
		Title:  task.Title(),
		Edited: task.Edited().Unix(),
	}

	if tmp := task.Description(); tmp != "" {
		result.Description = &tmp
	}

	return &result, nil
}
