package command

import (
	"strings"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task"
)

type helpCommand struct {
	command
	helpList string
}

func (c *helpCommand) Execute(args string) string {
	return c.helpList
}

func newHelpCommand(m task.IManager, commands map[string]ICommand) *helpCommand {
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
