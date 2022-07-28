package command

import (
	"fmt"
	"strconv"
	"time"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task"
)

type getCommand struct {
	command
}

func (c *getCommand) Execute(args string) string {
	argsArr, err := extractArgs(args)
	if err != nil {
		return err.Error()
	} else if len(argsArr) == 0 {
		return errNoEnoughArgs.Error()
	}

	id, err := strconv.ParseUint(argsArr[0], 10, 64)
	if err != nil {
		return fmt.Sprintf("Cannot parse %s", argsArr[0])
	}

	t, err := c.manager.Get(uint(id))
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("Title: %s \nDescription: %s \nCreated: %s",
		t.Title,
		t.Description,
		t.Created.Format(time.Stamp))
}

func newGetCommand(m task.IManager) *getCommand {
	return &getCommand{
		command{
			name:        "get",
			description: "get task info",
			subArgs:     "<ID>",
			manager:     m},
	}
}
