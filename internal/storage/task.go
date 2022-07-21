package storage

import (
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/pkg/errors"
)

var lastID = uint(0)

// Task структура для хранения задачи
type Task struct {
	taskID      uint
	title       string
	description string
	created     time.Time
}

// Created чтение времени создания задачи
func (t *Task) Created() time.Time {
	return t.created
}

// ID чтение идентификатора задачи
func (t *Task) ID() uint {
	return t.taskID
}

// Title чтение названия задачи
func (t *Task) Title() string {
	return t.title
}

// Description чтение описания задачи
func (t *Task) Description() string {
	return t.description
}

const maxNameLen = 64

// SetTitle установка названия задачи
func (t *Task) SetTitle(title string) error {
	if title == "" {
		return errors.New("title must be not empty")
	}
	if utf8.RuneCountInString(title) > maxNameLen {
		return errors.Errorf("title must be less than %d", maxNameLen)
	}

	t.title = title

	return nil
}

const maxDescriptionLen = 256

// SetDescription установка описания задачи
func (t *Task) SetDescription(description string) error {
	if utf8.RuneCountInString(description) > maxDescriptionLen {
		return errors.Errorf("description must be less than %d", maxDescriptionLen)
	}

	t.description = description

	return nil
}

// NewTask Создание задачи
func NewTask(title, description string) (*Task, error) {
	t := Task{}
	if err := t.SetTitle(title); err != nil {
		return nil, err
	}
	if err := t.SetDescription(description); err != nil {
		return nil, err
	}
	t.created = time.Now()

	lastID++
	t.taskID = lastID
	return &t, nil
}

func (t Task) String() string {
	return fmt.Sprintf("%d: %s", t.taskID, t.title)
}
