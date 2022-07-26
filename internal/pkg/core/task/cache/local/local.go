package local

import (
	"github.com/pkg/errors"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"
)

// Cache структура локального кэша
type Cache struct {
	data map[uint]models.Task
}

const maxTasks = 8

var (
	ErrTaskExist    = errors.New("task already exist")
	ErrTaskNotExist = errors.New("task doesn't exist")
)

func New() *Cache {
	return &Cache{data: make(map[uint]models.Task)}
}

// List чтение списка задач
func (c *Cache) List() []models.Task {
	res := make([]models.Task, 0, len(c.data))

	for _, t := range c.data {
		res = append(res, t)
	}

	return res
}

// Add добавление задачи
func (c *Cache) Add(t models.Task) error {
	if len(c.data) >= maxTasks {
		return errors.New("Has no space for tasks, please delete one")
	}
	if _, ok := c.data[t.ID]; ok {
		return errors.Wrapf(ErrTaskExist, "ID: [%d]", t.ID)
	}

	c.data[t.ID] = t
	return nil
}

// Update обновление задачи
func (c *Cache) Update(t models.Task) error {
	if _, ok := c.data[t.ID]; !ok {
		return errors.Wrapf(ErrTaskNotExist, "ID: [%d]", t.ID)
	}

	c.data[t.ID] = t
	return nil
}

// Delete удаление задачи
func (c *Cache) Delete(ID uint) error {
	if _, ok := c.data[ID]; !ok {
		return errors.Wrapf(ErrTaskNotExist, "ID: [%d]", ID)
	}

	delete(c.data, ID)
	return nil
}

// Get чтение задачи
func (c *Cache) Get(ID uint) (models.Task, error) {
	if _, ok := c.data[ID]; !ok {
		return models.Task{}, errors.Wrapf(ErrTaskNotExist, "ID: [%d}", ID)
	}

	return c.data[ID], nil
}
