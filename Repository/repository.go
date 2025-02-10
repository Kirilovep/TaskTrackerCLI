package Repository

import (
	"encoding/json"
	"fmt"
	"os"
	"roadmap.sh/task-cli/Domain"
)

type Repository interface {
	LoadTasks()
	SaveTasks()
	AddTask(description string)
	UpdateTask(taskId int, description string)
	DeleteTask(taskId int)
	StartTask(taskId int)
	CompleteTask(taskId int)
	ListTasks()
	ListTasksByStatus(status string)
}

type repositoryImpl struct{}

var tasks []Domain.Task
var tasksFile = "tasks.json"

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r repositoryImpl) LoadTasks() {
	file, err := os.ReadFile(tasksFile)
	if err != nil {
		if os.IsNotExist(err) {
			tasks = []Domain.Task{}
			return
		}
		fmt.Println("Error reading tasks file:", err)
		os.Exit(1)
	}

	err = json.Unmarshal(file, &tasks)
	if err != nil {
		fmt.Println("Error parsing tasks file:", err)
		os.Exit(1)
	}
}

func (r repositoryImpl) ListTasksByStatus(status string) {
	found := false
	for _, task := range tasks {
		if task.Status == status {
			fmt.Printf("%d - %s [%s]\n", task.ID, task.Description, task.Status)
			found = true
		}
	}
	if !found {
		fmt.Printf("No tasks with status '%s' found\n", status)
	}
}

func (r repositoryImpl) ListTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}
	for _, task := range tasks {
		fmt.Printf("%d - %s [%s]\n", task.ID, task.Description, task.Status)
	}
}

func (r repositoryImpl) CompleteTask(taskId int) {
	for i, task := range tasks {
		if task.ID == taskId {
			tasks[i].Status = "done"
			fmt.Printf("Task %d is now done\n", taskId)
			return
		}
	}
	fmt.Println("Error: Task not found")
}

func (r repositoryImpl) StartTask(taskId int) {
	for i, task := range tasks {
		if task.ID == taskId {
			tasks[i].Status = "in_progress"
			fmt.Printf("Task %d is now in progress\n", taskId)
			return
		}
	}
	fmt.Println("Error: Task not found")
}

func (r repositoryImpl) SaveTasks() {
	file, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling tasks:", err)
		os.Exit(1)
	}

	err = os.WriteFile(tasksFile, file, 0644)
	if err != nil {
		fmt.Println("Error writing tasks file:", err)
		os.Exit(1)
	}
}

func (r repositoryImpl) AddTask(description string) {
	task := Domain.Task{
		ID:          len(tasks) + 1,
		Description: description,
		Status:      "todo",
	}
	tasks = append(tasks, task)
	fmt.Printf("Added task: %d - %s\n", task.ID, task.Description)
}

func (r repositoryImpl) UpdateTask(taskId int, description string) {
	for i, task := range tasks {
		if task.ID == taskId {
			tasks[i].Description = description
			fmt.Printf("Updated task: %d - %s\n", taskId, description)
			return
		}
	}
	fmt.Println("Error: Task not found")
}

func (r repositoryImpl) DeleteTask(taskId int) {
	for i, task := range tasks {
		if task.ID == taskId {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("Deleted task: %d\n", taskId)
			return
		}
	}
	fmt.Println("Error: Task not found")
}
