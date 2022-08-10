package models

// AddTaskData структура запроса на добавление задачи
type AddTaskData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// UpdateTaskData структура запроса на обновление задачи
type UpdateTaskData struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// ListTaskData структура запроса на получение списка задач
type ListTaskData struct {
	Limit  uint `json:"limit"`
	Offset uint `json:"offset"`
}

// DeleteTaskData структура запроса на удаления задачи
type DeleteTaskData struct {
	ID uint `json:"id"`
}

// GetTaskData структура запроса на получение описания задачи
type GetTaskData struct {
	ID uint `json:"id"`
}
