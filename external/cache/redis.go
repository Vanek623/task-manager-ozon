package cache

import (
	"context"
	"encoding/json"
	"time"

	redisPkg "github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
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
	bytes, marshErr := json.Marshal(task)
	if marshErr != nil {
		return marshErr
	}

	log.WithField("task", task).Debug("Write to redis")

	return r.sendResponse(ctx, ID, bytes, err)
}

func (r *redis) WriteListResponse(ctx context.Context, ID *uuid.UUID, tasks []*models.Task, err error) error {
	bytes, marshErr := json.Marshal(tasks)
	if marshErr != nil {
		return marshErr
	}

	return r.sendResponse(ctx, ID, bytes, err)
}

const (
	connectTimeout    = 5 * time.Second
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

type dto struct {
	Data []byte
	Err  error
}

func (r *redis) sendResponse(ctx context.Context, ID *uuid.UUID, data []byte, err error) error {
	d := dto{}
	if err != nil {
		d.Err = err
	} else {
		d.Data = data
	}

	resp, err := json.Marshal(d)
	if err != nil {
		return err
	}

	if _, err = r.c.Set(ctx, ID.String(), resp, expirationTimeout).Result(); err != nil {
		return err
	}

	return nil
}
