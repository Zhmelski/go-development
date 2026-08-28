package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"task_02/models"
)

const DefaultFileName = "tasks.json"

// 'LoadTasks' loads tasks from a JSON file
func LoadTasks(fileName string) (*models.TaskList, error) {
	// Checking the existence of a file
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// If the file does not exist, return an empty list
		return &models.TaskList{
			Tasks:  []models.Task{},
			LastID: 0,
		}, nil
	}

	// Reading the contents of the file
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Parse JSON
	var taskList models.TaskList
	if err := json.Unmarshal(data, &taskList); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &taskList, nil
}

// 'SaveTasks' saves tasks to a JSON file
func SaveTasks(fileName string, taskList *models.TaskList) error {
	// Convert to JSON with indentation for readability
	data, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize data: %w", err)
	}

	// Write to file
	if err := os.WriteFile(fileName, data, 0644); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
