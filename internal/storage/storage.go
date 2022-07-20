package storage

import (
	"github.com/pkg/errors"
	"log"
	"strconv"
)

var tasks map[uint]*Task

const maxTasks = 8

func init() {
	tasks = make(map[uint]*Task)

	if t, e := NewTask("Create new Task",
		"Test task creating"); e != nil {
		log.Panicln(e)
	} else if e = Add(t); e != nil {
		log.Panicln(e)
	}
}

func Tasks() []*Task {
	res := make([]*Task, 0, len(tasks))

	for _, t := range tasks {
		res = append(res, t)
	}

	return res
}

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

func Update(t *Task) error {
	if _, ok := tasks[t.ID()]; !ok {
		return makeTaskExistError(false, t.ID())
	}

	tasks[t.ID()] = t
	return nil
}

func Delete(id uint) error {
	if _, ok := tasks[id]; !ok {
		return makeTaskExistError(false, id)
	}

	delete(tasks, id)
	return nil
}

func Get(id uint) (*Task, error) {
	if t, ok := tasks[id]; !ok {
		return nil, makeTaskExistError(false, id)
	} else {
		return t, nil
	}
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
