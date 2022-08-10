package command

import (
	"context"

	serviceModelsPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
)

type addCommand struct {
	command
}

func (c *addCommand) Execute(ctx context.Context, args string) string {
	argsArr, err := extractArgsCounted(args, 1, 2)
	if err != nil {
		return err.Error()
	}

	if _, err = c.service.AddTask(ctx, serviceModelsPkg.AddTaskData{
		Title:       argsArr[0],
		Description: argsArr[1],
	}); err != nil {
		return err.Error()
	}

	return "TaskBrief added"
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
