package commander

import (
	"fmt"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/storage"
	"strconv"
)

func NewDeleteCommand() Command {
	return Command{"delete", "delete task", "<id>",
		func(args string) (string, error) {
			argsArr, err := extractArgs(args)
			if err != nil {
				return "", err
			} else if len(argsArr) == 0 {
				return "", hasNoEnoughArgs
			}

			id, err := strconv.ParseUint(argsArr[0], 10, 64)
			if err != nil {
				return "", err
			}

			if err = storage.Delete(uint(id)); err != nil {
				return "", err
			}

			return fmt.Sprintf("Task %s deleted", argsArr[0]), nil
		}}
}
