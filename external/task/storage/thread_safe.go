package storage

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"

	"gitlab.ozon.dev/Vanek623/task-manager-system/external/task/models"
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
func (s *threadSafe) Add(ctx context.Context, t *models.Task) (ID *uuid.UUID, err error) {
	if err = s.limiter.start(ctx); err != nil {
		return
	}

	defer s.limiter.end()

	s.mu.Lock()
	defer s.mu.Unlock()

	return s.impl.Add(ctx, t)
}

// Delete удаление задачи
func (s *threadSafe) Delete(ctx context.Context, ID *uuid.UUID) error {
	if err := s.limiter.start(ctx); err != nil {
		return err
	}

	defer s.limiter.end()

	s.mu.Lock()
	defer s.mu.Unlock()

	return s.impl.Delete(ctx, ID)
}

// List чтение списка задач
func (s *threadSafe) List(ctx context.Context, limit, offset uint64) ([]*models.Task, error) {
	if err := s.limiter.start(ctx); err != nil {
		return nil, err
	}

	defer s.limiter.end()

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.impl.List(ctx, limit, offset)
}

// Update обновление задачи
func (s *threadSafe) Update(ctx context.Context, t *models.Task) error {
	if err := s.limiter.start(ctx); err != nil {
		return err
	}

	defer s.limiter.end()

	s.mu.Lock()
	defer s.mu.Unlock()

	return s.impl.Update(ctx, t)
}

// Get чтение задачи
func (s *threadSafe) Get(ctx context.Context, ID *uuid.UUID) (*models.Task, error) {
	if err := s.limiter.start(ctx); err != nil {
		return nil, err
	}

	defer s.limiter.end()

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.impl.Get(ctx, ID)
}
