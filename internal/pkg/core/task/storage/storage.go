package storage

import (
	"context"
	"time"
	"unicode/utf8"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/pkg/errors"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"
)

type iTaskStorage interface {
	Add(ctx context.Context, t models.Task) (uint, error)
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
	maxWorkers        = 10
	workerIdleTimeout = 100 * time.Millisecond
)

// Storage реализация хранилища
type Storage struct {
	iTaskStorage
}

// NewLocal создание локального многопоточного хранилища
func NewLocal(isThreadSafe bool) *Storage {
	if isThreadSafe {
		return &Storage{newThreadSafeStorage(newLocal(), maxWorkers, workerIdleTimeout)}
	}

	return &Storage{newLocal()}
}

// NewPostgres создание хранилища PostgreSQL
func NewPostgres(pool *pgxpool.Pool, isThreadSafe bool) *Storage {
	if isThreadSafe {
		return &Storage{newThreadSafeStorage(&postgres{pool: pool}, maxWorkers, workerIdleTimeout)}
	}

	return &Storage{&postgres{pool: pool}}
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
