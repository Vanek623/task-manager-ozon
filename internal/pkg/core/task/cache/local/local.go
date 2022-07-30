package local

import (
	"sync"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"
)

const poolSize = 10

// Cache структура локального кэша
type Cache struct {
	mu   sync.RWMutex
	data map[uint]models.Task
	pool chan struct{}
}

const maxTasks = 8

var (
	errTaskExist    = errors.New("task already exist")
	errTaskNotExist = errors.New("task doesn't exist")
)

// New Создание локального кэша
func New() Cache {
	return Cache{
		mu:   sync.RWMutex{},
		data: make(map[uint]models.Task),
		pool: make(chan struct{}, poolSize),
	}
}

// List чтение списка задач
func (c *Cache) List() []models.Task {
	c.rLock()
	defer c.rUnlock()

	res := make([]models.Task, 0, len(c.data))

	for _, t := range c.data {
		res = append(res, t)
	}

	return res
}

// Add добавление задачи
func (c *Cache) Add(t models.Task) error {
	c.lock()
	defer c.unlock()

	if len(c.data) >= maxTasks {
		return errors.New("Has no space for tasks, please delete one")
	}
	if _, ok := c.data[t.ID]; ok {
		return errors.Wrapf(errTaskExist, "ID: [%d]", t.ID)
	}

	c.data[t.ID] = t
	return nil
}

// Update обновление задачи
func (c *Cache) Update(t models.Task) error {
	c.lock()
	defer c.unlock()

	if _, ok := c.data[t.ID]; !ok {
		return errors.Wrapf(errTaskNotExist, "ID: [%d]", t.ID)
	}

	t.Created = c.data[t.ID].Created
	c.data[t.ID] = t

	return nil
}

// Delete удаление задачи
func (c *Cache) Delete(ID uint) error {
	c.lock()
	defer c.unlock()

	if _, ok := c.data[ID]; !ok {
		return errors.Wrapf(errTaskNotExist, "ID: [%d]", ID)
	}

	delete(c.data, ID)
	return nil
}

// Get чтение задачи
func (c *Cache) Get(ID uint) (models.Task, error) {
	c.rLock()
	defer c.rUnlock()

	if _, ok := c.data[ID]; !ok {
		return models.Task{}, errors.Wrapf(errTaskNotExist, "ID: [%d]", ID)
	}

	return c.data[ID], nil
}

func (c *Cache) decPool() {
	c.pool <- struct{}{}
}

func (c *Cache) incPool() {
	<-c.pool
}

func (c *Cache) lock() {
	c.mu.Lock()
	c.decPool()
}

func (c *Cache) unlock() {
	c.mu.Unlock()
	c.incPool()
}

func (c *Cache) rLock() {
	c.mu.RLock()
	c.decPool()
}

func (c *Cache) rUnlock() {
	c.mu.RUnlock()
	c.incPool()
}
