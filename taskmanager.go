package taskmanager

import (
	"encoding/json"
	"fmt"
)

// Task 表示JSON数组中的一个对象
type Task map[string]interface{}

// TaskManager 提供对JSON任务数组的操作接口
type TaskManager struct{}

// RemoveTaskByName 从JSON数组中删除指定name的任务
func (tm *TaskManager) RemoveTaskByName(jsonStr string, name string) (string, error) {
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

// AddTask 向JSON数组中添加新的任务
func (tm *TaskManager) AddTask(jsonStr string, newTask map[string]interface{}) (string, error) {
	var tasks []map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &tasks)
	if err != nil {
		return "", err
	}

	// 检查新任务是否包含name字段
	if _, ok := newTask["name"]; !ok || newTask["name"] == "" {
		return "", fmt.Errorf("new task must have a 'name' field")
	}

	// 检查jsonStr中是否已存在具有相同name字段的任务
	for _, existingTask := range tasks {
		if existingName, ok := existingTask["name"]; ok {
			if existingName == newTask["name"] {
				return "", fmt.Errorf("a task with the same name '%v' already exists", newTask["name"])
			}
		}
	}

	// 添加新任务到切片
	tasks = append(tasks, newTask)

	updatedJSON, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return "", err
	}

	return string(updatedJSON), nil
}

// UpdateTask 更新JSON数组中指定name的任务
func (tm *TaskManager) UpdateTask(jsonStr string, updateTask map[string]interface{}) (string, error) {
	var tasks []map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &tasks)
	if err != nil {
		return "", err
	}

	// 检查新任务是否包含name字段
	if _, ok := updateTask["name"]; !ok {
		return "", fmt.Errorf("new task must have a 'name' field")
	}

	// 查找并更新已有任务
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
