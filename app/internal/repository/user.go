package repository

import (
	"context"
	"fmt"
	"strings"
	"telegrambot_new_emploee/internal/models"
	"time"
)

var (
	ErrNoUser = fmt.Errorf("no user with this tag")
)

type Deadline time.Time

func (d *Deadline) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02 15:04", s)
	if err != nil {
		return err
	}
	*d = Deadline(t)
	return nil
}

type Task struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	StoryPoints int      `json:"story_points"`
	Deadline    Deadline `json:"deadLine"`
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
