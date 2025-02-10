package main

import "time"

type Task struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	InProgress bool      `json:"inProgress"`
	IsDone     bool      `json:"isDone"`
}

func NewTask(id string, name string) *Task {
	return &Task{
		ID:         id,
		Name:       name,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		InProgress: false,
		IsDone:     false,
	}
}
