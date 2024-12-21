package repositories

import (
	"database/sql"
	"github.com/oswgg/todo-htmx/internal/models"
)

type MariadbTaskRepository struct {
	db *sql.DB
}

func NewMariadbTaskRepository(db *sql.DB) TaskRepository {
	return &MariadbTaskRepository{
		db: db,
	}
}

func (m *MariadbTaskRepository) List() ([]*models.Task, error) {
	var tasks []*models.Task
	rows, err := m.db.Query("SELECT id, name, completed, created_at FROM tasks")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var task models.Task
		err = rows.Scan(&task.ID, &task.Name, &task.Completed, &task.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}
func (m *MariadbTaskRepository) FindById(id int64) (*models.Task, error) {
	var task models.Task
	row := m.db.QueryRow("SELECT id, name, completed, created_at FROM tasks WHERE id = ?", id)
	err := row.Scan(&task.ID, &task.Name, &task.Completed, &task.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (m *MariadbTaskRepository) Create(task *models.Task) (*models.Task, error) {
	row := m.db.QueryRow("INSERT INTO tasks(id, name, completed, created_at) VALUES(?, ?, ?, ?) RETURNING *", task.ID, task.Name, task.Completed, task.CreatedAt)
	err := row.Scan(&task.ID, &task.Name, &task.Completed, &task.CreatedAt)
	if err != nil {
		return nil, err
	}
	return task, nil
}

//func (m *MariadbTaskRepository) Update(task *models.Task) ([]*models.Task, error) {
//}
//
//func (m *MariadbTaskRepository) Delete(id int64) error {
//}
//
//func (m *MariadbTaskRepository) Toggle(id int64) (*models.Task, error) {
//}
