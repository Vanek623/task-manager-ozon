package commander

import (
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/storage"
	"strings"
)

func NewListCommand() Command {
	return Command{"list", "tasks list", "",
		func(args string) (string, error) {
			tasks := storage.Tasks()

			if len(tasks) == 0 {
				return "There are no tasks!", nil
			}

			out := make([]string, 0, len(tasks))
			for _, task := range tasks {
				out = append(out, task.String())
			}

			return strings.Join(out, "\n"), nil
		}}
}
