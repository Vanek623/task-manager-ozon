package async

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Shopify/sarama"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/counters"
	serviceModelsPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
	storageModelsPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/storage/models"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/tracer"
	"go.opentelemetry.io/otel"
)

const (
	topicAddRequestName    = "income_add_request"
	topicDeleteRequestName = "income_delete_request"
	topicUpdateRequestName = "income_update_request"
	topicGetRequestName    = "income_get_request"
	topicListRequestName   = "income_list_request"

	timeout = 100 * time.Millisecond
)

type kafka struct {
	producer sarama.SyncProducer
	cs       *counters.Counters
}

func newKafka(ctx context.Context, brokers []string, cs *counters.Counters) (*kafka, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, cfg)
	if err != nil {
		return nil, err
	}

	go func() {
		<-ctx.Done()
		if err := producer.Close(); err != nil {
			log.Error(err)
		} else {
			log.Info("Kafka down")
		}
	}()

	return &kafka{
		producer: producer,
		cs:       cs,
	}, nil
}

func (k *kafka) send(ctx context.Context, obj []byte, topicName string) error {
	k.cs.Inc(counters.Outbound)

	ctx, cl := context.WithTimeout(ctx, timeout)
	defer cl()

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
		log.Error("Request timeout")
		return nil
	}
}

func (k *kafka) WriteAddRequest(ctx context.Context, data *storageModelsPkg.AddTaskData) (*uuid.UUID, error) {
	requestID := uuid.New()

	tmp := struct {
		RequestID          uuid.UUID
		ID                 uuid.UUID
		Title, Description string
	}{
		RequestID:   requestID,
		ID:          data.ID(),
		Title:       data.Title(),
		Description: data.Description(),
	}

	obj, err := json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	_, span := otel.Tracer(tracer.Name).Start(ctx, tracer.MakeSpanName("Add Kafka"))
	defer span.End()

	if err = k.send(ctx, obj, topicAddRequestName); err != nil {
		return nil, err
	}

	return &requestID, err
}

func (k *kafka) WriteDeleteRequest(ctx context.Context, data *serviceModelsPkg.DeleteTaskData) (*uuid.UUID, error) {
	requestID := uuid.New()

	tmp := struct {
		RequestID uuid.UUID
		ID        uuid.UUID
	}{
		RequestID: requestID,
		ID:        data.ID(),
	}

	obj, err := json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	_, span := otel.Tracer(tracer.Name).Start(ctx, tracer.MakeSpanName("Delete Kafka"))
	defer span.End()

	if err = k.send(ctx, obj, topicDeleteRequestName); err != nil {
		return nil, err
	}

	return &requestID, nil
}

func (k *kafka) WriteUpdateRequest(ctx context.Context, data *serviceModelsPkg.UpdateTaskData) (*uuid.UUID, error) {
	requestID := uuid.New()

	tmp := struct {
		RequestID          uuid.UUID
		ID                 uuid.UUID
		Title, Description string
	}{
		RequestID:   requestID,
		ID:          data.ID(),
		Title:       data.Title(),
		Description: data.Description(),
	}

	obj, err := json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	if err = k.send(ctx, obj, topicUpdateRequestName); err != nil {
		return nil, err
	}

	return &requestID, nil
}

func (k *kafka) WriteListRequest(ctx context.Context, data *serviceModelsPkg.ListTaskData) (*uuid.UUID, error) {
	requestID := uuid.New()

	tmp := struct {
		RequestID     uuid.UUID
		Limit, Offset uint64
	}{
		RequestID: requestID,
		Limit:     data.Limit(),
		Offset:    data.Offset(),
	}

	obj, err := json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	if err = k.send(ctx, obj, topicListRequestName); err != nil {
		return nil, err
	}

	return &requestID, nil
}

func (k *kafka) WriteGetRequest(ctx context.Context, data *serviceModelsPkg.GetTaskData) (*uuid.UUID, error) {
	requestID := uuid.New()

	tmp := struct {
		RequestID uuid.UUID
		ID        uuid.UUID
	}{
		RequestID: requestID,
		ID:        data.ID(),
	}

	obj, err := json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	_, span := otel.Tracer(tracer.Name).Start(ctx, tracer.MakeSpanName("Get Kafka"))
	defer span.End()

	if err = k.send(ctx, obj, topicGetRequestName); err != nil {
		return nil, err
	}

	return &requestID, nil
}
