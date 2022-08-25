package models

import (
	"encoding/json"

	"github.com/google/uuid"
)

// AddTaskData новый запрос на создание задачи для хранилища
type AddTaskData struct {
	id                 *uuid.UUID
	title, description string
}

// Description описание задачи
func (d *AddTaskData) Description() string {
	return d.description
}

// ID Идентификатор задачи
func (d *AddTaskData) ID() uuid.UUID {
	return *d.id
}

// Title заголовок задачи
func (d *AddTaskData) Title() string {
	return d.title
}

// MarshalJSON создать json из структуры
func (d *AddTaskData) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		ID          *uuid.UUID
		Title       string
		Description string
	}{
		ID:          d.id,
		Title:       d.title,
		Description: d.description,
	})

	if err != nil {
		return nil, err
	}

	return j, nil
}

// NewAddTaskData новый запрос на создание задачи для хранилища
func NewAddTaskData(ID *uuid.UUID, title, description string) *AddTaskData {
	return &AddTaskData{
		id:          ID,
		title:       title,
		description: description,
	}
}
