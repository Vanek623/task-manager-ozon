package storage

import (
	"github.com/pkg/errors"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"
	"sync"
)

type iTaskStorage interface {
	Add(t models.Task) error
	Delete(ID uint) error
	List() []models.Task
	Update(t models.Task) error
	Get(ID uint) (models.Task, error)
}

var (
	// ErrTaskExist задача уже есть в хранилище
	ErrTaskExist = errors.New("task already exist")

	// ErrTaskNotExist задачи нет в хранилище
	ErrTaskNotExist = errors.New("task doesn't exist")
)

type storage struct {
	capacity int
	mu       sync.RWMutex
	pool     chan struct{}
}

func (c *storage) decPool() {
	c.pool <- struct{}{}
}

func (c *storage) incPool() {
	<-c.pool
}

func (c *storage) lock() {
	c.mu.Lock()
	c.decPool()
}

func (c *storage) unlock() {
	c.mu.Unlock()
	c.incPool()
}

func (c *storage) rLock() {
	c.mu.RLock()
	c.decPool()
}

func (c *storage) rUnlock() {
	c.mu.RUnlock()
	c.incPool()
}
