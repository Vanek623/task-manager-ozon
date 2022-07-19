package commander

import "strings"

type HelpCommand struct {
	bc Command
}

func NewHelpCommand() ICommand {
	return HelpCommand{Command{HELP_NAME, "commands list", ""}}
}

func (c HelpCommand) Help() string {
	return c.bc.Help()
}

func (c HelpCommand) Execute(args string) (string, error) {
	var commands []ICommand

	commands = append(commands, NewListCommand())
	commands = append(commands, NewAddCommand())
	commands = append(commands, NewUpdateCommand())
	commands = append(commands, NewDeleteCommand())
	commands = append(commands, NewGetCommand())

	out := make([]string, 0, len(commands)+1)
	out = append(out, c.Help())

	for _, c := range commands {
		out = append(out, c.Help())
	}

	return strings.Join(out, "\n"), nil
}
