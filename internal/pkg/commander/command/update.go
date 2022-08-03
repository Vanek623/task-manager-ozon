package command

import (
	"context"
	"fmt"
	"strconv"
)

type updateCommand struct {
	command
}

func (c *updateCommand) Execute(ctx context.Context, args string) string {
	argsArr, err := extractArgsCounted(args, 2, 3)
	if err != nil {
		return err.Error()
	}

	id, err := strconv.ParseUint(argsArr[0], 10, 64)
	if err != nil {
		return fmt.Sprintf("Cannot parse %s", argsArr[0])
	}

	t, err := c.manager.Get(ctx, uint(id))
	if err != nil {
		return err.Error()
	}

	t.Title = argsArr[1]
	t.Description = argsArr[2]

	if err = c.manager.Update(ctx, *t); err != nil {
		return err.Error()
	}

	return "Task updated"
}

func newUpdateCommand(m iTaskStorage) *updateCommand {
	return &updateCommand{
		command{
			name:        "update",
			description: "update task",
			subArgs:     "<ID> <name> <description>",
			manager:     m},
	}
}
