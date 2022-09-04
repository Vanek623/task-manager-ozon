//go:generate mockgen -source ./service.go -destination=./mocks/service.go -package=mock_service

package service

import (
	"context"

	"github.com/google/uuid"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/validation"
)

type iService interface {
	AddTask(ctx context.Context, data *models.AddTaskData) (*uuid.UUID, error)
	DeleteTask(ctx context.Context, data *models.DeleteTaskData) error
	TasksList(ctx context.Context, data *models.ListTaskData) ([]*models.Task, error)
	UpdateTask(ctx context.Context, data *models.UpdateTaskData) error
	GetTask(ctx context.Context, data *models.GetTaskData) (*models.DetailedTask, error)
}

// Service структура бизнес логики
type Service struct {
	iService
}

// AddTask добавить задачу
func (s *Service) AddTask(ctx context.Context, data *models.AddTaskData) (*uuid.UUID, error) {
	if err := validation.IsTitleAndDescriptionOk(data.Title(), data.Description()); err != nil {
		return nil, err
	}

	return s.iService.AddTask(ctx, data)
}

// UpdateTask обновить задачу
func (s *Service) UpdateTask(ctx context.Context, data *models.UpdateTaskData) error {
	if err := validation.IsTitleAndDescriptionOk(data.Title(), data.Description()); err != nil {
		return err
	}

	return s.iService.UpdateTask(ctx, data)
}

// NewServiceWithSyncStorage создать структуру бизнес логики c синхронным общением с хранилищем
func NewServiceWithSyncStorage(s iSyncStorage) *Service {
	return &Service{
		&serviceWithSyncStorage{
			storage: s,
		},
	}
}

// NewServiceWithAsyncStorage создать структуру бизнес логики c асинхронным общением с хранилищем
func NewServiceWithAsyncStorage(sw iStorageWriter, sr iStorageReader) *Service {
	return &Service{
		&serviceWithAsyncStorage{
			sw: sw,
			sr: sr,
		},
	}
}
