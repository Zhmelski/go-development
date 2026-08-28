package models

import "time"

// 'Task' create a task on the TODO list
type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

// 'TaskList' presents a list of all tasks
type TaskList struct {
	Tasks  []Task `json:"tasks"`
	LastID int    `json:"last_id"`
}
