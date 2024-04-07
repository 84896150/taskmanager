package taskmanipulator

import (
	"encoding/json"
	"fmt"
)

// Task represents a JSON object in an array.
type Task map[string]interface{}

// TaskManipulator provides methods for manipulating tasks within a JSON array.
type TaskManipulator struct{}

// RemoveTaskByName removes a task from a JSON array by its name.
func (tm *TaskManipulator) RemoveTaskByName(jsonStr string, name string) (string, error) {
	var tasks []Task
	err := json.Unmarshal([]byte(jsonStr), &tasks)
	if err != nil {
		return "", err
	}

	var updatedTasks []Task
	for _, task := range tasks {
		if taskName, ok := task["name"].(string); ok && taskName != name {
			updatedTasks = append(updatedTasks, task)
		}
	}

	updatedJSON, err := json.MarshalIndent(updatedTasks, "", "  ")
	if err != nil {
		return "", err
	}

	return string(updatedJSON), nil
}

// AddTask adds a new task to a JSON array.
func (tm *TaskManipulator) AddTask(jsonStr string, newTask map[string]interface{}) (string, error) {
	var tasks []map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &tasks)
	if err != nil {
		return "", err
	}

	// Check if the new task has a 'name' field
	if _, ok := newTask["name"]; !ok || newTask["name"] == "" {
		return "", fmt.Errorf("new task must have a 'name' field")
	}

	// Check if a task with the same name already exists
	for _, existingTask := range tasks {
		if existingName, ok := existingTask["name"]; ok {
			if existingName == newTask["name"] {
				return "", fmt.Errorf("a task with the same name '%v' already exists", newTask["name"])
			}
		}
	}

	// Add the new task to the slice
	tasks = append(tasks, newTask)

	updatedJSON, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return "", err
	}

	return string(updatedJSON), nil
}

// UpdateTask updates an existing task in a JSON array.
func (tm *TaskManipulator) UpdateTask(jsonStr string, updateTask map[string]interface{}) (string, error) {
	var tasks []map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &tasks)
	if err != nil {
		return "", err
	}

	// Check if the update task has a 'name' field
	if _, ok := updateTask["name"]; !ok {
		return "", fmt.Errorf("update task must have a 'name' field")
	}

	// Find and update the existing task
	found := false
	for i, task := range tasks {
		if taskName, ok := task["name"].(string); ok && taskName == updateTask["name"].(string) {
			tasks[i] = updateTask
			found = true
			break
		}
	}

	if !found {
		return "", fmt.Errorf("update with name '%v' not found", updateTask["name"])
	}

	updatedJSON, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return "", err
	}

	return string(updatedJSON), nil
}
