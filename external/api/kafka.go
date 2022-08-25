package api

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"log"
	"time"

	"gitlab.ozon.dev/Vanek623/task-manager-system/external/task/models"
)

const (
	topicAddRequestName    = "income_add_request"
	topicDeleteRequestName = "income_delete_request"
	topicUpdateRequestName = "income_update_request"
)

type iTaskStorage interface {
	Add(ctx context.Context, t *models.Task) error
	Delete(ctx context.Context, ID *uuid.UUID) error
	List(ctx context.Context, limit, offset uint64) ([]*models.Task, error)
	Update(ctx context.Context, t *models.Task) error
	Get(ctx context.Context, ID *uuid.UUID) (*models.Task, error)
}

type kafkaConsumer struct {
	client  sarama.ConsumerGroup
	storage iTaskStorage
}

func newKafka(brokers []string, group string) (*kafkaConsumer, error) {
	cfg := sarama.NewConfig()
	client, err := sarama.NewConsumerGroup(brokers, group, cfg)
	if err != nil {
		return nil, err
	}

	return &kafkaConsumer{client: client}, nil
}

func (k *kafkaConsumer) Consume(ctx context.Context, topics []string) {
	for {
		if err := k.client.Consume(ctx, topics, k); err != nil {
			log.Printf("on consume: %v", err)
			time.Sleep(time.Second * 5)
		}
	}
}

func (k *kafkaConsumer) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (k *kafkaConsumer) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (k *kafkaConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case <-session.Context().Done():
			log.Print("Done")
			return nil
		case msg, ok := <-claim.Messages():
			if !ok {
				log.Print("Data channel closed")
				return nil
			}

			if err := k.handleMessage(session.Context(), msg); err != nil {
				log.Println(err)
			}

			session.MarkMessage(msg, "")
		}
	}
}

func (k *kafkaConsumer) handleMessage(ctx context.Context, msg *sarama.ConsumerMessage) error {
	switch msg.Topic {
	case topicAddRequestName:
		return k.AddTask(ctx, msg.Value)
	case topicDeleteRequestName:
		return k.DeleteTask(ctx, msg.Value)
	case topicUpdateRequestName:
		return k.UpdateTask(ctx, msg.Value)
	default:
		return errors.Errorf("unknown topic - %s", msg.Topic)
	}
}

func (k *kafkaConsumer) AddTask(ctx context.Context, data []byte) error {
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

func (k *kafkaConsumer) DeleteTask(ctx context.Context, data []byte) error {
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

func (k *kafkaConsumer) UpdateTask(ctx context.Context, data []byte) error {
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
