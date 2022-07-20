package commander

import "fmt"

func NewStartCommand() Command {
	return Command{"start", "get hello message", "",
		func(args string) (string, error) {
			return fmt.Sprintf("Hello %s!", args), nil
		}}
}
