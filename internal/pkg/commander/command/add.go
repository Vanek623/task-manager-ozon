package command

import (
	"context"
	"fmt"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
)

type addCommand struct {
	command
}

func (c *addCommand) Execute(ctx context.Context, args string) string {
	argsArr, err := extractArgsCounted(args, 1, 2)
	if err != nil {
		return err.Error()
	}

	id, err := c.service.AddTask(ctx, models.NewAddTaskData(argsArr[0], argsArr[1]))
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("Task #%d added", id)
}

func newAddCommand(s iService) *addCommand {
	return &addCommand{
		command{
			name:        "add",
			description: "add task",
			subArgs:     "<name> <description>",
			service:     s},
	}
}
