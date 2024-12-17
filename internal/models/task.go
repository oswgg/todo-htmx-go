package models

import "time"

type Task struct {
	ID        int64
	Name      string
	Completed bool
	CreatedAt time.Time
}
