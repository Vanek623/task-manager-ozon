package commander

import (
	"TaskAlertBot/internal/storage"
	"strings"
)

type ListCommand struct {
	bc Command
}

func NewListCommand() ICommand {
	return ListCommand{Command{LIST_NAME, "tasks list", ""}}
}

func (c ListCommand) Help() string {
	return c.bc.Help()
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
