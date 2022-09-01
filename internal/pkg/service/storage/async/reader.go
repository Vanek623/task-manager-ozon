package async

import (
	"context"

	redisPkg "github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
)

type iReader interface {
	ReadAddResponse(ctx context.Context, requestID *uuid.UUID) (*uuid.UUID, error)
	ReadDeleteResponse(ctx context.Context, requestID *uuid.UUID) error
	ReadListResponse(ctx context.Context, requestID *uuid.UUID) ([]*models.Task, error)
	ReadUpdateResponse(ctx context.Context, requestID *uuid.UUID) error
	ReadGetResponse(ctx context.Context, requestID *uuid.UUID) (*models.Task, error)
}

var (
	// ErrNoExistID ошибка отсутствия ключа в хранилище
	ErrNoExistID = errors.New("has no such key")
	// ErrParse ошибка разбора ответа
	ErrParse = errors.New("cannot parse message")
)

// Reader читалка KV хранилища
type Reader struct {
	iReader
}

// NewRedisReader создать читалку для Redis
func NewRedisReader(ctx context.Context, opts *redisPkg.Options) (*Reader, error) {
	r, err := newRedis(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &Reader{r}, nil
}
