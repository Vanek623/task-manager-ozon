package commander

import (
	"fmt"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/storage"
	"strconv"
)

type UpdateCommand struct {
	bc Command
}

func NewUpdateCommand() ICommand {
	return UpdateCommand{Command{UPDATE_NAME, "edit task", "<id> <title> <description>"}}
}

func (c UpdateCommand) Help() string {
	return c.bc.Help()
}

func (c UpdateCommand) Execute(args string) (string, error) {
	argsArr, err := extractArgs(args)
	if err != nil {
		return "", err
	}

	if len(argsArr) < 2 {
		return "", errors.New("Has no enough args")
	}

	var t *storage.Task
	if id, err := strconv.ParseUint(argsArr[0], 10, 64); err != nil {
		return "", err
	} else if t, err = storage.Get(uint(id)); err != nil {
		return "", err
	}

	if err = t.SetTitle(argsArr[1]); err != nil {
		return "", err
	} else if len(argsArr) > 2 {
		if err = t.SetDescription(argsArr[2]); err != nil {
			return "", err
		}
	}

	if err = storage.Update(t); err != nil {
		return "", err
	}

	return fmt.Sprintf("Task %d: %s updated", t.ID(), t.Title()), nil
}
