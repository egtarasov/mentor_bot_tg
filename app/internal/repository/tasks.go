package repository

import (
	"context"
	"telegrambot_new_emploee/internal/models"
)

type TasksRepo interface {
	GetTasks(ctx context.Context, employeeId int64) []models.Task
	GetTodoList(ctx context.Context, employeeId int64) []models.Todo
}
