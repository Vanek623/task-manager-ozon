package command

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"

	"github.com/pkg/errors"
)

// ICommand интерфейс команды
type ICommand interface {
	Name() string
	Description() string
	SubArgs() string
	Execute(ctx context.Context, args string) string
	Help() string
}

type iService interface {
	AddTask(ctx context.Context, data *models.AddTaskData) (*uuid.UUID, error)
	DeleteTask(ctx context.Context, data *models.DeleteTaskData) error
	TasksList(ctx context.Context, data *models.ListTaskData) ([]*models.Task, error)
	UpdateTask(ctx context.Context, data *models.UpdateTaskData) error
	GetTask(ctx context.Context, data *models.GetTaskData) (*models.DetailedTask, error)
}

type command struct {
	name        string
	description string
	subArgs     string
	service     iService
}

func (c *command) Name() string {
	return c.name
}

func (c *command) Description() string {
	return c.description
}

func (c *command) SubArgs() string {
	return c.subArgs
}

func (c *command) Execute(_ context.Context, _ string) string {
	return "Cannot exec this command"
}

func (c *command) Help() string {
	return fmt.Sprintf("/%s %s - %s", c.name, c.subArgs, c.description)
}

// ErrNoEnoughArgs недостаточно аргументов
var ErrNoEnoughArgs = errors.New("has no enough arguments")

func extractArgs(args string) ([]string, error) {
	var out []string
	for len(args) != 0 {
		if args[0] == ' ' {
			args = args[1:]
			continue
		}

		var subArgs []string
		if args[0] == '"' {
			subArgs = strings.SplitAfterN(args[1:], "\"", 2)
			if len(subArgs) != 2 {
				return nil, errors.Errorf("Cannot parse %s", args)
			}
		} else {
			subArgs = strings.SplitAfterN(args, " ", 2)
			if len(subArgs) == 1 {
				out = append(out, subArgs[0])
				break
			}
		}

		out = append(out, subArgs[0][0:len(subArgs[0])-1])
		args = subArgs[1]
	}

	return out, nil
}

func extractArgsCounted(args string, min, max int) ([]string, error) {
	argsArr, err := extractArgs(args)
	if err != nil {
		return argsArr, err
	}

	if len(argsArr) < min {
		return argsArr, ErrNoEnoughArgs
	}

	if len(argsArr) < max {
		tmp := make([]string, max-len(argsArr))
		argsArr = append(argsArr, tmp...)
	}

	return argsArr, nil
}
