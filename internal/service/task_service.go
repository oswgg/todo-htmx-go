package service

import (
	"github.com/oswgg/todo-htmx/internal/models"
	"github.com/oswgg/todo-htmx/internal/repositories"
)

type TaskService interface {
	Create(task *models.Task) (*models.Task, error)
	Update(task *models.Task) ([]*models.Task, error)
	Delete(id int64) error
	Toggle(id int64) (*models.Task, error)
	List() ([]*models.Task, error)
}

type TaskServiceImpl struct {
	repo repositories.TaskRepository
}

func NewTaskService(repo repositories.TaskRepository) TaskService {
	return &TaskServiceImpl{
		repo: repo,
	}
}

func (s *TaskServiceImpl) Create(task *models.Task) (*models.Task, error) {
	newTaskList, err := s.repo.Create(task)
	if err != nil {
		return nil, err
	}
	return newTaskList, nil
}
func (s *TaskServiceImpl) Update(task *models.Task) ([]*models.Task, error) {
	updatedTaskList, err := s.repo.Update(task)
	if err != nil {
		return nil, err
	}
	return updatedTaskList, nil
}
func (s *TaskServiceImpl) Delete(id int64) error {
	return s.repo.Delete(id)
}
func (s *TaskServiceImpl) List() ([]*models.Task, error) {
	return s.repo.List()
}
func (s *TaskServiceImpl) Toggle(id int64) (*models.Task, error) {
	return s.repo.Toggle(id)
}
