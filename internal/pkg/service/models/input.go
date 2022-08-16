package models

// AddTaskData структура запроса на добавление задачи
type AddTaskData struct {
	Title, Description string
}

// UpdateTaskData структура запроса на обновление задачи
type UpdateTaskData struct {
	ID                 uint64
	Title, Description string
}

// ListTaskData структура запроса на получение списка задач
type ListTaskData struct {
	MaxTasksCount uint64
	Offset        uint64
}

// DeleteTaskData структура запроса на удаления задачи
type DeleteTaskData struct {
	ID uint64
}

// GetTaskData структура запроса на получение описания задачи
type GetTaskData struct {
	ID uint64
}
