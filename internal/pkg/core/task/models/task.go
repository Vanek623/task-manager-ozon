package models

import (
	"fmt"
	"time"
)

// Task структура для хранения задачи
type Task struct {
	ID          uint
	Title       string
	Description string
	Created     time.Time
}

func (t *Task) String() string {
	return fmt.Sprintf("%d: %s", t.ID, t.Title)
}
