package cache

import (
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"
)

type ICache interface {
	Add(t models.Task) error
	Delete(ID uint) error
	List() []models.Task
	Update(t models.Task) error
	Get(ID uint) (models.Task, error)
}
