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

// String формирование краткого описания
func (t *Task) String() string {
	return fmt.Sprintf("%d: %s", t.ID, t.Title)
}

//// MarshalBinary создание бинарного представления
//func (t Task) MarshalBinary() ([]byte, error) {
//	return json.Marshal(t)
//}
//
//// UnmarshalBinary создание задачи из бинарного представления
//func (t *Task) UnmarshalBinary(data []byte) error {
//	return json.Unmarshal(data, t)
//}
