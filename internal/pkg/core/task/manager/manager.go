package manager

import (
	"unicode/utf8"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"
)

type iTaskStorage interface {
	Add(t models.Task) error
	Delete(ID uint) error
	List() ([]models.Task, error)
	Update(t models.Task) error
	Get(ID uint) (models.Task, error)
}

// ErrValidation ошибка валидации данных
var ErrValidation = errors.New("invalid data")

const (
	maxNameLen        = 64
	maxDescriptionLen = 256
)

// checkTitleAndDescription проверка корректности заголовка и описания задачи
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
