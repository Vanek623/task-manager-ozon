package service

import (
	"context"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
	storageModelsPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/storage/models"
)

type iSyncStorage interface {
	Add(ctx context.Context, data *storageModelsPkg.AddTaskData) error
	Delete(ctx context.Context, data *models.DeleteTaskData) error
	List(ctx context.Context, data *models.ListTaskData) ([]*models.Task, error)
	Update(ctx context.Context, data *models.UpdateTaskData) error
	Get(ctx context.Context, data *models.GetTaskData) (*models.DetailedTask, error)
}

type serviceWithSyncStorage struct {
	storage iSyncStorage
}

// AddTask добавить задачу
func (s *serviceWithSyncStorage) AddTask(ctx context.Context, data *models.AddTaskData) (*uuid.UUID, error) {
	id := uuid.New()
	sData := storageModelsPkg.NewAddTaskData(&id, data.Title(), data.Description())

	log.Debugf("Adding task %s", id)

	if err := s.storage.Add(ctx, sData); err != nil {
		return nil, err
	}

	return &id, nil
}

// DeleteTask удалить задачу
func (s *serviceWithSyncStorage) DeleteTask(ctx context.Context, data *models.DeleteTaskData) error {
	return s.storage.Delete(ctx, data)
}

// TasksList получить список задач
func (s *serviceWithSyncStorage) TasksList(ctx context.Context, data *models.ListTaskData) ([]*models.Task, error) {
	return s.storage.List(ctx, data)
}

// UpdateTask обновить задачу
func (s *serviceWithSyncStorage) UpdateTask(ctx context.Context, data *models.UpdateTaskData) error {
	return s.storage.Update(ctx, data)
}

// GetTask получить подробное описание задачи
func (s *serviceWithSyncStorage) GetTask(ctx context.Context, data *models.GetTaskData) (*models.DetailedTask, error) {
	return s.storage.Get(ctx, data)
}
