package repository

import (
	"context"
	"fmt"
	"telegrambot_new_emploee/internal/models"
)

var (
	ErrNoData = fmt.Errorf("")
)

type TasksRepo interface {
	GetTasksById(ctx context.Context, employeeId int64) ([]models.Task, error)
	GetTodoListById(ctx context.Context, employeeId int64) ([]models.Todo, error)
	CheckTodo(ctx context.Context, todo *models.Todo) error
	CheckTask(ctx context.Context, task *models.Task) error
	GetGoalsByUser(ctx context.Context, user *models.User) ([]models.Goal, error)
	GetOccupationMaterial(ctx context.Context, occupationId int64) (*models.Occupation, error)
}
