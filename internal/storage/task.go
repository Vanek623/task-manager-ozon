package storage

import (
	"errors"
	"fmt"
)

var lastId = uint(0)

type Task struct {
	id          uint
	name        string
	description string
}

func (t *Task) Id() uint {
	return t.id
}

func (t *Task) Name() string {
	return t.name
}

func (t *Task) Description() string {
	return t.description
}

var maxNameLen = 64

func (t *Task) SetName(name string) error {
	if len(name) == 0 {
		return errors.New("name must be not empty")
	} else if len(name) > maxNameLen {
		return errors.New(fmt.Sprintf("Name must be less than %d", maxNameLen))
	}

	t.name = name

	return nil
}

func (t *Task) SetDescription(description string) error {
	if len(description) > 256 {
		return errors.New("description too large")
	}

	t.description = description

	return nil
}

func NewTask(name, description string) (*Task, error) {
	t := Task{}
	if err := t.SetName(name); err != nil {
		return nil, err
	}
	if err := t.SetDescription(description); err != nil {
		return nil, err
	}
	lastId++
	t.id = lastId
	return &t, nil
}

func (t Task) String() string {
	if len(t.description) == 0 {
		return fmt.Sprintf("%d: %s", t.id, t.name)
	} else {
		return fmt.Sprintf("%d: %s - %s", t.id, t.name, t.description)
	}
}
