package task

import (
	"time"
	"unicode/utf8"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/cache"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/cache/local"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"
)

// IManager интерфейс управления задачами
type IManager interface {
	Create(task models.Task) error
	Update(task models.Task) error
	Delete(ID uint) error
	List() []models.Task
	Get(ID uint) (models.Task, error)
}

// Manager Реализация интерфейса управления задачами
type Manager struct {
	cache  cache.ICache
	lastID uint
}

// NewLocalManager Создание модуля управления задачами
func NewLocalManager() *Manager {
	tmp := local.New()

	return &Manager{
		cache:  &tmp,
		lastID: 1,
	}
}

const (
	maxNameLen        = 64
	maxDescriptionLen = 256
)

var errValidation = errors.New("invalid data")

func checkTitleAndDescription(t models.Task) error {
	if t.Title == "" {
		return errors.Wrap(errValidation, "field: [title] is empty")
	}
	if utf8.RuneCountInString(t.Title) > maxNameLen {
		return errors.Wrap(errValidation, "field: [title] too large")
	}

	if utf8.RuneCountInString(t.Description) > maxDescriptionLen {
		return errors.Wrap(errValidation, "field: [description] too large")
	}

	return nil
}

// Create создание задачи
func (c *Manager) Create(t models.Task) error {
	t.ID = c.lastID
	c.lastID++
	t.Created = time.Now()

	if err := checkTitleAndDescription(t); err != nil {
		return err
	}

	return c.cache.Add(t)
}

// Update обновление задачи
func (c *Manager) Update(t models.Task) error {
	if err := checkTitleAndDescription(t); err != nil {
		return err
	}

	return c.cache.Update(t)
}

// Delete удаление задачи
func (c *Manager) Delete(ID uint) error {
	return c.cache.Delete(ID)
}

// List получение списка задач
func (c *Manager) List() []models.Task {
	return c.cache.List()
}

// Get получение задачи
func (c *Manager) Get(ID uint) (models.Task, error) {
	return c.cache.Get(ID)
}
