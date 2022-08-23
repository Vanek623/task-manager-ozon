package storage

import (
	"context"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/task/models"
)

type kafka struct {
}

func (k kafka) Add(ctx context.Context, t *models.Task) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (k kafka) Delete(ctx context.Context, ID uint64) error {
	//TODO implement me
	panic("implement me")
}

func (k kafka) List(ctx context.Context, limit, offset uint64) ([]*models.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (k kafka) Update(ctx context.Context, t *models.Task) error {
	//TODO implement me
	panic("implement me")
}

func (k kafka) Get(ctx context.Context, ID uint64) (*models.Task, error) {
	//TODO implement me
	panic("implement me")
}
