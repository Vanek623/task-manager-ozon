package async

import (
	"context"
	"encoding/json"
	"time"

	redisPkg "github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
)

type redis struct {
	iReader
	client *redisPkg.Client
}

type dto struct {
	Data []byte
	Err  error
}

type storageTask struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Edited      time.Time `json:"edited"`
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
	_, err := r.read(ctx, requestID)
	if err != nil {
		return err
	}

	return nil
}

func (r *redis) ReadDeleteResponse(ctx context.Context, requestID *uuid.UUID) error {
	_, err := r.read(ctx, requestID)
	if err != nil {
		return err
	}

	return nil
}

func (r *redis) ReadListResponse(ctx context.Context, requestID *uuid.UUID) ([]*models.Task, error) {
	resp, err := r.read(ctx, requestID)
	if err != nil {
		return nil, err
	}

	var tmp []*storageTask
	if err = json.Unmarshal(resp, &tmp); err != nil {
		return nil, err
	}

	tasks := make([]*models.Task, 0, len(tmp))
	for _, t := range tmp {
		tasks = append(tasks, models.NewTask(&t.ID, t.Title))
	}

	return tasks, nil
}

func (r *redis) ReadUpdateResponse(ctx context.Context, requestID *uuid.UUID) error {
	_, err := r.read(ctx, requestID)
	if err != nil {
		return err
	}

	return nil
}

func (r *redis) ReadGetResponse(ctx context.Context, requestID *uuid.UUID) (*models.DetailedTask, error) {
	resp, err := r.read(ctx, requestID)
	if err != nil {
		return nil, err
	}

	var tmp storageTask
	if err = json.Unmarshal(resp, &tmp); err != nil {
		return nil, err
	}

	log.WithField("task", tmp).Debug("Read from redis")

	return models.NewDetailedTask(tmp.Title, tmp.Description, tmp.Edited), nil
}

func (r *redis) read(ctx context.Context, requestID *uuid.UUID) ([]byte, error) {
	res := r.client.GetDel(ctx, requestID.String())
	if err := res.Err(); err == redisPkg.Nil {
		return nil, ErrNoExistID
	} else if err != nil {
		return nil, err
	}

	bytes, _ := res.Bytes()
	var tmp dto
	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return nil, err
	}
	if tmp.Err != nil {
		return nil, tmp.Err
	}

	return tmp.Data, nil
}
