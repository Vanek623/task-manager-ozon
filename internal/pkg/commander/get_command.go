package commander

import (
	"fmt"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/storage"
	"strconv"
	"time"
)

func newGetCommand() command {
	return command{"get", "getting task", "<id>",
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

			t, err := storage.Get(uint(id))
			if err != nil {
				return "", err
			}

			return fmt.Sprintf("Title: %s \nDescription: %s \nCreated: %s",
				t.Title(),
				t.Description(),
				t.Created().Format(time.Stamp)), nil
		}}
}
