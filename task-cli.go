package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Empty command")
		return
	}

	if len(os.Args) < 3 {
		fmt.Println("Empty command input")
		return
	}

	repo := NewTaskRepository()
	service := NewTaskService(repo)

	command := os.Args[1]

	switch command {
	case "add":
		taskId, err := service.AddTask(os.Args[2])

		if err != nil {
			fmt.Printf("%s", err.Error())
		} else {
			fmt.Println("Task added successfully", "ID:", taskId)
		}

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Empty task id or new task name")
			return
		}

		err := service.UpdateTask(os.Args[2], os.Args[3])

		if err != nil {
			fmt.Println("Task id incorrect")
		}

	case "delete":
		err := service.DeleteTask(os.Args[2])

		if err != nil {
			fmt.Printf("%s", err.Error())
		}

	case "mark-in-progress":
		err := service.SetInProgressTask(os.Args[2])

		if err != nil {
			fmt.Printf("%s", err.Error())
		}

	case "mark-done":
		err := service.SetDoneTask(os.Args[2])

		if err != nil {
			fmt.Printf("%s", err.Error())
		}
	// case "list-all":
	// 	ListAllTasks()
	// case "list-done":
	// 	ListDoneTasks()
	// case "list-not-done":
	// 	ListNotDoneTasks()
	// case "list-in-progress":
	// 	ListInProgressTasks()
	default:
		fmt.Println("Unknown command. Available commands: add, update, delete, mark-in-progress, mark-done, list-all, list-done, list-not-done, list-in-progress")
	}
}
