package async

import (
	"context"
	"encoding/json"

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

func (r *redis) ReadAddResponse(ctx context.Context, requestID *uuid.UUID) (*uuid.UUID, error) {
	resp, err := r.readBytes(ctx, requestID)
	if err != nil {
		return nil, err
	}

	taskID, err := uuid.FromBytes(resp)
	if err != nil {
		return nil, err
	}

	return &taskID, nil
}

func (r *redis) ReadDeleteResponse(ctx context.Context, requestID *uuid.UUID) error {
	resp, err := r.readMessage(ctx, requestID)
	if err != nil {
		return err
	}

	if resp != "" {
		return errors.New(resp)
	}

	return nil
}

func (r *redis) ReadListResponse(ctx context.Context, requestID *uuid.UUID) ([]*models.Task, error) {
	resp, err := r.readBytes(ctx, requestID)
	if err != nil {
		return nil, err
	}

	var tasks []*models.Task
	if err = json.Unmarshal(resp, &tasks); err != nil {
		return nil, ErrParse
	}

	return tasks, nil
}

func (r *redis) ReadUpdateResponse(ctx context.Context, requestID *uuid.UUID) error {
	resp, err := r.readMessage(ctx, requestID)
	if err != nil {
		return err
	}

	if resp != "" {
		return errors.New(resp)
	}

	return nil
}

func (r *redis) ReadGetResponse(ctx context.Context, requestID *uuid.UUID) (*models.Task, error) {
	resp, err := r.readBytes(ctx, requestID)
	if err != nil {
		return nil, err
	}

	var task *models.Task
	if err = json.Unmarshal(resp, task); err != nil {
		return nil, ErrParse
	}

	return task, nil
}

func (r *redis) readBytes(ctx context.Context, requestID *uuid.UUID) ([]byte, error) {
	res, err := r.client.GetDel(ctx, requestID.String()).Bytes()
	if err == redisPkg.Nil {
		return nil, ErrNoExistID
	} else if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *redis) readMessage(ctx context.Context, requestID *uuid.UUID) (string, error) {
	res, err := r.client.GetDel(ctx, requestID.String()).Result()
	if err == redisPkg.Nil {
		return "", ErrNoExistID
	} else if err != nil {
		return "", err
	}

	return res, nil
}
