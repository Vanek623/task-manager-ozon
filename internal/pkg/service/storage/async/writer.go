package async

import (
	"context"

	"github.com/google/uuid"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/counters"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
	storageModelsPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/storage/models"
)

type iWriter interface {
	WriteAddRequest(ctx context.Context, data *storageModelsPkg.AddTaskData) (*uuid.UUID, error)
	WriteDeleteRequest(ctx context.Context, data *models.DeleteTaskData) (*uuid.UUID, error)
	WriteListRequest(ctx context.Context, data *models.ListTaskData) (*uuid.UUID, error)
	WriteUpdateRequest(ctx context.Context, data *models.UpdateTaskData) (*uuid.UUID, error)
	WriteGetRequest(ctx context.Context, data *models.GetTaskData) (*uuid.UUID, error)
}

// Writer асинхронный отправитель запросов в хранилище
type Writer struct {
	iWriter
}

// NewKafkaWriter асинхронный отправитель запросов в хранилище на основе Kafka
func NewKafkaWriter(ctx context.Context, brokers []string, cs *counters.Counters) (*Writer, error) {
	s, err := newKafka(ctx, brokers, cs)
	if err != nil {
		return nil, err
	}

	return &Writer{s}, nil
}
