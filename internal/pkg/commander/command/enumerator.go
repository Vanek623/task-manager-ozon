package command

import (
	"github.com/google/uuid"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
)

type taskEnumerator struct {
	tasks [tasksOnPage]*uuid.UUID
}

func (t *taskEnumerator) Update(tasks []*models.Task) {
	for i := range t.tasks {
		if i < len(tasks) {
			t.tasks[i] = tasks[i].ID()
		} else {
			t.tasks[i] = nil
		}
	}
}

func (t *taskEnumerator) Get(num uint64) *uuid.UUID {
	if isOk, id := t.isOk(num); isOk {
		return t.tasks[id]
	}

	return nil
}

func (t *taskEnumerator) Delete(num uint64) (deleted *uuid.UUID) {
	if isOk, id := t.isOk(num); isOk {
		deleted = t.tasks[id]
		t.tasks[id] = nil
	}

	return
}

func (t *taskEnumerator) isOk(num uint64) (isOk bool, id uint64) {
	id = num - 1
	isOk = id < tasksOnPage && t.tasks[id] != nil

	return
}
