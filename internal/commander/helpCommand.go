package commander

import "strings"

type HelpCommand struct {
	bc baseCommand
}

func NewHelpCommand() ICommand {
	return HelpCommand{baseCommand{"help", "commands list", ""}}
}

func (c HelpCommand) Help() string {
	return c.bc.help()
}

func (c HelpCommand) Execute(args string) (string, error) {
	var commands []ICommand

	commands = append(commands, NewListCommand())
	commands = append(commands, NewAddCommand())

	out := make([]string, 0, 3)
	out = append(out, c.Help())

	for _, c := range commands {
		out = append(out, c.Help())
	}

	return strings.Join(out, "\n"), nil
}
