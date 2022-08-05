package storage

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"
)

const localCapacity = 100

// local структура локального хранилища
type local struct {
	data   map[uint]models.Task
	lastID uint64
}

func newLocal() *local {
	return &local{
		data:   make(map[uint]models.Task),
		lastID: 1,
	}
}

func (s *local) Add(_ context.Context, t models.Task) (uint, error) {
	if len(s.data) >= localCapacity {
		return 0, ErrHasNoSpace
	}

	t.ID = uint(s.lastID)
	s.lastID++
	t.Created = time.Now()

	s.data[t.ID] = t

	return t.ID, nil
}

func (s *local) Delete(_ context.Context, ID uint) error {
	if _, ok := s.data[ID]; !ok {
		return errors.Wrapf(ErrTaskNotExist, "ID: [%d]", ID)
	}

	delete(s.data, ID)
	return nil
}

func (s *local) List(_ context.Context, limit, offset uint) ([]models.Task, error) {
	res := make([]models.Task, 0, limit)

	for i, t := range s.data {
		if i >= offset+limit {
			break
		}
		if i < offset {
			continue
		}

		res = append(res, t)
	}

	return res, nil
}

func (s *local) Update(_ context.Context, t models.Task) error {
	if _, ok := s.data[t.ID]; !ok {
		return errors.Wrapf(ErrTaskNotExist, "ID: [%d]", t.ID)
	}

	t.Created = s.data[t.ID].Created
	s.data[t.ID] = t

	return nil
}

func (s *local) Get(_ context.Context, ID uint) (*models.Task, error) {
	time.Sleep(time.Second)

	task, ok := s.data[ID]
	if !ok {
		return nil, errors.Wrapf(ErrTaskNotExist, "ID: [%d]", ID)
	}

	return &task, nil
}
