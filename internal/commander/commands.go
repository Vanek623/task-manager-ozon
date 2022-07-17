package commander

import (
	"TaskAlertBot/internal/storage"
	"fmt"
	"github.com/pkg/errors"
	"strings"
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
	Execute(args []string) (string, error)
}

type BaseCommand struct {
	Name        string
	Description string
	SubArgs     string
}

func (c BaseCommand) Help() string {
	return fmt.Sprintf("/%s %s - %s", c.Name, c.SubArgs, c.Description)
}

type HelpCommand struct {
	bc BaseCommand
}

func NewHelpCommand() ICommand {
	return HelpCommand{BaseCommand{"help", "commands list", ""}}
}

func (c HelpCommand) Help() string {
	return c.bc.Help()
}

func (c HelpCommand) Execute(args []string) (string, error) {
	var commands []ICommand

	commands = append(commands, NewListCommand())
	commands = append(commands, NewAddCommand())

	out := make([]string, 0, 3)
	out = append(out, c.Help())

	for _, c := range commands {
		out = append(out, c.Help())
	}

	return strings.Join(out, "/n"), nil
}

type ListCommand struct {
	bc BaseCommand
}

func NewListCommand() ICommand {
	return ListCommand{BaseCommand{"list", "tasks list", ""}}
}

func (c ListCommand) Help() string {
	return c.bc.Help()
}

func (c ListCommand) Execute(args []string) (string, error) {
	tasks := storage.Tasks()

	if len(tasks) == 0 {
		return "There are no tasks!", nil
	}

	out := make([]string, 0, len(tasks))
	for _, task := range tasks {
		out = append(out, fmt.Sprintf("%d: %s - %s", task.Id(), task.Name(), task.Description()))
	}

	return strings.Join(out, "/n"), nil
}

type AddCommand struct {
	bc BaseCommand
}

func NewAddCommand() ICommand {
	return AddCommand{BaseCommand{"add", "add new task", "<title> <description>"}}
}

func (c AddCommand) Help() string {
	return c.bc.Help()
}

func (c AddCommand) Execute(args []string) (string, error) {
	if len(args) == 0 || len(args) > 2 {
		return "", errors.New("Invalid args")
	}

	var t *storage.Task
	var err error
	if len(args) == 1 {
		t, err = storage.NewTask(args[0], "")
	} else {
		t, err = storage.NewTask(args[0], args[1])
	}

	if err != nil {
		return "", err
	}

	err = storage.Add(t)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Task %d: \"%s\" added", t.Id(), t.Name()), nil
}

func NewCommand(s string) (ICommand, error) {
	unknownCommandErr := errors.New("Unknown command")
	words := strings.Split(s, " ")
	if len(words) == 0 {
		return nil, unknownCommandErr
	}

	if ctor, ok := commands[words[0]]; ok {
		return ctor(), nil
	}

	return nil, unknownCommandErr
}
