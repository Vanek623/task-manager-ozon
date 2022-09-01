package api

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/google/uuid"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/counters"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/task/models"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/storage"
)

type iTaskStorage interface {
	Add(ctx context.Context, t *models.Task) error
	Delete(ctx context.Context, ID *uuid.UUID) error
	List(ctx context.Context, limit, offset uint64) ([]*models.Task, error)
	Update(ctx context.Context, t *models.Task) error
	Get(ctx context.Context, ID *uuid.UUID) (*models.Task, error)
}

// NewProtobufAPI создание обработчика синхронных запросов
func NewProtobufAPI(s iTaskStorage, cs *counters.Counters) pb.StorageServer {
	return newProtobuf(s, cs)
}

// NewKafkaAPI создание обработчика асинхронных запросов
func NewKafkaAPI(s iTaskStorage, cs *counters.Counters) sarama.ConsumerGroupHandler {
	return newKafka(s, cs)
}
