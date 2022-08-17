package command

import (
	"context"
	"fmt"
	"strconv"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
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

	data := models.NewUpdateTaskData(id, argsArr[1], argsArr[2])
	if err = c.service.UpdateTask(ctx, data); err != nil {
		return err.Error()
	}

	return "TaskBrief updated"
}

func newUpdateCommand(s iService) *updateCommand {
	return &updateCommand{
		command{
			name:        "update",
			description: "update task",
			subArgs:     "<ID> <name> <description>",
			service:     s},
	}
}
