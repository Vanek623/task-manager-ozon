package command

import (
	"fmt"
)

type startCommand struct {
	command
}

func (c *startCommand) Execute(args string) string {
	return fmt.Sprintf("Hello %s!", args)
}

func newStartCommand() *startCommand {
	return &startCommand{
		command{
			name:        "start",
			description: "get hello message",
		},
	}
}
