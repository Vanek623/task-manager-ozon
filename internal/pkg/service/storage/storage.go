package storage

import (
	"context"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/counters"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
	storageModelsPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/storage/models"
)

type iStorage interface {
	Add(ctx context.Context, data *storageModelsPkg.AddTaskData) error
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
func NewGRPC(ctx context.Context, address string, cs *counters.Counters) (*Storage, error) {
	s, err := newGRPC(ctx, address, cs)
	if err != nil {
		return nil, err
	}
	return &Storage{s}, nil
}

// NewKafka Хранилище на основе очереди для операций добавления, изменения, удаления
// и синхронного хранилища для операций чтения
func NewKafka(ctx context.Context, brokers []string, syncStorage iStorage, cs *counters.Counters) (*Storage, error) {
	s, err := newKafka(ctx, brokers, syncStorage, cs)
	if err != nil {
		return nil, err
	}

	return &Storage{s}, nil
}
