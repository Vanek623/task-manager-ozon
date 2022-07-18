package commander

import (
	"fmt"
)

type baseCommand struct {
	Name        string
	Description string
	SubArgs     string
}

func (c baseCommand) help() string {
	return fmt.Sprintf("/%s %s - %s", c.Name, c.SubArgs, c.Description)
}

func extractQuotArgs(args string) []string {
	argsArr := make([]string, 0, 2)
	isBeginFind := false
	var begId int
	for i, ch := range args {
		if ch == '"' {
			if isBeginFind {
				if begId != i {
					argsArr = append(argsArr, args[begId:i])
				} else {
					argsArr = append(argsArr, "")
				}
			}

			begId = i + 1
			isBeginFind = !isBeginFind
		}
	}

	return argsArr
}
