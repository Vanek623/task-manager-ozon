package validation

import (
	"unicode/utf8"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/sync"
)

const (
	maxNameLen        = 64
	maxDescriptionLen = 256
)

// IsTitleOk проверка правильности заголовка
func IsTitleOk(t string) error {
	if t == "" {
		return errors.Wrap(sync.ErrValidation, "title is empty")
	}

	if utf8.RuneCountInString(t) > maxNameLen {
		return errors.Wrap(sync.ErrValidation, "title too large")
	}

	return nil
}

// IsDescriptionOk проверка правильности описания
func IsDescriptionOk(d string) error {
	if utf8.RuneCountInString(d) > maxDescriptionLen {
		return errors.Wrap(sync.ErrValidation, "description too large")
	}

	return nil
}

// IsTitleAndDescriptionOk проверка правильности заголовка и описания
func IsTitleAndDescriptionOk(title, description string) error {
	if err := IsTitleOk(title); err != nil {
		return err
	}

	if err := IsDescriptionOk(description); err != nil {
		return err
	}

	return nil
}
