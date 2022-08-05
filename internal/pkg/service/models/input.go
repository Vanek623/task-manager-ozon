package models

type AddTaskData struct {
	Title, Description string
}

type UpdateTaskData struct {
	ID                 uint
	Title, Description string
}

type ListTaskData struct{}

type DeleteTaskData struct {
	ID uint
}

type GetTaskData struct {
	ID uint
}
