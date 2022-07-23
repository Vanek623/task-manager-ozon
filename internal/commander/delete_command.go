package commander

import (
	"fmt"
	"strconv"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/storage"
)

func newDeleteCommand() command {
	return command{"delete", "delete task", "<id>",
		func(args string) (string, error) {
			argsArr, err := extractArgs(args)
			if err != nil {
				return "", err
			} else if len(argsArr) == 0 {
				return "", errNoEnoughArgs
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
