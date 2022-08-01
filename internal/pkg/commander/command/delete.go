package command

import (
	"fmt"
	"strconv"
)

type deleteCommand struct {
	command
}

func (c *deleteCommand) Execute(args string) string {
	argsArr, err := extractArgsCounted(args, 1, 1)
	if err != nil {
		return err.Error()
	}

	id, err := strconv.ParseUint(argsArr[0], 10, 64)
	if err != nil {
		return fmt.Sprintf("Cannot parse %s", argsArr[0])
	}

	if err = c.manager.Delete(uint(id)); err != nil {
		return err.Error()
	}

	return "Task deleted"
}

func newDeleteCommand(m iTaskManager) *deleteCommand {
	return &deleteCommand{
		command{
			name:        "delete",
			description: "delete task",
			subArgs:     "<ID>",
			manager:     m},
	}
}
