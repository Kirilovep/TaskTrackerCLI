package main

import (
	"encoding/json"
	"os"
	"strconv"
	"time"
)

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (service *TaskService) AddTask(taskName string) (int, error) {
	taskId := 1

	for {
		_, err := os.ReadFile(strconv.Itoa(taskId))

		if err != nil {
			break
		}

		taskId += 1
	}

	task := NewTask(strconv.Itoa(taskId), taskName)

	if err := service.repo.Save(task); err != nil {
		return 0, err
	}

	return taskId, nil
}

func (service *TaskService) UpdateTask(taskId string, taskName string) error {
	byteValue, err := os.ReadFile(service.repo.FileName(taskId))

	if err != nil {
		return err
	}

	var task *Task

	if err := json.Unmarshal(byteValue, &task); err != nil {
		return err
	}

	task.Name = taskName
	setCurrentUpdatedAtTime(task)

	return service.repo.Update(task)
}

func (service *TaskService) SetInProgressTask(taskId string) error {
	byteValue, err := os.ReadFile(service.repo.FileName(taskId))

	if err != nil {
		return err
	}

	var task *Task

	if err := json.Unmarshal(byteValue, &task); err != nil {
		return err
	}

	task.InProgress = true
	setCurrentUpdatedAtTime(task)

	return service.repo.Update(task)
}

func (service *TaskService) SetDoneTask(taskId string) error {
	byteValue, err := os.ReadFile(taskId)

	if err != nil {
		return err
	}

	var task *Task

	if err := json.Unmarshal(byteValue, &task); err != nil {
		return err
	}

	task.InProgress = false
	task.IsDone = true
	setCurrentUpdatedAtTime(task)

	return service.repo.Update(task)
}

func (service *TaskService) DeleteTask(taskId string) error {
	byteValue, err := os.ReadFile(taskId)

	if err != nil {
		return err
	}

	var task *Task

	if err := json.Unmarshal(byteValue, &task); err != nil {
		return err
	}

	return service.repo.Delete(task)
}

func setCurrentUpdatedAtTime(task *Task) {
	task.UpdatedAt = time.Now()
}
