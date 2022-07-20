package commander

import (
	"fmt"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/storage"
	"strconv"
	"time"
)

func NewGetCommand() Command {
	return Command{"get", "getting task", "<id>",
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
