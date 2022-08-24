package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Task структура для хранения задачи
type Task struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Created     time.Time `db:"created"`
	Edited      time.Time `db:"edited"`
}

func (t *Task) String() string {
	return fmt.Sprintf("%d: %s", t.ID, t.Title)
}
