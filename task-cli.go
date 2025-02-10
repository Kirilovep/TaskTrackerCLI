package main

import (
	"fmt"
	"os"
	"roadmap.sh/task-cli/Repository"
	"roadmap.sh/task-cli/Service"
	"strings"
)

func main() {
	repository := Repository.NewRepository()
	service := Service.NewService(repository)

	service.LoadTasks()

	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing task description")
			return
		}
		description := strings.Join(os.Args[2:], " ")
		service.AddTask(description)

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Error: Missing task ID or new description")
			return
		}
		id := parseInt(os.Args[2])
		description := strings.Join(os.Args[3:], " ")
		service.UpdateTask(id, description)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing task ID")
			return
		}
		id := parseInt(os.Args[2])
		service.DeleteTask(id)

	case "start":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing task ID")
			return
		}
		id := parseInt(os.Args[2])
		service.StartTask(id)

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing task ID")
			return
		}
		id := parseInt(os.Args[2])
		service.CompleteTask(id)

	case "list":
		service.ListTasks()

	case "list-done":
		service.ListTasksByStatus("done")

	case "list-todo":
		service.ListTasksByStatus("todo")

	case "list-in-progress":
		service.ListTasksByStatus("in_progress")

	default:
		fmt.Println("Error: Unknown command")
		printUsage()
	}

	service.SaveTasks()
}

// Parse an integer from a string
func parseInt(s string) int {
	var id int
	_, err := fmt.Sscanf(s, "%d", &id)
	if err != nil {
		fmt.Println("Error: Invalid task ID")
		os.Exit(1)
	}
	return id
}

// Print usage instructions
func printUsage() {
	fmt.Println(`Usage:
  task-tracker <command> [arguments]

Commands:
  add <description>       Add a new task
  update <id> <description> Update a task
  delete <id>            Delete a task
  start <id>             Mark a task as in progress
  done <id>              Mark a task as done
  list                   List all tasks
  list-done              List all done tasks
  list-todo              List all todo tasks
  list-in-progress       List all in-progress tasks`)
}
