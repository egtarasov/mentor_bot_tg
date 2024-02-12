package repository

import (
	"context"
	"telegrambot_new_emploee/internal/models"
)

type TasksRepo interface {
	GetTasksById(ctx context.Context, employeeId int64) ([]models.Task, error)
	GetTodoListById(ctx context.Context, employeeId int64) ([]models.Todo, error)
	CheckTodo(ctx context.Context, todoId *models.Todo) error
	GetGoalsById(ctx context.Context, employeeId int64) ([]models.Goal, error)
}
