package commander

import (
	"TaskAlertBot/internal/storage"
	"fmt"
	"strconv"
)

type DeleteCommand struct {
	bc Command
}

func (c DeleteCommand) Help() string {
	return c.bc.Help()
}

func (c DeleteCommand) Execute(args string) (string, error) {
	if argsArr, err := extractArgs(args); err != nil {
		return "", err
	} else if id, err := strconv.ParseUint(argsArr[0], 10, 64); err != nil {
		return "", err
	} else if err = storage.Delete(uint(id)); err != nil {
		return "", err
	} else {
		return fmt.Sprintf("Task %s deleted", argsArr[0]), nil
	}
}

func NewDeleteCommand() ICommand {
	return DeleteCommand{Command{DELETE_NAME, "delete task", "<id>"}}
}
