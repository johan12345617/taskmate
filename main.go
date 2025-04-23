package main

import (
	"flag"
	"fmt"
	"taskmate/storage"
	"taskmate/task"
	"taskmate/utils"
)

var listTasks []task.Task = []task.Task{}

func main() {
	//read the json file once
	listTasks, err := storage.ReadTasksFromFile()
	if err != nil {
		fmt.Printf("Error open the list of task\n")
		return
	}
	for _, name := range utils.CommandNames {
		fmt.Println(name)
	}

	//declare all commands
	name := flag.String("name", "dick", "The name to greet.")
	flag.Parse()
	fmt.Printf("Hello, %s!\n", *name)
}
