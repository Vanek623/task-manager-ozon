package command

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
)

type deleteCommand struct {
	command
}

func (c *deleteCommand) Execute(ctx context.Context, args string) string {
	argsArr, err := extractArgsCounted(args, 1, 1)
	if err != nil {
		return err.Error()
	}

	id, err := uuid.Parse(argsArr[0])
	if err != nil {
		return fmt.Sprintf("Cannot parse %s", argsArr[0])
	}

	if err = c.service.DeleteTask(ctx, models.NewDeleteTaskData(&id)); err != nil {
		return err.Error()
	}

	return "Task deleted"
}

func newDeleteCommand(s iService) *deleteCommand {
	return &deleteCommand{
		command{
			name:        "delete",
			description: "delete task",
			subArgs:     "<ID>",
			service:     s},
	}
}
