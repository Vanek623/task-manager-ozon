package commander

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

type baseCommand struct {
	Name        string
	Description string
	SubArgs     string
}

func (c baseCommand) help() string {
	return fmt.Sprintf("/%s %s - %s", c.Name, c.SubArgs, c.Description)
}

func extractArgs(args string) ([]string, error) {
	var out []string
	for len(args) != 0 {
		if args[0] == ' ' {
			args = args[1:]
			continue
		}

		var subargs []string
		if args[0] == '"' {
			subargs = strings.SplitAfterN(args[1:], "\"", 2)
			if len(subargs) != 2 {
				return nil, errors.New(fmt.Sprintf("Cannot parse %s", args))
			}
		} else {
			subargs = strings.SplitAfterN(args, " ", 2)
			if len(subargs) == 1 {
				out = append(out, subargs[0])
				break
			}
		}

		out = append(out, subargs[0][0:len(subargs[0])-1])
		args = subargs[1]
	}

	return out, nil
}