package cache

import (
	"context"
	"time"

	redisPkg "github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/task/models"
)

type redis struct {
	c *redisPkg.Client
}

func (r *redis) WriteAddResponse(ctx context.Context, ID *uuid.UUID, err error) error {
	return r.sendResponse(ctx, ID, nil, err)
}

func (r *redis) WriteDeleteResponse(ctx context.Context, ID *uuid.UUID, err error) error {
	return r.sendResponse(ctx, ID, nil, err)
}

func (r *redis) WriteUpdateResponse(ctx context.Context, ID *uuid.UUID, err error) error {
	return r.sendResponse(ctx, ID, nil, err)
}

func (r *redis) WriteGetResponse(ctx context.Context, ID *uuid.UUID, task *models.Task, err error) error {
	return r.sendResponse(ctx, ID, task, err)
}

func (r *redis) WriteListResponse(ctx context.Context, ID *uuid.UUID, tasks []*models.Task, err error) error {
	return r.sendResponse(ctx, ID, tasks, err)
}

const (
	connectTimeout    = 5 * time.Second
	writeTimeout      = 100 * time.Millisecond
	expirationTimeout = time.Second
)

func newRedis(ctx context.Context, opts *redisPkg.Options) (*redis, error) {
	client := redisPkg.NewClient(opts)

	pingCtx, cl := context.WithTimeout(ctx, connectTimeout)
	defer cl()
	if err := client.Ping(pingCtx).Err(); err != nil {
		return nil, err
	}

	return &redis{c: client}, nil
}

func (r *redis) sendResponse(ctx context.Context, ID *uuid.UUID, data interface{}, prevErr error) error {
	if prevErr != nil {
		data = prevErr
	}

	writeCtx, cl := context.WithTimeout(ctx, writeTimeout)
	defer cl()
	if _, err := r.c.Set(writeCtx, ID.String(), data, expirationTimeout).Result(); err != nil {
		return err
	}

	return prevErr
}
