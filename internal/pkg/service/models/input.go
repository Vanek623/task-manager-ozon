package models

// AddTaskData структура запроса на добавление задачи
type AddTaskData struct {
	Title, Description string
}

// UpdateTaskData структура запроса на обновление задачи
type UpdateTaskData struct {
	ID                 uint
	Title, Description string
}

// ListTaskData структура запроса на получение списка задач
type ListTaskData struct {
	MaxTasksCount uint
	Offset        uint
}

// DeleteTaskData структура запроса на удаления задачи
type DeleteTaskData struct {
	ID uint
}

// GetTaskData структура запроса на получение описания задачи
type GetTaskData struct {
	ID uint
}
