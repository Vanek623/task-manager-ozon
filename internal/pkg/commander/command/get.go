package command

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
)

type getCommand struct {
	command
}

func (c *getCommand) Execute(ctx context.Context, args string) string {
	argsArr, err := extractArgs(args)
	if err != nil {
		return err.Error()
	} else if len(argsArr) == 0 {
		return ErrNoEnoughArgs.Error()
	}

	id, err := uuid.Parse(argsArr[0])
	if err != nil {
		return fmt.Sprintf("Cannot parse %s", argsArr[0])
	}

	t, err := c.service.GetTask(ctx, models.NewGetTaskData(&id))
	if err != nil {
		return err.Error()
	}

	return t.String()
}

func newGetCommand(s iService) *getCommand {
	return &getCommand{
		command{
			name:        "get",
			description: "get task info",
			subArgs:     "<ID>",
			service:     s},
	}
}
