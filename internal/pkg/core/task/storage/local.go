package storage

import (
	"github.com/pkg/errors"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"
)

const localPoolSize = 10
const localCapacity = 8

// Local структура локального кэша
type Local struct {
	storage
	data map[uint]models.Task
}

// New Создание локального хранилища
func New() *Local {
	return &Local{
		storage: storage{
			capacity: localCapacity,
			pool:     make(chan struct{}, localPoolSize),
		},
		data: make(map[uint]models.Task),
	}
}

// List чтение списка задач
func (c *Local) List() []models.Task {
	c.rLock()
	defer c.rUnlock()

	res := make([]models.Task, 0, len(c.data))

	for _, t := range c.data {
		res = append(res, t)
	}

	return res
}

// Add добавление задачи
func (c *Local) Add(t models.Task) error {
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

// Update обновление задачи
func (c *Local) Update(t models.Task) error {
	c.lock()
	defer c.unlock()

	if _, ok := c.data[t.ID]; !ok {
		return errors.Wrapf(ErrTaskNotExist, "ID: [%d]", t.ID)
	}

	t.Created = c.data[t.ID].Created
	c.data[t.ID] = t

	return nil
}

// Delete удаление задачи
func (c *Local) Delete(ID uint) error {
	c.lock()
	defer c.unlock()

	if _, ok := c.data[ID]; !ok {
		return errors.Wrapf(ErrTaskNotExist, "ID: [%d]", ID)
	}

	delete(c.data, ID)
	return nil
}

// Get чтение задачи
func (c *Local) Get(ID uint) (models.Task, error) {
	c.rLock()
	defer c.rUnlock()

	if _, ok := c.data[ID]; !ok {
		return models.Task{}, errors.Wrapf(ErrTaskNotExist, "ID: [%d]", ID)
	}

	return c.data[ID], nil
}
