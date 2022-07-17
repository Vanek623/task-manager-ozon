package commander

const (
	Help = iota
	List
	Add

	//del
	//edit
	//get
)

type Command struct {
	Name        string
	Subargs     string
	Description string
	Function    func()
}

func newCmdHelp() Command {
	return Command{"help", "", "list of commands", func() {

	}}
}

func newCmdList() Command {
	return Command{"list", "", "list of tasks", func() {

	}}
}

func newCmdAdd() Command {
	return Command{"add", "<name> <password>",
		"add new task by name and description", func() {

		}}
}

func newCommand(cmdType int) (Command, error) {
	switch cmdType {
	case Help:
		return
	}
}
