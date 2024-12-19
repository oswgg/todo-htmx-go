package repositories

import (
	"fmt"
	"github.com/oswgg/todo-htmx/internal/models"
)

type TaskRepository interface {
	FindById(id int64) (*models.Task, error)
	List() ([]*models.Task, error)
	Create(task *models.Task) ([]*models.Task, error)
	Update(task *models.Task) ([]*models.Task, error)
	Toggle(id int64) (*models.Task, error)
	Delete(id int64) error
}

type MockTaskRepository struct {
	tasks  []*models.Task
	nextID int64
}

func NewMockTaskRepository() TaskRepository {
	return &MockTaskRepository{
		tasks: []*models.Task{},
	}
}

func (m *MockTaskRepository) List() ([]*models.Task, error) {
	return m.tasks, nil
}
func (m *MockTaskRepository) FindById(id int64) (*models.Task, error) {
	for _, task := range m.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return nil, fmt.Errorf("task not found")
}

func (m *MockTaskRepository) Create(task *models.Task) ([]*models.Task, error) {
	task.ID = m.nextID
	m.tasks = append(m.tasks, task)
	m.nextID++
	return m.tasks, nil
}

func (m *MockTaskRepository) Update(task *models.Task) ([]*models.Task, error) {
	for i, itTask := range m.tasks {
		if itTask.ID == task.ID {
			m.tasks[i].Name = task.Name
			m.tasks[i].Completed = task.Completed
			return m.tasks, nil
		}
	}
	return nil, fmt.Errorf("task not found")
}

func (m *MockTaskRepository) Delete(id int64) error {
	for i, itTask := range m.tasks {
		if itTask.ID == id {
			m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task not found")
}

func (m *MockTaskRepository) Toggle(id int64) (*models.Task, error) {
	for i, itTask := range m.tasks {
		if itTask.ID == id {
			m.tasks[i].Completed = !m.tasks[i].Completed
		}
		return m.tasks[i], nil
	}
	return nil, fmt.Errorf("task not found")
}
