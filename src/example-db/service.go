package main

import (
	"encoding/json"
	"log"
	"os"
)

func listTasks() ([]Task, error) {
	taskBytes, err := os.ReadFile("./data/tasks.json")
	if err != nil {
		log.Println("Failed to read tasks.json...will return default task(s)")
		taskBytes = []byte(`[{"id":1,"name":"Default task"}]`)
	}
	var tasks []Task
	err = json.Unmarshal(taskBytes, &tasks)
	if err != nil {
		log.Println("Failed to unmarshal tasks.json due to:", err)
		return nil, err
	}
	return tasks, nil
}

func insertTask(task Task) (Task, error) {
	tasks, err := listTasks()
	if err != nil {
		return Task{}, err
	}
	task.ID = len(tasks) + 1
	tasks = append(tasks, task)
	taskBytes, err := json.Marshal(tasks)
	if err != nil {
		log.Println("Failed to marshal tasks due to:", err)
		return Task{}, err
	}
	err = os.WriteFile("./data/tasks.json", taskBytes, 0777)
	if err != nil {
		log.Println("Failed to write tasks.json due to:", err)
		return Task{}, err
	}
	return task, nil
}
