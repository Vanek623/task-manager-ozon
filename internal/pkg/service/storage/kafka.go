package storage

import (
	"context"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/counters"
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
	cs       *counters.Counters
}

func newKafka(ctx context.Context, brokers []string, syncStorage iStorage, cs *counters.Counters) (*kafka, error) {
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
		cs:       cs,
	}

	<-k.ctx.Done()

	go func() {
		if err := producer.Close(); err != nil {
			log.Error(err)
		} else {
			log.Info("producer closed")
		}
	}()

	return k, nil
}

func (k *kafka) send(ctx context.Context, obj []byte, topicName string) error {
	k.cs.Inc(counters.Outbound)

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

func (k *kafka) Add(ctx context.Context, data *storageModelsPkg.AddTaskData) error {
	obj, err := data.MarshalJSON()

	if err != nil {
		return err
	}

	err = k.send(ctx, obj, topicAddRequestName)

	return err
}

func (k *kafka) Delete(ctx context.Context, data *models.DeleteTaskData) error {
	obj, err := data.MarshalJSON()

	if err != nil {
		return err
	}

	err = k.send(ctx, obj, topicDeleteRequestName)

	return err
}

func (k *kafka) Update(ctx context.Context, data *models.UpdateTaskData) error {
	obj, err := data.MarshalJSON()

	if err != nil {
		return err
	}

	err = k.send(ctx, obj, topicUpdateRequestName)

	return err
}
