package commander

import "fmt"

type baseCommand struct {
	Name        string
	Description string
	SubArgs     string
}

func (c baseCommand) help() string {
	return fmt.Sprintf("/%s %s - %s", c.Name, c.SubArgs, c.Description)
}
