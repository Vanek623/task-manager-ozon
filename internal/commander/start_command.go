package commander

import "fmt"

func newStartCommand() command {
	return command{"start", "get hello message", "",
		func(args string) (string, error) {
			return fmt.Sprintf("Hello %s!", args), nil
		}}
}
