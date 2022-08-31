package storage

import (
	"context"

	redisPkg "github.com/go-redis/redis"
	"github.com/google/uuid"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/counters"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/task/models"
)

type redis struct {
	storage iTaskStorage
	client  *redisPkg.Client
	cs      *counters.Counters
}

func (r *redis) Add(ctx context.Context, t *models.Task) error {
	//TODO implement me
	panic("implement me")
}

func (r *redis) Delete(ctx context.Context, ID *uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (r *redis) List(ctx context.Context, limit, offset uint64) ([]*models.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (r *redis) Update(ctx context.Context, t *models.Task) error {
	//TODO implement me
	panic("implement me")
}

func (r *redis) Get(ctx context.Context, ID *uuid.UUID) (*models.Task, error) {
	//TODO implement me
	panic("implement me")
}

func newRedis(opts *redisPkg.Options, storage iTaskStorage, cs *counters.Counters) (*redis, error) {
	client := redisPkg.NewClient(opts)
	if res := client.Ping(); res.Err() != nil {
		return nil, res.Err()
	}

	return &redis{
		storage: storage,
		client:  client,
		cs:      cs,
	}, nil
}
