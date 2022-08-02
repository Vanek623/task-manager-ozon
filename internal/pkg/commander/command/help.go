package command

import (
	"context"
	"strings"
)

type helpCommand struct {
	command
	helpList string
}

func (c *helpCommand) Execute(_ context.Context, _ string) string {
	return c.helpList
}

func newHelpCommand(m iTaskStorage, commands map[string]ICommand) *helpCommand {
	tmp := helpCommand{
		command: command{
			name:        "help",
			description: "get commands list",
			manager:     m},
		helpList: "",
	}

	helpArr := make([]string, 0, len(commands)+1)
	for _, s := range commands {
		helpArr = append(helpArr, s.Help())
	}
	helpArr = append(helpArr, tmp.Help())

	tmp.helpList = strings.Join(helpArr, "\n")

	return &tmp
}
