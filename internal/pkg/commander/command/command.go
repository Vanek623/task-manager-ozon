package command

import (
	"fmt"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"
	"strings"

	"github.com/pkg/errors"
)

// ICommand интерфейс команды
type ICommand interface {
	Name() string
	Description() string
	SubArgs() string
	Execute(args string) string
	Help() string
}

type iTaskManager interface {
	Add(t models.Task) error
	Delete(ID uint) error
	List() []models.Task
	Update(t models.Task) error
	Get(ID uint) (models.Task, error)
}

type command struct {
	name        string
	description string
	subArgs     string
	manager     iTaskManager
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

func (c *command) Execute(args string) string {
	return "Cannot exec this command"
}

func (c *command) Help() string {
	return fmt.Sprintf("/%s %s - %s", c.name, c.subArgs, c.description)
}

var errNoEnoughArgs = errors.New("has no enough arguments")

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
		return argsArr, errors.New("Has no enough args")
	}

	if len(argsArr) < max {
		tmp := make([]string, max-len(argsArr))
		argsArr = append(argsArr, tmp...)
	}

	return argsArr, nil
}
