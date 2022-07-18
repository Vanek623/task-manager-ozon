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
	commands = append(commands, NewUpdateCommand())

	out := make([]string, 0, len(commands)+1)
	out = append(out, c.Help())

	for _, c := range commands {
		out = append(out, c.Help())
	}

	return strings.Join(out, "\n"), nil
}
