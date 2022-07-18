package commander

import (
	"github.com/pkg/errors"
)

var commands map[string]func() ICommand

func init() {
	commands = make(map[string]func() ICommand)

	commands["help"] = NewHelpCommand
	commands["list"] = NewListCommand
	commands["add"] = NewAddCommand
}

type ICommand interface {
	Help() string
	Execute(args string) (string, error)
}

func NewCommand(command string) (ICommand, error) {
	if ctor, ok := commands[command]; ok {
		return ctor(), nil
	}

	return nil, errors.New("Unknown command")
}
