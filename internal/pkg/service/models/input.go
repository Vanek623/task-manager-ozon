package models

// AddTaskData структура запроса на добавление задачи
type AddTaskData struct {
	title, description string
}

// Title заголовок задачи
func (d *AddTaskData) Title() string {
	return d.title
}

// Description описание задачи
func (d *AddTaskData) Description() string {
	return d.description
}

// NewAddTaskData новый запрос на создание задачи
func NewAddTaskData(title, description string) *AddTaskData {
	return &AddTaskData{
		title:       title,
		description: description,
	}
}

// UpdateTaskData структура запроса на обновление задачи
type UpdateTaskData struct {
	id                 uint64
	title, description string
}

// ID ID задачи
func (d *UpdateTaskData) ID() uint64 {
	return d.id
}

// Title заголовок задачи
func (d *UpdateTaskData) Title() string {
	return d.title
}

// Description описание задачи
func (d *UpdateTaskData) Description() string {
	return d.description
}

// NewUpdateTaskData создание запроса на обновление
func NewUpdateTaskData(ID uint64, title string, description string) *UpdateTaskData {
	return &UpdateTaskData{
		id:          ID,
		title:       title,
		description: description,
	}
}

// ListTaskData структура запроса на получение списка задач
type ListTaskData struct {
	maxTasksCount uint64
	offset        uint64
}

// MaxTasksCount максимальное кол-во выдаваемых задач
func (d *ListTaskData) MaxTasksCount() uint64 {
	return d.maxTasksCount
}

// Offset сколько задач пропустить
func (d *ListTaskData) Offset() uint64 {
	return d.offset
}

// NewListTaskData новый запрос на чтение списка
func NewListTaskData(maxTasksCount uint64, offset uint64) *ListTaskData {
	return &ListTaskData{
		maxTasksCount: maxTasksCount,
		offset:        offset,
	}
}

// DeleteTaskData структура запроса на удаления задачи
type DeleteTaskData struct {
	id uint64
}

// ID ID задачи
func (d *DeleteTaskData) ID() uint64 {
	return d.id
}

// NewDeleteTaskData новый запрос на удаление
func NewDeleteTaskData(ID uint64) *DeleteTaskData {
	return &DeleteTaskData{
		id: ID,
	}
}

// GetTaskData структура запроса на получение описания задачи
type GetTaskData struct {
	id uint64
}

// ID ID задачи
func (d *GetTaskData) ID() uint64 {
	return d.id
}

// NewGetTaskData новый запрос на чтение задачи
func NewGetTaskData(ID uint64) *GetTaskData {
	return &GetTaskData{
		id: ID,
	}
}
