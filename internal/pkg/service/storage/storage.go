package storage

import (
	"context"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
)

type iStorage interface {
	Add(ctx context.Context, data *models.AddTaskData) (uint64, error)
	Delete(ctx context.Context, data *models.DeleteTaskData) error
	List(ctx context.Context, data *models.ListTaskData) ([]*models.Task, error)
	Update(ctx context.Context, data *models.UpdateTaskData) error
	Get(ctx context.Context, data *models.GetTaskData) (*models.DetailedTask, error)
}

// Storage Хранилище сервиса
type Storage struct {
	iStorage
}

// NewGRPC GRPC хранилище
func NewGRPC(address string) (*Storage, error) {
	s, err := newGRPC(address)
	if err != nil {
		return nil, err
	}
	return &Storage{s}, nil
}
