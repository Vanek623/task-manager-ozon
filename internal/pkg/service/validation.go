package service

import (
	"github.com/pkg/errors"
	"unicode/utf8"
)

const (
	maxNameLen        = 64
	maxDescriptionLen = 256
)

func isTitleOk(t string) error {
	if t == "" {
		return errors.Wrap(ErrValidation, "title is empty")
	}

	if utf8.RuneCountInString(t) > maxNameLen {
		return errors.Wrap(ErrValidation, "title too large")
	}

	return nil
}

func isDescriptionOk(d string) error {
	if utf8.RuneCountInString(d) > maxDescriptionLen {
		return errors.Wrap(ErrValidation, "description too large")
	}

	return nil
}

func isTitleAndDescriptionOk(title, description string) error {
	if err := isTitleOk(title); err != nil {
		return err
	}

	if err := isDescriptionOk(description); err != nil {
		return err
	}

	return nil
}
