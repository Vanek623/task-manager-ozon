package storage

import (
	"context"
	"log"

	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
	storageModelsPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/storage/models"
)

const (
	topicAddRequestName    = "income_add_request"
	topicDeleteRequestName = "income_delete_request"
	topicUpdateRequestName = "income_update_request"
)

type kafka struct {
	iStorage
	producer sarama.SyncProducer
	ctx      context.Context
}

func newKafka(ctx context.Context, brokers []string, syncStorage iStorage) (*kafka, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, cfg)
	if err != nil {
		return nil, err
	}

	k := &kafka{
		iStorage: syncStorage,
		producer: producer,
		ctx:      ctx,
	}

	go k.closeWhenCtxDone()

	return k, nil
}

func (k *kafka) send(ctx context.Context, obj []byte, topicName string) error {
	ch := make(chan error)
	go func() {
		var err error
		_, _, err = k.producer.SendMessage(&sarama.ProducerMessage{
			Topic: topicName,
			Value: sarama.ByteEncoder(obj),
		})

		ch <- err
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (k *kafka) closeWhenCtxDone() {
	<-k.ctx.Done()

	if err := k.producer.Close(); err != nil {
		log.Println(err)
	} else {
		log.Println("producer closed")
	}
}

func (k *kafka) Add(ctx context.Context, data *storageModelsPkg.AddTaskData) error {
	obj, err := data.MarshalJSON()
	if err != nil {
		return err
	}

	return k.send(ctx, obj, topicAddRequestName)
}

func (k *kafka) Delete(ctx context.Context, data *models.DeleteTaskData) error {
	obj, err := data.MarshalJSON()
	if err != nil {
		return err
	}

	return k.send(ctx, obj, topicDeleteRequestName)
}

func (k *kafka) Update(ctx context.Context, data *models.UpdateTaskData) error {
	obj, err := data.MarshalJSON()
	if err != nil {
		return err
	}

	return k.send(ctx, obj, topicUpdateRequestName)
}
