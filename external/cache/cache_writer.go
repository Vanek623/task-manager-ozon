package cache

import (
	"context"

	redisPkg "github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/task/models"
)

type iWriter interface {
	WriteAddResponse(ctx context.Context, ID *uuid.UUID, err error) error
	WriteDeleteResponse(ctx context.Context, ID *uuid.UUID, err error) error
	WriteUpdateResponse(ctx context.Context, ID *uuid.UUID, err error) error
	WriteGetResponse(ctx context.Context, ID *uuid.UUID, task *models.Task, err error) error
	WriteListResponse(ctx context.Context, ID *uuid.UUID, tasks []*models.Task, err error) error
}

// Writer записывает информацию к кэш
type Writer struct {
	iWriter
}

func NewRedisWriter(ctx context.Context, opts *redisPkg.Options) (*Writer, error) {
	r, err := newRedis(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &Writer{r}, nil
}
