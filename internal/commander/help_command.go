package commander

import (
	"strings"
)

var helpList string

func newHelpCommand(commands map[string]command) command {
	helpArr := make([]string, 0, len(commands))
	for _, s := range commands {
		helpArr = append(helpArr, s.Help())
	}

	tmp := command{name: "help", description: "commands list"}
	helpArr = append(helpArr, tmp.Help())

	helpList = strings.Join(helpArr, "\n")

	tmp.Execute = func(args string) (string, error) {
		return helpList, nil
	}

	return tmp
}
