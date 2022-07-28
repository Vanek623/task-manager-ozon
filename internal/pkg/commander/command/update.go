package command

import (
	"fmt"
	"strconv"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task"
)

type updateCommand struct {
	command
}

func (c *updateCommand) Execute(args string) string {
	argsArr, err := extractArgsCounted(args, 2, 3)
	if err != nil {
		return err.Error()
	}

	id, err := strconv.ParseUint(argsArr[0], 10, 64)
	if err != nil {
		return fmt.Sprintf("Cannot parse %s", argsArr[0])
	}

	t, err := c.manager.Get(uint(id))
	if err != nil {
		return err.Error()
	}

	t.Title = argsArr[1]
	t.Description = argsArr[2]

	if err = c.manager.Update(t); err != nil {
		return err.Error()
	}

	return "Task updated"
}

func newUpdateCommand(m task.IManager) *updateCommand {
	return &updateCommand{
		command{
			name:        "update",
			description: "update task",
			subArgs:     "<id> <name> <description>",
			manager:     m},
	}
}
