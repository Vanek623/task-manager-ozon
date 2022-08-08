package storage

import (
	"context"
	"sync"
	"time"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"
)

type threadSafe struct {
	impl    iTaskStorage
	limiter workerLimiter
	mu      sync.RWMutex
}

func newThreadSafeStorage(storage iTaskStorage, maxWorkers uint, timeout time.Duration) *threadSafe {
	return &threadSafe{
		impl:    storage,
		limiter: newWorkerLimiter(maxWorkers, timeout),
	}
}

// Add добавление задачи
func (s *threadSafe) Add(ctx context.Context, t models.Task) (ID uint, err error) {
	if err = s.limiter.start(ctx); err != nil {
		return
	}

	defer s.limiter.end()

	s.mu.Lock()
	defer s.mu.Unlock()

	return s.impl.Add(ctx, t)
}

// Delete удаление задачи
func (s *threadSafe) Delete(ctx context.Context, ID uint) error {
	if err := s.limiter.start(ctx); err != nil {
		return err
	}

	defer s.limiter.end()

	s.mu.Lock()
	defer s.mu.Unlock()

	return s.impl.Delete(ctx, ID)
}

// List чтение списка задач
func (s *threadSafe) List(ctx context.Context, limit, offset uint) ([]models.Task, error) {
	if err := s.limiter.start(ctx); err != nil {
		return nil, err
	}

	defer s.limiter.end()

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.impl.List(ctx, limit, offset)
}

// Update обновление задачи
func (s *threadSafe) Update(ctx context.Context, t models.Task) error {
	if err := s.limiter.start(ctx); err != nil {
		return err
	}

	defer s.limiter.end()

	s.mu.Lock()
	defer s.mu.Unlock()

	return s.impl.Update(ctx, t)
}

// Get чтение задачи
func (s *threadSafe) Get(ctx context.Context, ID uint) (*models.Task, error) {
	if err := s.limiter.start(ctx); err != nil {
		return nil, err
	}

	defer s.limiter.end()

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.impl.Get(ctx, ID)
}
