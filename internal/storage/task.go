package storage

import (
	"errors"
	"fmt"
	"time"
)

var lastId = uint(0)

type Task struct {
	id          uint
	title       string
	description string
	created     time.Time
}

func (t *Task) Created() time.Time {
	return t.created
}

func (t *Task) Id() uint {
	return t.id
}

func (t *Task) Title() string {
	return t.title
}

func (t *Task) Description() string {
	return t.description
}

const maxNameLen = 64

func (t *Task) SetTitle(title string) error {
	if len(title) == 0 {
		return errors.New("title must be not empty")
	} else if len(title) > maxNameLen {
		return errors.New(fmt.Sprintf("Title must be less than %d", maxNameLen))
	}

	t.title = title

	return nil
}

func (t *Task) SetDescription(description string) error {
	if len(description) > 256 {
		return errors.New("description too large")
	}

	t.description = description

	return nil
}

func NewTask(title, description string) (*Task, error) {
	t := Task{}
	if err := t.SetTitle(title); err != nil {
		return nil, err
	}
	if err := t.SetDescription(description); err != nil {
		return nil, err
	}
	t.created = time.Now()

	lastId++
	t.id = lastId
	return &t, nil
}

func (t Task) String() string {
	return fmt.Sprintf("%d: %s", t.id, t.title)
}
