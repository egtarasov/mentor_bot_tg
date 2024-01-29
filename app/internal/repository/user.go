package repository

import (
	"context"
	"fmt"
	"telegrambot_new_emploee/internal/models"
)

var (
	ErrNoUser = fmt.Errorf("no user with this tag")
)

type UserRepo interface {
	GetUserByTag(ctx context.Context, userId int64) (*models.User, error)
}