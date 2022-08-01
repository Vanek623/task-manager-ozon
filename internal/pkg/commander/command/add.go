package command

import (
	"time"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"
)

type addCommand struct {
	command
}

func (c *addCommand) Execute(args string) string {
	argsArr, err := extractArgsCounted(args, 1, 2)
	if err != nil {
		return err.Error()
	}

	if err = c.manager.Add(models.Task{
		Title:       argsArr[0],
		Description: argsArr[1],
		Created:     time.Now(),
	}); err != nil {
		return err.Error()
	}

	return "Task added"
}

func newAddCommand(m iTaskManager) *addCommand {
	return &addCommand{
		command{
			name:        "add",
			description: "add task",
			subArgs:     "<name> <description>",
			manager:     m},
	}
}
