package database

import "context"

type Task struct {
	Id          int64  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	StoryPoints int64  `db:"story_points"`
	Completed   bool   `db:"completed"`
	EmployeeId  int64  `db:"employee_id"`
}

type TasksRepo interface {
	GetTasks(ctx context.Context, employeeId int64) []Task
}
