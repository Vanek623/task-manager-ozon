package command

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	serviceModelsPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
)

const tasksOnPage = 5

type listCommand struct {
	command
}

func (c *listCommand) Execute(ctx context.Context, args string) string {
	argsArr, err := extractArgsCounted(args, 0, 1)
	if err != nil {
		return err.Error()
	}

	var pageNum uint64
	if argsArr[0] == "" {
		pageNum = 1
	} else {
		if pageNum, err = strconv.ParseUint(argsArr[0], 10, 64); err != nil {
			return err.Error()
		}

		if pageNum == 0 {
			return "page number must be not zero"
		}
	}

	data := serviceModelsPkg.ListTaskData{
		MaxTasksCount: tasksOnPage,
		Offset:        tasksOnPage * (uint(pageNum) - 1),
	}

	tasks, err := c.service.TasksList(ctx, data)
	if err != nil {
		return err.Error()
	}

	if len(tasks) == 0 {
		return "There are no tasks!"
	}

	out := make([]string, 0, len(tasks))
	for _, t := range tasks {
		out = append(out, fmt.Sprintf("%d. %s", t.ID, t.Title))
	}

	return strings.Join(out, "\n")
}

func newListCommand(s iService) *listCommand {
	return &listCommand{
		command{
			name:        "list",
			description: "tasks list",
			subArgs:     "<page>",
			service:     s},
	}
}
