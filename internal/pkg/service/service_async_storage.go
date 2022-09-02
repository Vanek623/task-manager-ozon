package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/storage/async"
	storageModelsPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/storage/models"
)

type iStorageReader interface {
	ReadAddResponse(ctx context.Context, requestID *uuid.UUID) error
	ReadDeleteResponse(ctx context.Context, requestID *uuid.UUID) error
	ReadListResponse(ctx context.Context, requestID *uuid.UUID) ([]*models.Task, error)
	ReadUpdateResponse(ctx context.Context, requestID *uuid.UUID) error
	ReadGetResponse(ctx context.Context, requestID *uuid.UUID) (*models.DetailedTask, error)
}

type iStorageWriter interface {
	WriteAddRequest(ctx context.Context, data *storageModelsPkg.AddTaskData) (*uuid.UUID, error)
	WriteDeleteRequest(ctx context.Context, data *models.DeleteTaskData) (*uuid.UUID, error)
	WriteListRequest(ctx context.Context, data *models.ListTaskData) (*uuid.UUID, error)
	WriteUpdateRequest(ctx context.Context, data *models.UpdateTaskData) (*uuid.UUID, error)
	WriteGetRequest(ctx context.Context, data *models.GetTaskData) (*uuid.UUID, error)
}

const (
	maxReadAttempts    = 3
	nextAttemptTimeout = 100 * time.Millisecond
)

type serviceWithAsyncStorage struct {
	sw iStorageWriter
	sr iStorageReader
}

func (s *serviceWithAsyncStorage) AddTask(ctx context.Context, data *models.AddTaskData) (*uuid.UUID, error) {
	id := uuid.New()
	sData := storageModelsPkg.NewAddTaskData(&id, data.Title(), data.Description())

	reqID, err := s.sw.WriteAddRequest(ctx, sData)
	if err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{
		"task ID":    id.String(),
		"request ID": reqID.String(),
	}).Debug("Waiting for add task")

	for i := 0; i < maxReadAttempts; i++ {
		if err = s.sr.ReadAddResponse(ctx, reqID); err == nil {
			log.Debug("Task added")
			return &id, nil
		} else if i+1 < maxReadAttempts {
			if errors.Is(err, async.ErrNoExistID) {
				log.Error(err)
			}
			time.Sleep(nextAttemptTimeout)
		}
	}

	return nil, err
}

func (s *serviceWithAsyncStorage) DeleteTask(ctx context.Context, data *models.DeleteTaskData) error {
	reqID, err := s.sw.WriteDeleteRequest(ctx, data)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"task ID":    data.ID().String(),
		"request ID": reqID.String(),
	}).Debug("Waiting for delete task")

	for i := 0; i < maxReadAttempts; i++ {
		if err = s.sr.ReadDeleteResponse(ctx, reqID); err == nil {
			log.Debug("Task deleted")
			return nil
		} else if i+1 < maxReadAttempts {
			if !errors.Is(err, async.ErrNoExistID) {
				log.Error(err)
			}
			time.Sleep(nextAttemptTimeout)
		}
	}

	return err
}

func (s *serviceWithAsyncStorage) TasksList(ctx context.Context, data *models.ListTaskData) ([]*models.Task, error) {
	reqID, err := s.sw.WriteListRequest(ctx, data)
	if err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{
		"request ID": reqID.String(),
		"limit":      data.Limit(),
		"offset":     data.Offset(),
	}).Debug("Waiting for list tasks")

	var tasks []*models.Task
	for i := 0; i < maxReadAttempts; i++ {
		if tasks, err = s.sr.ReadListResponse(ctx, reqID); err == nil {
			log.Debug("Tasks got")
			return tasks, nil
		} else if i+1 < maxReadAttempts {
			if !errors.Is(err, async.ErrNoExistID) {
				log.Error(err)
			}

			time.Sleep(nextAttemptTimeout)
		}
	}

	return nil, err
}

func (s *serviceWithAsyncStorage) UpdateTask(ctx context.Context, data *models.UpdateTaskData) error {
	reqID, err := s.sw.WriteUpdateRequest(ctx, data)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"task ID":    data.ID().String(),
		"request ID": reqID.String(),
	}).Debug("Waiting for update task")

	for i := 0; i < maxReadAttempts; i++ {
		if err = s.sr.ReadUpdateResponse(ctx, reqID); err == nil {
			log.Debug("Task updated")
			return nil
		} else if i+1 < maxReadAttempts {
			if errors.Is(err, async.ErrNoExistID) {
				log.Error(err)
			}
			time.Sleep(nextAttemptTimeout)
		}
	}

	return err
}

func (s *serviceWithAsyncStorage) GetTask(ctx context.Context, data *models.GetTaskData) (*models.DetailedTask, error) {
	reqID, err := s.sw.WriteGetRequest(ctx, data)
	if err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{
		"request ID": reqID.String(),
		"task ID":    data.ID().String(),
	}).Debug("Waiting for get task")

	var task *models.DetailedTask
	for i := 0; i < maxReadAttempts; i++ {
		if task, err = s.sr.ReadGetResponse(ctx, reqID); err == nil {
			log.Debug("Task got")
			return task, nil
		} else if i+1 < maxReadAttempts {
			if !errors.Is(err, async.ErrNoExistID) {
				log.Error(err)
			}

			time.Sleep(nextAttemptTimeout)
		}
	}

	return nil, err
}
