package models

import (
	"fmt"
	"time"
)

// Task структура для хранения задачи
type Task struct {
	ID          uint64    `storage:"id"`
	Title       string    `storage:"title"`
	Description string    `storage:"description"`
	Created     time.Time `storage:"created"`
	Edited      time.Time `storage:"edited"`
}

func (t *Task) String() string {
	return fmt.Sprintf("%d: %s", t.ID, t.Title)
}
