package commander

type StartCommand struct {
	bc Command
}

func (c StartCommand) Help() string {
	return c.bc.Help()
}

func (c StartCommand) Execute(args string) (string, error) {
	return "Hello, user!", nil
}

func NewStartCommand() ICommand {
	return StartCommand{Command{START_COMMAND, "get hello message", ""}}
}
