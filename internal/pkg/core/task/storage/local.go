package storage

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"
)

const localPoolSize = 10
const localCapacity = 8
const localTimeout = 100 * time.Millisecond

// Local структура локального кэша
type Local struct {
	storage
	data map[uint]models.Task
}

// NewLocal Создание локального хранилища
func NewLocal() *Local {
	return &Local{
		storage: storage{
			capacity: localCapacity,
			timeout:  localTimeout,
			pool:     make(chan struct{}, localPoolSize),
		},
		data: make(map[uint]models.Task),
	}
}

// List чтение списка задач
func (c *Local) List() ([]models.Task, error) {
	proc := make(chan struct{}, 1)
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	var tmp []models.Task
	go func() {
		tmp = c.list()
		proc <- struct{}{}
	}()

	for {
		select {
		case <-ctx.Done():
			return nil, ErrActionTimeout
		case <-proc:
			return tmp, nil
		}
	}
}

// Add добавление задачи
func (c *Local) Add(t models.Task) error {
	proc := make(chan struct{}, 1)
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	var err error
	go func() {
		//time.Sleep(time.Second)
		err = c.add(t)
		proc <- struct{}{}
	}()

	for {
		select {
		case <-ctx.Done():
			return ErrActionTimeout
		case <-proc:
			return err
		}
	}
}

// Update обновление задачи
func (c *Local) Update(t models.Task) error {
	proc := make(chan struct{}, 1)
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	var err error
	go func() {
		err = c.update(t)
		proc <- struct{}{}
	}()

	for {
		select {
		case <-ctx.Done():
			return ErrActionTimeout
		case <-proc:
			return err
		}
	}
}

// Delete удаление задачи
func (c *Local) Delete(ID uint) error {
	proc := make(chan struct{}, 1)
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	var err error
	go func() {
		err = c.delete(ID)
		proc <- struct{}{}
	}()

	for {
		select {
		case <-ctx.Done():
			return ErrActionTimeout
		case <-proc:
			return err
		}
	}
}

// Get чтение задачи
func (c *Local) Get(ID uint) (models.Task, error) {
	proc := make(chan struct{}, 1)
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	var err error
	var task models.Task
	go func() {
		task, err = c.get(ID)
		time.Sleep(time.Second)
		proc <- struct{}{}
	}()

	for {
		select {
		case <-ctx.Done():
			return task, ErrActionTimeout
		case <-proc:
			return task, err
		}
	}
}

func (c *Local) list() []models.Task {
	c.rLock()
	defer c.rUnlock()

	res := make([]models.Task, 0, len(c.data))

	for _, t := range c.data {
		res = append(res, t)
	}

	return res
}

func (c *Local) add(t models.Task) error {
	c.lock()
	defer c.unlock()

	if len(c.data) >= c.capacity {
		return errors.New("Has no space for tasks, please delete one")
	}
	if _, ok := c.data[t.ID]; ok {
		return errors.Wrapf(ErrTaskExist, "ID: [%d]", t.ID)
	}

	c.data[t.ID] = t
	return nil
}

func (c *Local) update(t models.Task) error {
	c.lock()
	defer c.unlock()

	if _, ok := c.data[t.ID]; !ok {
		return errors.Wrapf(ErrTaskNotExist, "ID: [%d]", t.ID)
	}

	t.Created = c.data[t.ID].Created
	c.data[t.ID] = t

	return nil
}

func (c *Local) delete(ID uint) error {
	c.lock()
	defer c.unlock()

	if _, ok := c.data[ID]; !ok {
		return errors.Wrapf(ErrTaskNotExist, "ID: [%d]", ID)
	}

	delete(c.data, ID)
	return nil
}

func (c *Local) get(ID uint) (models.Task, error) {
	c.rLock()
	defer c.rUnlock()

	if _, ok := c.data[ID]; !ok {
		return models.Task{}, errors.Wrapf(ErrTaskNotExist, "ID: [%d]", ID)
	}

	return c.data[ID], nil
}
