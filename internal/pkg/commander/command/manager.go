package command

import (
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task"
)

// Manager структура содержащая команды
type Manager struct {
	commands map[string]ICommand
}

// NewManager создание менеджера команд
func NewManager(tm task.IManager) Manager {
	m := Manager{}
	m.createCommands(tm)

	return m
}

// GetCommand получение команды по имени
func (m *Manager) GetCommand(name string) ICommand {
	return m.commands[name]
}

func (m *Manager) createCommands(tm task.IManager) {
	m.commands = make(map[string]ICommand)

	m.registerCommand(newAddCommand(tm))
	m.registerCommand(newDeleteCommand(tm))
	m.registerCommand(newGetCommand(tm))
	m.registerCommand(newListCommand(tm))
	m.registerCommand(newStartCommand())
	m.registerCommand(newUpdateCommand(tm))

	m.registerCommand(newHelpCommand(tm, m.commands))
}

func (m *Manager) registerCommand(c ICommand) {
	m.commands[c.Name()] = c
}
