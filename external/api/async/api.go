package async

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/google/uuid"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/counters"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/task/models"
)

type iTaskStorage interface {
	Add(ctx context.Context, t *models.Task) error
	Delete(ctx context.Context, ID *uuid.UUID) error
	List(ctx context.Context, limit, offset uint64) ([]*models.Task, error)
	Update(ctx context.Context, t *models.Task) error
	Get(ctx context.Context, ID *uuid.UUID) (*models.Task, error)
}

// NewKafkaAPI создание обработчика асинхронных запросов
func NewKafkaAPI(s iTaskStorage, cs *counters.Counters, cw iCacheWriter) sarama.ConsumerGroupHandler {
	return newKafka(s, cs, cw)
}
