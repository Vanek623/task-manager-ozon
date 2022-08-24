package storage

import (
	"context"
	"time"

	"github.com/google/uuid"

	"gitlab.ozon.dev/Vanek623/task-manager-system/external/task/models"

	"github.com/pkg/errors"
)

const localCapacity = 100

// local структура локального хранилища
type local struct {
	data map[uuid.UUID]*models.Task
}

func newLocal() *local {
	return &local{
		data: make(map[uuid.UUID]*models.Task),
	}
}

func (s *local) Add(_ context.Context, t *models.Task) (*uuid.UUID, error) {
	if len(s.data) >= localCapacity {
		return nil, ErrHasNoSpace
	}

	t.Created = time.Now()

	s.data[t.ID] = t

	return &t.ID, nil
}

func (s *local) Delete(_ context.Context, ID *uuid.UUID) error {
	if _, ok := s.data[*ID]; !ok {
		return errors.Wrapf(ErrTaskNotExist, "ID: [%d]", ID)
	}

	delete(s.data, *ID)
	return nil
}

func (s *local) List(_ context.Context, limit, offset uint64) ([]*models.Task, error) {
	res := make([]*models.Task, 0, limit)

	count := uint64(0)
	for _, t := range s.data {
		if count >= offset+limit {
			break
		}
		if count < offset {
			continue
		}

		res = append(res, t)
		count++
	}

	return res, nil
}

func (s *local) Update(_ context.Context, t *models.Task) error {
	if _, ok := s.data[t.ID]; !ok {
		return errors.Wrapf(ErrTaskNotExist, "ID: [%d]", t.ID)
	}

	t.Created = s.data[t.ID].Created
	s.data[t.ID] = t

	return nil
}

func (s *local) Get(_ context.Context, ID *uuid.UUID) (*models.Task, error) {
	time.Sleep(time.Second)

	task, ok := s.data[*ID]
	if !ok {
		return nil, errors.Wrapf(ErrTaskNotExist, "ID: [%d]", ID)
	}

	return task, nil
}
