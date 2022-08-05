package service

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	taskModelsPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/storage"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
)

// ErrValidation ошибка валидации данных
var ErrValidation = errors.New("invalid data")

type iService interface {
	AddTask(ctx context.Context, data models.AddTaskData) (uint, error)
	DeleteTask(ctx context.Context, data models.DeleteTaskData) error
	TasksList(ctx context.Context, data models.ListTaskData) ([]models.Task, error)
	UpdateTask(ctx context.Context, data models.UpdateTaskData) error
	GetTask(ctx context.Context, data models.GetTaskData) (*models.DetailedTask, error)
}

type iTaskStorage interface {
	Add(ctx context.Context, t taskModelsPkg.Task) (uint, error)
	Delete(ctx context.Context, ID uint) error
	List(ctx context.Context) ([]taskModelsPkg.Task, error)
	Update(ctx context.Context, t taskModelsPkg.Task) error
	Get(ctx context.Context, ID uint) (*taskModelsPkg.Task, error)
}

type Service struct {
	iService
	storage iTaskStorage
}

func NewService(pool *pgxpool.Pool) *Service {
	return &Service{
		storage: storage.NewPostgres(pool, true),
	}
}

func (s *Service) AddTask(ctx context.Context, data models.AddTaskData) (uint, error) {
	if err := isTitleAndDescriptionOk(data.Title, data.Description); err != nil {
		return 0, err
	}

	task := taskModelsPkg.Task{
		Title:       data.Title,
		Description: data.Description,
	}

	return s.storage.Add(ctx, task)
}

func (s *Service) DeleteTask(ctx context.Context, data models.DeleteTaskData) error {
	return s.storage.Delete(ctx, data.ID)
}

func (s *Service) TasksList(ctx context.Context, _ models.ListTaskData) ([]models.Task, error) {
	tasks, err := s.storage.List(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]models.Task, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, models.Task{
			ID:    task.ID,
			Title: task.Title,
		})
	}

	return result, nil
}

func (s *Service) UpdateTask(ctx context.Context, data models.UpdateTaskData) error {
	if err := isTitleAndDescriptionOk(data.Title, data.Description); err != nil {
		return err
	}

	task := taskModelsPkg.Task{
		ID:          data.ID,
		Title:       data.Title,
		Description: data.Description,
	}

	return s.storage.Update(ctx, task)
}

func (s *Service) GetTask(ctx context.Context, data models.GetTaskData) (*models.DetailedTask, error) {
	task, err := s.storage.Get(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	return &models.DetailedTask{
		Title:       task.Title,
		Description: task.Description,
		Edited:      task.Edited,
	}, nil
}
