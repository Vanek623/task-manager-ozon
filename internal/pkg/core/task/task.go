package task

import (
	"github.com/pkg/errors"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/cache"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/cache/local"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"
	"unicode/utf8"
)

type ITask interface {
	Create(task *models.Task) error
	Update(task *models.Task) error
	Delete(ID uint) error
	List() []*models.Task
	Get(ID uint)
}

type Core struct {
	cache cache.ICache
}

// New Создание задачи
func New() *Core {
	return &Core{
		cache: local.New(),
	}
}

const (
	maxNameLen        = 64
	maxDescriptionLen = 256
)

var ErrValidation = errors.New("invalid data")

func checkTitleAndDescription(t models.Task) error {
	if t.Title == "" {
		return errors.Wrap(ErrValidation, "field: [title] is empty")
	}
	if utf8.RuneCountInString(t.Title) > maxNameLen {
		return errors.Wrap(ErrValidation, "field: [title] too large")
	}

	if t.Description == "" {
		return errors.Wrap(ErrValidation, "field: [description] is empty")
	}
	if utf8.RuneCountInString(t.Description) > maxDescriptionLen {
		return errors.Wrap(ErrValidation, "field: [description] too large")
	}

	return nil
}

var lastID uint = 0

func (c *Core) Create(t models.Task) error {
	t.ID = lastID
	lastID++

	if err := checkTitleAndDescription(t); err != nil {
		return err
	}

	return c.cache.Add(t)
}

func (c *Core) Update(t models.Task) error {
	if err := checkTitleAndDescription(t); err != nil {
		return err
	}

	return c.cache.Update(t)
}

func (c *Core) Delete(ID uint) error {
	return c.cache.Delete(ID)
}

func (c *Core) List() []models.Task {
	return c.cache.List()
}
