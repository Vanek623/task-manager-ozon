package storage

import (
	"strconv"

	"github.com/pkg/errors"
)

var tasks map[uint]*Task

const maxTasks = 8

func init() {
	tasks = make(map[uint]*Task)
}

// Tasks чтение списка задач
func Tasks() []*Task {
	res := make([]*Task, 0, len(tasks))

	for _, t := range tasks {
		res = append(res, t)
	}

	return res
}

// Add добавление задачи
func Add(t *Task) error {
	if len(tasks) >= maxTasks {
		return errors.New("Has no space for tasks, please delete one")
	}
	if _, ok := tasks[t.ID()]; ok {
		return makeTaskExistError(true, t.ID())
	}

	tasks[t.ID()] = t
	return nil
}

// Update обновление задачи
func Update(t *Task) error {
	if _, ok := tasks[t.ID()]; !ok {
		return makeTaskExistError(false, t.ID())
	}

	tasks[t.ID()] = t
	return nil
}

// Delete удаление задачи
func Delete(id uint) error {
	if _, ok := tasks[id]; !ok {
		return makeTaskExistError(false, id)
	}

	delete(tasks, id)
	return nil
}

// Get чтение задачи
func Get(id uint) (*Task, error) {
	if _, ok := tasks[id]; !ok {
		return nil, makeTaskExistError(false, id)
	}

	return tasks[id], nil
}

func makeTaskExistError(isExist bool, id uint) error {
	var e error
	if isExist {
		e = errors.New("task already exist")
	} else {
		e = errors.New("task doesn't exist")
	}

	return errors.Wrap(e, strconv.FormatUint(uint64(id), 10))
}
