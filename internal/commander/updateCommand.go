package commander

import (
	"TaskAlertBot/internal/storage"
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

type UpdateCommand struct {
	bc baseCommand
}

func NewUpdateCommand() ICommand {
	return UpdateCommand{baseCommand{"update", "edit task", "<id> <title> <description>"}}
}

func (c UpdateCommand) Help() string {
	return c.bc.help()
}

func (c UpdateCommand) Execute(args string) (string, error) {
	id, err := extractUint(args)
	if err != nil {
		return "", err
	}

	var t *storage.Task
	t, err = storage.Get(id)
	if err != nil {
		return "", err
	}

	argsArr := extractQuotArgs(args)
	if len(argsArr) == 0 {
		return "", errors.New("Title must be not empty")
	}

	if err = t.SetTitle(argsArr[0]); err != nil {
		return "", err
	}

	if len(argsArr) >= 2 {
		if err = t.SetDescription(argsArr[1]); err != nil {
			return "", err
		}
	}

	return fmt.Sprintf("Task %d: %s updated", t.Id(), t.Title()), err
}

func extractUint(s string) (uint, error) {
	begId := len(s)
	endId := len(s)
	for i, ch := range s {
		isChar := ch >= '0' && ch <= '9'
		if isChar && begId == len(s) {
			begId = i
		} else if !isChar && begId != len(s) {
			endId = i
			break
		}
	}

	if begId >= len(s) {
		return 0, errors.New(fmt.Sprintf("Cannot parse <%s>", s))
	}

	if num, err := strconv.ParseUint(s[begId:endId], 10, 64); err == nil {
		return uint(num), nil
	} else {
		return 0, err
	}
}
