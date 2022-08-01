package manager

import (
	"sync/atomic"
	"time"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"
	storagePkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/storage"
)

// LocalManager Реализация интерфейса управления задачами
type LocalManager struct {
	storage iTaskStorage
	lastID  uint64
}

// NewLocal Создание модуля управления задачами
func NewLocal() *LocalManager {
	return &LocalManager{
		storage: storagePkg.NewLocal(),
	}
}

// Add создание задачи
func (c *LocalManager) Add(t models.Task) error {
	if err := checkTitleAndDescription(t); err != nil {
		return err
	}

	t.ID = uint(atomic.AddUint64(&c.lastID, 1))
	t.Created = time.Now()

	return c.storage.Add(t)
}

// Update обновление задачи
func (c *LocalManager) Update(t models.Task) error {
	if err := checkTitleAndDescription(t); err != nil {
		return err
	}

	return c.storage.Update(t)
}

// Delete удаление задачи
func (c *LocalManager) Delete(ID uint) error {
	return c.storage.Delete(ID)
}

// List получение списка задач
func (c *LocalManager) List() ([]models.Task, error) {
	return c.storage.List()
}

// Get получение задачи
func (c *LocalManager) Get(ID uint) (models.Task, error) {
	return c.storage.Get(ID)
}
