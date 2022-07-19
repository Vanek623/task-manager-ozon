package commander

import (
	"github.com/pkg/errors"
)

var commands map[string]func() ICommand

const (
	HELP_NAME     = "help"
	LIST_NAME     = "list"
	ADD_NAME      = "add"
	UPDATE_NAME   = "update"
	DELETE_NAME   = "delete"
	GET_NAME      = "get"
	START_COMMAND = "start"
)

func init() {
	commands = make(map[string]func() ICommand)

	commands[HELP_NAME] = NewHelpCommand
	commands[LIST_NAME] = NewListCommand
	commands[ADD_NAME] = NewAddCommand
	commands[UPDATE_NAME] = NewUpdateCommand
	commands[DELETE_NAME] = NewDeleteCommand
	commands[GET_NAME] = NewGetCommand
	commands[START_COMMAND] = NewStartCommand
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
