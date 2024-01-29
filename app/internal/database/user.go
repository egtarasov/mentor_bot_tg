package database

import (
	"context"
	"fmt"
)

type User struct {
	Id           int64  `db:"id"`
	Name         string `db:"name"`
	Surname      string `db:"surname"`
	TelegramTag  string `db:"telegram_tag"`
	OccupationId int64  `db:"occupation_id"`
}

var (
	ErrNoUser = fmt.Errorf("no user with this tag")
)

type UserRepo interface {
	GetUserByTag(ctx context.Context, tag string) (*User, error)
}
