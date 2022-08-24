package storage

import (
	"context"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
)

type kafka struct {
}

func (k *kafka) Add(ctx context.Context, data *models.AddTaskData) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (k *kafka) Delete(ctx context.Context, data *models.DeleteTaskData) error {
	//TODO implement me
	panic("implement me")
}

func (k *kafka) List(ctx context.Context, data *models.ListTaskData) ([]*models.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (k *kafka) Update(ctx context.Context, data *models.UpdateTaskData) error {
	//TODO implement me
	panic("implement me")
}

func (k *kafka) Get(ctx context.Context, data *models.GetTaskData) (*models.DetailedTask, error) {
	//TODO implement me
	panic("implement me")
}
