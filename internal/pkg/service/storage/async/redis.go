package async

import (
	"context"

	redisPkg "github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
)

type redis struct {
	iReader
	client *redisPkg.Client
}

func newRedis(ctx context.Context, opts *redisPkg.Options) (r *redis, err error) {
	client := redisPkg.NewClient(opts)
	if _, err = client.Ping(ctx).Result(); err != nil {
		return
	}

	r = &redis{
		client: client,
	}

	go func() {
		<-ctx.Done()
		if err = client.Close(); err != nil {
			log.Error(err)
		}
	}()

	return
}

func (r *redis) ReadAddResponse(ctx context.Context, requestID *uuid.UUID) error {
	resp, err := r.read(ctx, requestID)
	if err != nil {
		return err
	}

	if str := resp.String(); str != "" {
		return errors.New(str)
	}

	return nil
}

func (r *redis) ReadDeleteResponse(ctx context.Context, requestID *uuid.UUID) error {
	resp, err := r.read(ctx, requestID)
	if err != nil {
		return err
	}

	if str := resp.String(); str != "" {
		return errors.New(str)
	}

	return nil
}

func (r *redis) ReadListResponse(ctx context.Context, requestID *uuid.UUID) ([]*models.Task, error) {
	resp, mainErr := r.read(ctx, requestID)
	if mainErr != nil {
		return nil, mainErr
	}

	var tasks []*models.Task
	if err := resp.Scan(&tasks); err == nil {
		return tasks, nil
	}

	if err := resp.Scan(mainErr); err != nil {
		return nil, err
	}

	return nil, mainErr
}

func (r *redis) ReadUpdateResponse(ctx context.Context, requestID *uuid.UUID) error {
	resp, err := r.read(ctx, requestID)
	if err != nil {
		return err
	}

	if str := resp.String(); str != "" {
		return errors.New(str)
	}

	return nil
}

func (r *redis) ReadGetResponse(ctx context.Context, requestID *uuid.UUID) (*models.DetailedTask, error) {
	resp, mainErr := r.read(ctx, requestID)
	if mainErr != nil {
		return nil, mainErr
	}

	var task *models.DetailedTask
	if err := resp.Scan(task); err == nil {
		return task, nil
	}

	if err := resp.Scan(mainErr); err != nil {
		return nil, err
	}

	return nil, mainErr
}

func (r *redis) read(ctx context.Context, requestID *uuid.UUID) (*redisPkg.StringCmd, error) {
	res := r.client.GetDel(ctx, requestID.String())
	if err := res.Err(); err == nil {
		return res, nil
	} else if err == redisPkg.Nil {
		return nil, ErrNoExistID
	} else if err != nil {
		return nil, err
	}

	return res, nil
}
