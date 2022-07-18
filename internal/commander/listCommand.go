package commander

import (
	"TaskAlertBot/internal/storage"
	"strings"
)

type ListCommand struct {
	bc baseCommand
}

func NewListCommand() ICommand {
	return ListCommand{baseCommand{"list", "tasks list", ""}}
}

func (c ListCommand) Help() string {
	return c.bc.help()
}

func (c ListCommand) Execute(args string) (string, error) {
	tasks := storage.Tasks()

	if len(tasks) == 0 {
		return "There are no tasks!", nil
	}

	out := make([]string, 0, len(tasks))
	for _, task := range tasks {
		out = append(out, task.String())
	}

	return strings.Join(out, "\n"), nil
}
