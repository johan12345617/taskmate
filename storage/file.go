package storage

import (
	"encoding/json"
	"os"
	"taskmate/task"
)

const jsonFileName string = "tasks-data.json"

func ReadTasksFromFile() ([]task.Task, error) {
	// Open the file
	file, err := os.Open(jsonFileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the JSON into a slice of Task
	var tasks []task.Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}
