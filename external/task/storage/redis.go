package storage

import (
	"context"
	"time"

	redisPkg "github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
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

func newRedis(ctx context.Context, opts *redisPkg.Options, storage iTaskStorage, cs *counters.Counters) (*redis, error) {
	client := redisPkg.NewClient(opts)

	pingCtx, cl := context.WithTimeout(ctx, time.Second)
	defer cl()

	if res := client.Ping(pingCtx); res.Err() != nil {
		return nil, res.Err()
	}

	go func() {
		<-ctx.Done()
		if err := client.Close(); err != nil {
			log.Error(err)
		}
	}()

	return &redis{
		storage: storage,
		client:  client,
		cs:      cs,
	}, nil
}
