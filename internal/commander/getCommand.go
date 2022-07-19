package commander

import (
	"TaskAlertBot/internal/storage"
	"fmt"
	"strconv"
	"time"
)

type GetCommand struct {
	bc Command
}

func (c GetCommand) Help() string {
	return c.bc.Help()
}

func (c GetCommand) Execute(args string) (string, error) {
	if argsArr, err := extractArgs(args); err != nil {
		return "", err
	} else if id, err := strconv.ParseUint(argsArr[0], 10, 64); err != nil {
		return "", err
	} else if t, err := storage.Get(uint(id)); err != nil {
		return "", err
	} else {
		return fmt.Sprintf("Title: %s \nDescription: %s \nCreated: %s", t.Title(), t.Description(), t.Created().Format(time.Stamp)), err
	}
}

func NewGetCommand() ICommand {
	return GetCommand{Command{GET_NAME, "getting task", "<id>"}}
}
