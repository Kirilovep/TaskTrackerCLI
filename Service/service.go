package Service

import "roadmap.sh/task-cli/Repository"

type Service interface {
	LoadTasks()
	SaveTasks()
	AddTask(description string)
	UpdateTask(taskId int, description string)
	StartTask(taskId int)
	DeleteTask(taskId int)
	CompleteTask(taskId int)
	ListTasksByStatus(status string)
	ListTasks()
}

type serviceImpl struct {
	repository Repository.Repository
}

func NewService(repository Repository.Repository) Service {
	return &serviceImpl{
		repository: repository,
	}
}

func (s serviceImpl) LoadTasks() {
	s.repository.LoadTasks()
}

func (s serviceImpl) ListTasks() {
	s.repository.ListTasks()
}

func (s serviceImpl) SaveTasks() {
	s.repository.SaveTasks()
}

func (s serviceImpl) AddTask(description string) {
	s.repository.AddTask(description)
}

func (s serviceImpl) UpdateTask(taskId int, description string) {
	s.repository.UpdateTask(taskId, description)
}

func (s serviceImpl) DeleteTask(taskId int) {
	s.repository.DeleteTask(taskId)
}

func (s serviceImpl) StartTask(taskId int) {
	s.repository.StartTask(taskId)
}

func (s serviceImpl) CompleteTask(taskId int) {
	s.repository.CompleteTask(taskId)
}

func (s serviceImpl) ListTasksByStatus(status string) {
	s.repository.ListTasksByStatus(status)
}
