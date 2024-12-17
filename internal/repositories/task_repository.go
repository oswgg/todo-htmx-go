package repositories

import (
	"fmt"
	"github.com/oswgg/todo-htmx/internal/models"
)

type TaskRepository interface {
	Create(task *models.Task) error
	FindByID(id int64) (*models.Task, error)
	Update(task *models.Task) error
	Delete(id int64) error
}

type MockTaskRepository struct {
	Tasks []models.Task
}

func (r *MockTaskRepository) Create(task *models.Task) error {
	r.Tasks = append(r.Tasks, *task)
	return nil
}

func (r *MockTaskRepository) FindByID(id int64) (*models.Task, error) {
	var err error
	for _, task := range r.Tasks {
		if task.ID == id {
			return &task, nil
		}
	}

	err = fmt.Errorf("task with ID %v not found", id)
	return nil, err
}

func (r *MockTaskRepository) Update(task *models.Task) error {
	var err error
	for i, itTask := range r.Tasks {
		if itTask.ID == task.ID {
			r.Tasks[i].Completed = task.Completed
			r.Tasks[i].Name = task.Name

			return nil
		}
	}

	err = fmt.Errorf("task with ID %v not found", task.ID)
	return err
}

func (r *MockTaskRepository) Delete(id int64) error {
	var err error
	for i, itTask := range r.Tasks {
		if itTask.ID == id {
			r.Tasks = append(r.Tasks[:i], r.Tasks[i+1:]...)
			return nil
		}
	}
	err = fmt.Errorf("task with ID %v not found", id)
	return err
}
