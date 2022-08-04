package command

import (
	"context"
	"time"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"
)

type addCommand struct {
	command
}

func (c *addCommand) Execute(ctx context.Context, args string) string {
	argsArr, err := extractArgsCounted(args, 1, 2)
	if err != nil {
		return err.Error()
	}

	if _, err = c.manager.Add(ctx, models.Task{
		Title:       argsArr[0],
		Description: argsArr[1],
		Created:     time.Now(),
	}); err != nil {
		return err.Error()
	}

	return "Task added"
}

func newAddCommand(m iTaskStorage) *addCommand {
	return &addCommand{
		command{
			name:        "add",
			description: "add task",
			subArgs:     "<name> <description>",
			manager:     m},
	}
}
