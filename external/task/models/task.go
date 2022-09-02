package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Task структура для хранения задачи
type Task struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Created     time.Time `json:"created" db:"created"`
	Edited      time.Time `json:"edited" db:"edited"`
}

// String формирование краткого описания
func (t *Task) String() string {
	return fmt.Sprintf("%d: %s", t.ID, t.Title)
}
