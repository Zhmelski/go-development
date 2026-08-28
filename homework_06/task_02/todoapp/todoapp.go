package todoapp

import (
	"fmt"
	"time"

	"task_02/models"
	"task_02/storage"
)

// 'TodoApp' presents a task list application
type TodoApp struct {
	fileName string
	taskList *models.TaskList
}

// 'NewTodoApp' creates a new application instance
func NewTodoApp(fileName string) (*TodoApp, error) {
	taskList, err := storage.LoadTasks(fileName)
	if err != nil {
		return nil, err
	}

	return &TodoApp{
		fileName: fileName,
		taskList: taskList,
	}, nil
}

// 'AddTask' adds a new task
func (app *TodoApp) AddTask(description string) error {
	app.taskList.LastID++

	newTask := models.Task{
		ID:          app.taskList.LastID,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	app.taskList.Tasks = append(app.taskList.Tasks, newTask)

	// Save changes
	if err := storage.SaveTasks(app.fileName, app.taskList); err != nil {
		return fmt.Errorf("failed to save task: %w", err)
	}

	fmt.Printf("✓ Task #%d added: %s\n", newTask.ID, description)
	return nil
}

// 'CompleteTask' marks the task as completed
func (app *TodoApp) CompleteTask(taskID int) error {
	// Search for a task by ID
	found := false
	for i := range app.taskList.Tasks {
		if app.taskList.Tasks[i].ID == taskID {
			if app.taskList.Tasks[i].Completed {
				return fmt.Errorf("task #%d is already completed", taskID)
			}

			app.taskList.Tasks[i].Completed = true
			now := time.Now()
			app.taskList.Tasks[i].CompletedAt = &now
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task #%d not found", taskID)
	}

	// Save changes
	if err := storage.SaveTasks(app.fileName, app.taskList); err != nil {
		return fmt.Errorf("failed to save changes: %w", err)
	}

	fmt.Printf("✓ Task #%d is marked as completed\n", taskID)
	return nil
}

// 'ListTasks' displays a list of all tasks
func (app *TodoApp) ListTasks() {
	if len(app.taskList.Tasks) == 0 {
		fmt.Println("The task list is empty")
		return
	}

	fmt.Println("\n📋 To-do list:")
	fmt.Println("═══════════════════════════════════════════════════════════════")

	for _, task := range app.taskList.Tasks {
		status := "❌ Not completed"
		if task.Completed {
			status = "✅ Completed"
		}

		fmt.Printf("\n#%d | %s\n", task.ID, status)
		fmt.Printf("Description: %s\n", task.Description)
		fmt.Printf("Created:  %s\n", task.CreatedAt.Format("02.01.2006 15:04:05"))

		if task.CompletedAt != nil {
			fmt.Printf("Completed: %s\n", task.CompletedAt.Format("02.01.2006 15:04:05"))
		}
	}

	fmt.Println("\n═══════════════════════════════════════════════════════════════")
}
