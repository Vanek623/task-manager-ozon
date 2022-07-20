package commander

import (
	"errors"
	"fmt"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/storage"
)

type AddCommand struct {
	bc Command
}

func NewAddCommand() ICommand {
	return AddCommand{Command{ADD_NAME, "add new task", "<title> <description>"}}
}

func (c AddCommand) Help() string {
	return c.bc.Help()
}

func (c AddCommand) Execute(args string) (string, error) {
	var t *storage.Task
	if argsArr, err := extractArgs(args); err != nil || len(argsArr) == 0 {
		if err != nil {
			return "", err
		} else {
			return "", errors.New("has no enough args")
		}
	} else if len(argsArr) == 1 {
		if t, err = storage.NewTask(argsArr[0], ""); err != nil {
			return "", err
		}
	} else if len(argsArr) == 2 {
		if t, err = storage.NewTask(argsArr[0], argsArr[1]); err != nil {
			return "", err
		}
	}

	if err := storage.Add(t); err != nil {
		return "", err
	}

	return fmt.Sprintf("Task %d: \"%s\" added", t.ID(), t.Title()), nil
}
