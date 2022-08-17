package storage

import (
	"context"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/task/models"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/pkg/errors"
)

type iTaskStorage interface {
	Add(ctx context.Context, t *models.Task) (uint64, error)
	Delete(ctx context.Context, ID uint64) error
	List(ctx context.Context, limit, offset uint64) ([]*models.Task, error)
	Update(ctx context.Context, t *models.Task) error
	Get(ctx context.Context, ID uint64) (*models.Task, error)
}

var (
	// ErrTaskNotExist задачи нет в хранилище
	ErrTaskNotExist = errors.New("task doesn't exist")

	// ErrHasNoSpace отсутсвует место в хранилище
	ErrHasNoSpace = errors.New("Has no space for tasks, please delete one")
)

const (
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
