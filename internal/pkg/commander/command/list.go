package command

import (
	"context"
	"fmt"
	serviceModelsPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
	"strings"
)

type listCommand struct {
	command
}

func (c *listCommand) Execute(ctx context.Context, _ string) string {
	tasks, err := c.service.TasksList(ctx, serviceModelsPkg.ListTaskData{})
	if err != nil {
		return err.Error()
	}

	if len(tasks) == 0 {
		return "There are no tasks!"
	}

	out := make([]string, 0, len(tasks))
	for _, t := range tasks {
		out = append(out, fmt.Sprintf("%d. %s", t.ID, t.Title))
	}

	return strings.Join(out, "\n")
}

func newListCommand(s iService) *listCommand {
	return &listCommand{
		command{
			name:        "list",
			description: "tasks list",
			service:     s},
	}
}
