package models

import (
	"fmt"
	"time"
)

// Task краткая информация о задаче
type Task struct {
	id    uint64
	title string
}

// ID ID задачи
func (t *Task) ID() uint64 {
	return t.id
}

// Title заголовок задачи
func (t *Task) Title() string {
	return t.title
}

// NewTask создание краткого описания задачи
func NewTask(ID uint64, title string) *Task {
	return &Task{
		id:    ID,
		title: title,
	}
}

// DetailedTask подробная информаци о задаче
type DetailedTask struct {
	title       string
	description string
	edited      time.Time
}

// Title заголовок задачи
func (t *DetailedTask) Title() string {
	return t.title
}

// Description описание задачи
func (t *DetailedTask) Description() string {
	return t.description
}

// Edited время изменения задачи
func (t *DetailedTask) Edited() time.Time {
	return t.edited
}

func (t *DetailedTask) String() string {
	return fmt.Sprintf("Title: %s \nDescription: %s \nEdited: %s",
		t.Title(),
		t.Description(),
		t.Edited().Format(time.Stamp))
}

// NewDetailedTask создание подробного описания задачи
func NewDetailedTask(title string, description string, edited time.Time) *DetailedTask {
	return &DetailedTask{
		title:       title,
		description: description,
		edited:      edited,
	}
}
