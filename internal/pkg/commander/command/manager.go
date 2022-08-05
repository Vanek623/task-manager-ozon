package command

// Manager структура содержащая команды
type Manager struct {
	commands map[string]ICommand
}

// NewManager создание менеджера команд
func NewManager(s iService) Manager {
	m := Manager{}
	m.createCommands(s)

	return m
}

// GetCommand получение команды по имени
func (m *Manager) GetCommand(name string) ICommand {
	return m.commands[name]
}

func (m *Manager) createCommands(s iService) {
	m.commands = make(map[string]ICommand)

	m.registerCommand(newAddCommand(s))
	m.registerCommand(newDeleteCommand(s))
	m.registerCommand(newGetCommand(s))
	m.registerCommand(newListCommand(s))
	m.registerCommand(newStartCommand())
	m.registerCommand(newUpdateCommand(s))

	m.registerCommand(newHelpCommand(s, m.commands))
}

func (m *Manager) registerCommand(c ICommand) {
	m.commands[c.Name()] = c
}
