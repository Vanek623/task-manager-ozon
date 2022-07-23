package commander

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/storage"
)

func newUpdateCommand() command {
	return command{"update", "edit task", "<id> \"title\" \"description\">",
		func(args string) (string, error) {
			argsArr, err := extractArgs(args)
			if err != nil {
				return "", err
			}

			if len(argsArr) < 2 {
				return "", errors.New("Has no enough args")
			}

			id, err := strconv.ParseUint(argsArr[0], 10, 64)
			if err != nil {
				return "", err
			}

			t, err := storage.Get(uint(id))
			if err != nil {
				return "", err
			}

			if err = t.SetTitle(argsArr[1]); err != nil {
				return "", err
			}

			if len(argsArr) > 2 {
				if err = t.SetDescription(argsArr[2]); err != nil {
					return "", err
				}
			}

			if err = storage.Update(t); err != nil {
				return "", err
			}

			return fmt.Sprintf("Task %d: %s updated", t.ID(), t.Title()), nil
		}}
}
