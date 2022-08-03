package command

import (
	"context"
	"fmt"
	"strconv"
	"time"
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

	id, err := strconv.ParseUint(argsArr[0], 10, 64)
	if err != nil {
		return fmt.Sprintf("Cannot parse %s", argsArr[0])
	}

	t, err := c.manager.Get(ctx, uint(id))
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("Title: %s \nDescription: %s \nCreated: %s",
		t.Title,
		t.Description,
		t.Created.Format(time.Stamp))
}

func newGetCommand(m iTaskStorage) *getCommand {
	return &getCommand{
		command{
			name:        "get",
			description: "get task info",
			subArgs:     "<ID>",
			manager:     m},
	}
}
