package async

import (
	"context"

	"github.com/google/uuid"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/counters"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
	storageModelsPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/storage/models"
)

type iStorageRequester interface {
	Add(ctx context.Context, data *storageModelsPkg.AddTaskData) (*uuid.UUID, error)
	Delete(ctx context.Context, data *models.DeleteTaskData) (*uuid.UUID, error)
	List(ctx context.Context, data *models.ListTaskData) (*uuid.UUID, error)
	Update(ctx context.Context, data *models.UpdateTaskData) (*uuid.UUID, error)
	Get(ctx context.Context, data *models.GetTaskData) (*uuid.UUID, error)
}

type iStorageResponder interface {
	ReadAddResponse(ctx context.Context, ID *uuid.UUID) (*uuid.UUID, error)
	ReadDeleteResponse(ctx context.Context, ID *uuid.UUID) (*uuid.UUID, error)
	ReadListResponse(ctx context.Context, ID *uuid.UUID) (*uuid.UUID, error)
	ReadUpdateResponse(ctx context.Context, ID *uuid.UUID) (*uuid.UUID, error)
	ReadGetResponse(ctx context.Context, ID *uuid.UUID) (*uuid.UUID, error)
}

type StorageRequester struct {
	iStorageRequester
}

type StorageResponder struct {
	iStorageResponder
}

// NewKafka Хранилище на основе очереди для операций добавления, изменения, удаления
// и синхронного хранилища для операций чтения
func NewKafka(ctx context.Context, brokers []string, cs *counters.Counters) (*Storage, error) {
	s, err := newKafka(ctx, brokers, cs)
	if err != nil {
		return nil, err
	}

	return &Storage{s}, nil
}
