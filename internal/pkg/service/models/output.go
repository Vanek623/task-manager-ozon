package models

import "time"

// TaskBrief краткая информация о задаче
type TaskBrief struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

// DetailedTask подробная информаци о задаче
type DetailedTask struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Edited      time.Time `json:"edited"`
}
