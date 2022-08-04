package storage

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"
)

const localCapacity = 8

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

	if err := checkTitleAndDescription(t); err != nil {
		return 0, err
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

func (s *local) List(_ context.Context) ([]models.Task, error) {
	res := make([]models.Task, 0, len(s.data))

	for _, t := range s.data {
		res = append(res, t)
	}

	return res, nil
}

func (s *local) Update(_ context.Context, t models.Task) error {
	if _, ok := s.data[t.ID]; !ok {
		return errors.Wrapf(ErrTaskNotExist, "ID: [%d]", t.ID)
	}

	if err := checkTitleAndDescription(t); err != nil {
		return err
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
