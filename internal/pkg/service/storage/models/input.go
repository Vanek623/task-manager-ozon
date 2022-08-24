package models

import (
	"github.com/google/uuid"
)

// AddTaskData новый запрос на создание задачи для хранилища
type AddTaskData struct {
	id                 *uuid.UUID
	title, description string
}

// Description описание задачи
func (a *AddTaskData) Description() string {
	return a.description
}

// ID Идентификатор задачи
func (a *AddTaskData) ID() uuid.UUID {
	return *a.id
}

// Title заголовок задачи
func (a *AddTaskData) Title() string {
	return a.title
}

// NewAddTaskData новый запрос на создание задачи для хранилища
func NewAddTaskData(ID *uuid.UUID, title, description string) *AddTaskData {
	return &AddTaskData{
		id:          ID,
		title:       title,
		description: description,
	}
}
