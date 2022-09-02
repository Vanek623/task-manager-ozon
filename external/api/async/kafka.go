package async

import (
	"context"
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/counters"

	"gitlab.ozon.dev/Vanek623/task-manager-system/external/task/models"
)

const (
	topicAddRequestName    = "income_add_request"
	topicDeleteRequestName = "income_delete_request"
	topicUpdateRequestName = "income_update_request"
	topicGetRequestName    = "income_get_request"
	topicListRequestName   = "income_list_request"
)

type iCacheWriter interface {
	WriteAddResponse(ctx context.Context, ID *uuid.UUID, err error) error
	WriteDeleteResponse(ctx context.Context, ID *uuid.UUID, err error) error
	WriteUpdateResponse(ctx context.Context, ID *uuid.UUID, err error) error
	WriteGetResponse(ctx context.Context, ID *uuid.UUID, task *models.Task, err error) error
	WriteListResponse(ctx context.Context, ID *uuid.UUID, tasks []*models.Task, err error) error
}

type kafka struct {
	storage iTaskStorage
	cs      *counters.Counters
	cw      iCacheWriter
}

func newKafka(storage iTaskStorage, cs *counters.Counters, writer iCacheWriter) *kafka {
	return &kafka{
		storage: storage,
		cs:      cs,
		cw:      writer,
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
	case topicGetRequestName:
		return k.getTask(ctx, msg.Value)
	case topicListRequestName:
		return k.listTasks(ctx, msg.Value)
	default:
		return errors.Errorf("unknown topic - %s", msg.Topic)
	}
}

func (k *kafka) addTask(ctx context.Context, data []byte) error {
	req := struct {
		RequestID   uuid.UUID
		ID          uuid.UUID
		Title       string
		Description string
	}{}

	err := json.Unmarshal(data, &req)
	if err != nil {
		return err
	}

	err = k.storage.Add(ctx, &models.Task{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
	})

	return k.cw.WriteAddResponse(ctx, &req.RequestID, err)
}

func (k *kafka) deleteTask(ctx context.Context, data []byte) error {
	req := struct {
		RequestID uuid.UUID
		ID        uuid.UUID
	}{}

	err := json.Unmarshal(data, &req)
	if err != nil {
		return err
	}

	err = k.storage.Delete(ctx, &req.ID)

	return k.cw.WriteDeleteResponse(ctx, &req.RequestID, err)
}

func (k *kafka) updateTask(ctx context.Context, data []byte) error {
	req := struct {
		RequestID   uuid.UUID
		ID          uuid.UUID
		Title       string
		Description string
	}{}

	err := json.Unmarshal(data, &req)
	if err != nil {
		return err
	}

	err = k.storage.Update(ctx, &models.Task{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
	})

	return k.cw.WriteUpdateResponse(ctx, &req.RequestID, err)
}

func (k *kafka) getTask(ctx context.Context, data []byte) error {
	req := struct {
		RequestID uuid.UUID
		ID        uuid.UUID
	}{}

	err := json.Unmarshal(data, &req)
	if err != nil {
		return err
	}

	task, err := k.storage.Get(ctx, &req.ID)

	return k.cw.WriteGetResponse(ctx, &req.RequestID, task, err)
}

func (k *kafka) listTasks(ctx context.Context, data []byte) error {
	req := struct {
		RequestID     uuid.UUID
		Limit, Offset uint64
	}{}

	err := json.Unmarshal(data, &req)
	if err != nil {
		return err
	}

	tasks, err := k.storage.List(ctx, req.Limit, req.Offset)

	return k.cw.WriteListResponse(ctx, &req.RequestID, tasks, err)
}
