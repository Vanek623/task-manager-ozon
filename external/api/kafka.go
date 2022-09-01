package api

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"

	"github.com/Shopify/sarama"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/counters"

	"gitlab.ozon.dev/Vanek623/task-manager-system/external/task/models"
)

const (
	topicAddRequestName    = "income_add_request"
	topicDeleteRequestName = "income_delete_request"
	topicUpdateRequestName = "income_update_request"

	expiration = time.Minute
)

type kafka struct {
	storage iTaskStorage
	cs      *counters.Counters
	rc      *redis.Client
}

func newKafka(storage iTaskStorage, cs *counters.Counters, client *redis.Client) *kafka {
	return &kafka{
		storage: storage,
		cs:      cs,
		rc:      client,
	}
}

// Setup старт сессии
func (k *kafka) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup конец сессии
func (k *kafka) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim запуск цикла чтения
func (k *kafka) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case <-session.Context().Done():
			log.Info("Session done")
			return nil
		case msg, ok := <-claim.Messages():
			if !ok {
				log.Info("Data channel closed")
				return nil
			}

			k.cs.Inc(counters.Incoming)

			if err := k.handleMessage(session.Context(), msg); err != nil {
				k.cs.Inc(counters.Fail)
				log.Error(err)
			} else {
				k.cs.Inc(counters.Success)
			}

			session.MarkMessage(msg, "")
		}
	}
}

func (k *kafka) handleMessage(ctx context.Context, msg *sarama.ConsumerMessage) error {
	switch msg.Topic {
	case topicAddRequestName:
		return k.addTask(ctx, msg.Value)
	case topicDeleteRequestName:
		return k.deleteTask(ctx, msg.Value)
	case topicUpdateRequestName:
		return k.updateTask(ctx, msg.Value)
	default:
		return errors.Errorf("unknown topic - %s", msg.Topic)
	}
}

func (k *kafka) addTask(ctx context.Context, data []byte) error {
	obj := struct {
		ID          uuid.UUID
		Title       string
		Description string
	}{}

	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}

	err = k.storage.Add(ctx, &models.Task{
		ID:          obj.ID,
		Title:       obj.Title,
		Description: obj.Description,
	})

	return err
}

func (k *kafka) deleteTask(ctx context.Context, data []byte) error {
	obj := struct {
		ID uuid.UUID
	}{}

	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}

	err = k.storage.Delete(ctx, &obj.ID)

	return err
}

func (k *kafka) updateTask(ctx context.Context, data []byte) error {
	obj := struct {
		ID          uuid.UUID
		Title       string
		Description string
	}{}

	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}

	err = k.storage.Update(ctx, &models.Task{
		ID:          obj.ID,
		Title:       obj.Title,
		Description: obj.Description,
	})

	return err
}

func (k *kafka) getTask(ctx context.Context, data []byte) error {
	obj := struct {
		RequestID uuid.UUID
		ID        uuid.UUID
	}{}

	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}

	resp, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	_, err = k.rc.Set(ctx, obj.RequestID.String(), resp, expiration).Result()
	if err != nil {
		return err
	}

	return nil
}

func (k *kafka) listTasks(ctx context.Context, data []byte) error {
	obj := struct {
		RequestID     uuid.UUID
		Limit, Offset uint64
	}{}

	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}

	tasks, err := k.storage.List(ctx, obj.Limit, obj.Offset)
	if err != nil {
		return err
	}

	resp, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	_, err = k.rc.Set(ctx, obj.RequestID.String(), resp, expiration).Result()
	if err != nil {
		return err
	}

	return nil
}
