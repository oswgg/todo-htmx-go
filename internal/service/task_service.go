package service

import (
	"github.com/oswgg/todo-htmx/internal/models"
	"github.com/oswgg/todo-htmx/internal/repositories"
)

type TaskService interface {
	Create()
	FindByID()
	Update()
	Delete()
}

type TaskServiceImpl struct {
	Repository repositories.TaskRepository
}

func NewTaskService(repository repositories.TaskRepository) *TaskService {
	var taskService TaskService

	taskService = &TaskServiceImpl{
		Repository: &repositories.MockTaskRepository{Tasks: make([]models.Task, 0)},
	}

	return &taskService
}

func (s *TaskServiceImpl) Create() {}

func (s *TaskServiceImpl) FindByID() {}

func (s *TaskServiceImpl) Update() {}

func (s *TaskServiceImpl) Delete() {}
