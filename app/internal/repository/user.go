package repository

import (
	"context"
	"fmt"
	"telegrambot_new_emploee/internal/models"
)

var (
	ErrNoUser = fmt.Errorf("no user with this tag")
)

type Task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	StoryPoints int    `json:"story_points"`
}

type Todo string

type AddTasks struct {
	EmployeeId int64  `json:"employee_id"`
	Tasks      []Task `json:"tasks"`
	Todos      []Todo `json:"todos"`
}

type UserRepo interface {
	GetUserByTag(ctx context.Context, tag int64) (*models.User, error)
	GetUserById(ctx context.Context, userId int64) (*models.User, error)
	GetUsersOnAdaptation(ctx context.Context) ([]models.User, error)
	AddTasks(ctx context.Context, tasks *AddTasks) error
}
