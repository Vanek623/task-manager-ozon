//go:generate mockgen -source ./service.go -destination=./mocks/service.go -package=mock_service
package service

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
)

// ErrValidation ошибка валидации данных
var ErrValidation = errors.New("invalid data")

type iService interface {
	AddTask(ctx context.Context, data *models.AddTaskData) (uint64, error)
	DeleteTask(ctx context.Context, data *models.DeleteTaskData) error
	TasksList(ctx context.Context, data *models.ListTaskData) ([]*models.Task, error)
	UpdateTask(ctx context.Context, data *models.UpdateTaskData) error
	GetTask(ctx context.Context, data *models.GetTaskData) (*models.DetailedTask, error)
}

type iStorage interface {
	Add(ctx context.Context, data *models.AddTaskData) (uint64, error)
	Delete(ctx context.Context, data *models.DeleteTaskData) error
	List(ctx context.Context, data *models.ListTaskData) ([]*models.Task, error)
	Update(ctx context.Context, data *models.UpdateTaskData) error
	Get(ctx context.Context, data *models.GetTaskData) (*models.DetailedTask, error)
}

// Service структура бизнес логики
type Service struct {
	iService
	storage iStorage
}

// New создать структуру бизнес логики
func New(s iStorage) *Service {
	return &Service{
		storage: s,
	}
}

// AddTask добавить задачу
func (s *Service) AddTask(ctx context.Context, data *models.AddTaskData) (uint64, error) {
	if err := isTitleAndDescriptionOk(data.Title(), data.Description()); err != nil {
		return 0, err
	}

	return s.storage.Add(ctx, data)
}

// DeleteTask удалить задачу
func (s *Service) DeleteTask(ctx context.Context, data *models.DeleteTaskData) error {
	return s.storage.Delete(ctx, data)
}

// TasksList получить список задач
func (s *Service) TasksList(ctx context.Context, data *models.ListTaskData) ([]*models.Task, error) {
	tasks, err := s.storage.List(ctx, data)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// UpdateTask обновить задачу
func (s *Service) UpdateTask(ctx context.Context, data *models.UpdateTaskData) error {
	if err := isTitleAndDescriptionOk(data.Title(), data.Description()); err != nil {
		return err
	}

	return s.storage.Update(ctx, data)
}

// GetTask получить подробное описание задачи
func (s *Service) GetTask(ctx context.Context, data *models.GetTaskData) (*models.DetailedTask, error) {
	task, err := s.storage.Get(ctx, data)
	if err != nil {
		return nil, err
	}

	return task, nil
}
