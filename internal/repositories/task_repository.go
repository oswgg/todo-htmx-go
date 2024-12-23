package repositories

import (
	"github.com/oswgg/todo-htmx/internal/models"
)

type TaskRepository interface {
	FindById(id int64) (*models.Task, error)
	List() ([]*models.Task, error)
	Create(task *models.Task) (*models.Task, error)
	Update(task *models.Task) ([]*models.Task, error)
	Toggle(id int64) (*models.Task, error)
	Delete(id int64) error
}
