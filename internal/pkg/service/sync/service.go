//go:generate mockgen -source ./service.go -destination=./mocks/service.go -package=mock_service

package sync

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
	storageModelsPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/storage/models"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/validation"
)

// ErrValidation ошибка валидации данных
var ErrValidation = errors.New("invalid data")

type iService interface {
	AddTask(ctx context.Context, data *models.AddTaskData) (*uuid.UUID, error)
	DeleteTask(ctx context.Context, data *models.DeleteTaskData) error
	TasksList(ctx context.Context, data *models.ListTaskData) ([]*models.Task, error)
	UpdateTask(ctx context.Context, data *models.UpdateTaskData) error
	GetTask(ctx context.Context, data *models.GetTaskData) (*models.DetailedTask, error)
}

type iStorage interface {
	Add(ctx context.Context, data *storageModelsPkg.AddTaskData) error
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
func (s *Service) AddTask(ctx context.Context, data *models.AddTaskData) (*uuid.UUID, error) {
	if err := validation.IsTitleAndDescriptionOk(data.Title(), data.Description()); err != nil {
		return nil, err
	}

	id := uuid.New()
	sData := storageModelsPkg.NewAddTaskData(&id, data.Title(), data.Description())

	log.Debugf("Adding task %s", id)

	if err := s.storage.Add(ctx, sData); err != nil {
		return nil, err
	}

	return &id, nil
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
	if err := validation.IsTitleAndDescriptionOk(data.Title(), data.Description()); err != nil {
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
