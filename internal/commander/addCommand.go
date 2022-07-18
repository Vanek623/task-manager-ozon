package commander

import (
	"TaskAlertBot/internal/storage"
	"fmt"
	"github.com/pkg/errors"
)

type AddCommand struct {
	bc baseCommand
}

func NewAddCommand() ICommand {
	return AddCommand{baseCommand{"add", "add new task", "<title> <description>"}}
}

func (c AddCommand) Help() string {
	return c.bc.help()
}

func (c AddCommand) Execute(args string) (string, error) {
	fmt.Println("exec add")
	if len(args) == 0 {
		return "", errors.New("Invalid args")
	}

	argsArr := extractQuotArgs(args)

	var t *storage.Task
	var err error
	if len(argsArr) == 0 {
		t, err = nil, errors.New("Invalid arguments!")
	} else if len(argsArr) == 1 {
		t, err = storage.NewTask(argsArr[0], "")
	} else {
		t, err = storage.NewTask(argsArr[0], argsArr[1])
	}

	if err != nil {
		return "", err
	}

	if err = storage.Add(t); err != nil {
		return "", err
	}

	return fmt.Sprintf("Task %d: \"%s\" added", t.Id(), t.Title()), nil
}
