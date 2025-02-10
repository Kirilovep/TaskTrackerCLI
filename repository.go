package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type TaskRepository interface {
	Save(task *Task) error
	Update(task *Task) error
	Delete(task *Task) error
	FileName(taskID string) string
}

type taskRepositoryImpl struct{}

func NewTaskRepository() TaskRepository {
	return &taskRepositoryImpl{}
}

func (r *taskRepositoryImpl) Save(task *Task) error {
	taskJSON, err := json.Marshal(task)

	if err != nil {
		return err
	}

	err = os.WriteFile(r.FileName(task.ID), taskJSON, 0666)

	if err != nil {
		return err
	}

	return nil
}

func (r *taskRepositoryImpl) Update(task *Task) error {
	taskJSON, err := json.Marshal(task)

	if err != nil {
		return err
	}

	err = os.WriteFile(r.FileName(task.ID), taskJSON, 0666)

	if err != nil {
		return err
	}

	return nil
}

func (r *taskRepositoryImpl) Delete(task *Task) error {
	return os.Remove(r.FileName(task.ID))
}

func (r *taskRepositoryImpl) FileName(taskID string) string {
	return fmt.Sprintf("%s.json", taskID)
}
