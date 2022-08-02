package storage

import (
	"context"
	"time"
	"unicode/utf8"

	"github.com/pkg/errors"

	"sync"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"
)

type iTaskStorage interface {
	Add(ctx context.Context, t models.Task) error
	Delete(ctx context.Context, ID uint) error
	List(ctx context.Context) ([]models.Task, error)
	Update(ctx context.Context, t models.Task) error
	Get(ctx context.Context, ID uint) (*models.Task, error)
}

var (
	// ErrTaskNotExist задачи нет в хранилище
	ErrTaskNotExist = errors.New("task doesn't exist")

	// ErrHasNoSpace отсутсвует место в хранилище
	ErrHasNoSpace = errors.New("Has no space for tasks, please delete one")

	// ErrValidation ошибка валидации данных
	ErrValidation = errors.New("invalid data")
)

const (
	maxNameLen        = 64
	maxDescriptionLen = 256
)

// Storage реализация многопоточного хранилища
type Storage struct {
	impl    iTaskStorage
	limiter workerLimiter
	mu      sync.RWMutex
}

// NewLocal создание локального многопоточного хранилища
func NewLocal() *Storage {
	return &Storage{
		impl:    newLocal(),
		limiter: newWorkerLimiter(10, 100*time.Millisecond),
	}
}

// Add добавление задачи
func (s *Storage) Add(ctx context.Context, t models.Task) error {
	if err := s.limiter.start(ctx); err != nil {
		return err
	}

	defer s.limiter.end()

	s.mu.Lock()
	defer s.mu.Unlock()

	return s.impl.Add(ctx, t)
}

// Delete удаление задачи
func (s *Storage) Delete(ctx context.Context, ID uint) error {
	if err := s.limiter.start(ctx); err != nil {
		return err
	}

	defer s.limiter.end()

	s.mu.Lock()
	defer s.mu.Unlock()

	return s.impl.Delete(ctx, ID)
}

// List чтение списка задач
func (s *Storage) List(ctx context.Context) ([]models.Task, error) {
	if err := s.limiter.start(ctx); err != nil {
		return nil, err
	}

	defer s.limiter.end()

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.impl.List(ctx)
}

// Update обновление задачи
func (s *Storage) Update(ctx context.Context, t models.Task) error {
	if err := s.limiter.start(ctx); err != nil {
		return err
	}

	defer s.limiter.end()

	s.mu.Lock()
	defer s.mu.Unlock()

	return s.impl.Update(ctx, t)
}

// Get чтение задачи
func (s *Storage) Get(ctx context.Context, ID uint) (*models.Task, error) {
	if err := s.limiter.start(ctx); err != nil {
		return nil, err
	}

	defer s.limiter.end()

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.impl.Get(ctx, ID)
}

func checkTitleAndDescription(t models.Task) error {
	if t.Title == "" {
		return errors.Wrap(ErrValidation, "field: [title] is empty")
	}
	if utf8.RuneCountInString(t.Title) > maxNameLen {
		return errors.Wrap(ErrValidation, "field: [title] too large")
	}

	if utf8.RuneCountInString(t.Description) > maxDescriptionLen {
		return errors.Wrap(ErrValidation, "field: [description] too large")
	}

	return nil
}
