package api

import (
	"context"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/counters"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/tracer"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/service"
	"go.opentelemetry.io/otel"
)

type iService interface {
	AddTask(ctx context.Context, data *models.AddTaskData) (*uuid.UUID, error)
	DeleteTask(ctx context.Context, data *models.DeleteTaskData) error
	TasksList(ctx context.Context, data *models.ListTaskData) ([]*models.Task, error)
	UpdateTask(ctx context.Context, data *models.UpdateTaskData) error
	GetTask(ctx context.Context, data *models.GetTaskData) (*models.DetailedTask, error)
}

type implementation struct {
	pb.UnimplementedServiceServer
	s  iService
	cs *counters.Counters
}

//NewAPI создание обработчика сервиса
func NewAPI(s iService, cs *counters.Counters) pb.ServiceServer {
	return &implementation{
		s:  s,
		cs: cs,
	}
}

func (i *implementation) TaskCreate(ctx context.Context, in *pb.TaskCreateRequest) (*pb.TaskCreateResponse, error) {
	i.income()
	newCtx, span := otel.Tracer(tracer.Name).Start(ctx, tracer.MakeSpanName("TaskCreate"))
	defer span.End()

	log.WithFields(log.Fields{
		"title":       in.GetTitle(),
		"description": in.GetDescription(),
	}).Info("Incoming create request")

	data := models.NewAddTaskData(in.GetTitle(), in.GetDescription())

	id, err := i.s.AddTask(newCtx, data)

	if err != nil {
		i.fail(err)
		return nil, err
	}

	i.success()

	return &pb.TaskCreateResponse{ID: uuidToBytes(id)}, nil
}

func (i *implementation) TaskList(ctx context.Context, in *pb.TaskListRequest) (*pb.TaskListResponse, error) {
	i.income()

	log.WithFields(log.Fields{
		"offset": in.GetOffset(),
		"limit":  in.GetLimit(),
	}).Info("Incoming list request")

	data := models.NewListTaskData(in.GetLimit(), in.GetOffset())

	tasks, err := i.s.TasksList(ctx, data)

	if err != nil {
		i.fail(err)
		return nil, err
	}

	result := make([]*pb.TaskListResponse_Task, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, &pb.TaskListResponse_Task{
			ID:    uuidToBytes(task.ID()),
			Title: task.Title(),
		})
	}

	i.success()
	return &pb.TaskListResponse{Tasks: result}, nil
}

func (i *implementation) TaskUpdate(ctx context.Context, in *pb.TaskUpdateRequest) (*pb.TaskUpdateResponse, error) {
	i.income()

	log.WithFields(log.Fields{
		"id":          in.GetID(),
		"title":       in.GetTitle(),
		"description": in.GetDescription(),
	}).Info("Incoming update request")

	id, err := uuid.FromBytes(in.GetID())
	defer func() {
		if err != nil {
			i.fail(err)
		} else {
			i.success()
		}
	}()
	if err != nil {
		return nil, err
	}

	data := models.NewUpdateTaskData(&id, in.GetTitle(), in.GetDescription())

	err = i.s.UpdateTask(ctx, data)
	if err != nil {
		return nil, err
	}

	return &pb.TaskUpdateResponse{}, nil
}

func (i *implementation) TaskDelete(ctx context.Context, in *pb.TaskDeleteRequest) (*pb.TaskDeleteResponse, error) {
	i.income()

	newCtx, span := otel.Tracer(tracer.Name).Start(ctx, tracer.MakeSpanName("TaskDelete"))
	defer span.End()

	log.WithFields(log.Fields{
		"id": in.GetID(),
	}).Info("Incoming delete request")

	id, err := uuid.FromBytes(in.GetID())
	defer func() {
		if err != nil {
			i.fail(err)
		} else {
			i.success()
		}
	}()
	if err != nil {
		return nil, err
	}

	data := models.NewDeleteTaskData(&id)

	err = i.s.DeleteTask(newCtx, data)
	if err != nil {
		return nil, err
	}

	return &pb.TaskDeleteResponse{}, nil
}

func (i *implementation) TaskGet(ctx context.Context, in *pb.TaskGetRequest) (*pb.TaskGetResponse, error) {
	i.income()

	log.WithFields(log.Fields{
		"id": in.GetID(),
	}).Info("Incoming get request")

	id, err := uuid.FromBytes(in.GetID())
	if err != nil {
		i.fail(err)
		return nil, err
	}

	data := models.NewGetTaskData(&id)

	task, err := i.s.GetTask(ctx, data)
	if err != nil {
		i.fail(err)
		return nil, err
	}

	result := pb.TaskGetResponse{
		Title:  task.Title(),
		Edited: task.Edited().Unix(),
	}

	if tmp := task.Description(); tmp != "" {
		result.Description = &tmp
	}

	i.success()
	return &result, nil
}

func uuidToBytes(ID *uuid.UUID) []byte {
	bytes := make([]byte, 16)
	copy(bytes, ID[0:])

	return bytes
}

func (i *implementation) success() {
	i.cs.Inc(counters.Success)
}

func (i *implementation) fail(err error) {
	i.cs.Inc(counters.Fail)
	log.Error(err)
}

func (i *implementation) income() {
	i.cs.Inc(counters.Incoming)
}
