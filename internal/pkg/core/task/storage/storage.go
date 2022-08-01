package storage

import (
	"sync"
	"time"

	"github.com/pkg/errors"
)

var (
	// ErrTaskExist задача уже есть в хранилище
	ErrTaskExist = errors.New("task already exist")

	// ErrTaskNotExist задачи нет в хранилище
	ErrTaskNotExist = errors.New("task doesn't exist")

	// ErrActionTimeout таймаут операции
	ErrActionTimeout = errors.New("request timeout")
)

type storage struct {
	capacity int
	timeout  time.Duration
	mu       sync.RWMutex
	pool     chan struct{}
}

func (c *storage) lock() {
	c.mu.Lock()
	c.pool <- struct{}{}
}

func (c *storage) unlock() {
	c.mu.Unlock()
	<-c.pool
}

func (c *storage) rLock() {
	c.mu.RLock()
	c.pool <- struct{}{}
}

func (c *storage) rUnlock() {
	c.mu.RUnlock()
	<-c.pool
}
