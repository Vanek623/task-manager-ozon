package commander

import (
	"strings"
)

var helpList []string

func NewHelpCommand(commands map[string]Command) Command {
	helpList = make([]string, 0, len(commands))
	for _, s := range commands {
		helpList = append(helpList, s.Help())
	}

	tmp := Command{Name: "help", Description: "commands list"}
	helpList = append(helpList, tmp.Help())

	tmp.Execute = func(args string) (string, error) {
		return strings.Join(helpList, "\n"), nil
	}

	return tmp
}
