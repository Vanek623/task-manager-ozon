package commander

import (
	"fmt"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/storage"
)

func newAddCommand() command {
	return command{"add", "add new task", "\"title\" \"description\"",
		func(args string) (string, error) {
			argsArr, err := extractArgs(args)
			if err != nil {
				return "", err
			}

			var t *storage.Task
			switch len(argsArr) {
			case 0:
				return "", hasNoEnoughArgs
			case 1:
				if t, err = storage.NewTask(argsArr[0], ""); err != nil {
					return "", err
				}
			default:
				if t, err = storage.NewTask(argsArr[0], argsArr[1]); err != nil {
					return "", err
				}
			}

			if err := storage.Add(t); err != nil {
				return "", err
			}

			return fmt.Sprintf("Task %d: \"%s\" added", t.ID(), t.Title()), nil
		}}
}
