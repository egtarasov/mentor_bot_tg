package models

import (
	"database/sql"
	"time"
)

type Command struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	ActionId int    `db:"action_id"`
	ParentId *int64 `db:"parent_id"`
	IsAdmin  bool   `db:"is_admin"`
}

type Material struct {
	Id        int64  `db:"id"`
	Message   string `db:"message"`
	CommandId int64  `db:"command_id"`
}

type Task struct {
	Id          int64               `db:"id"`
	Name        string              `db:"name"`
	Description string              `db:"description"`
	StoryPoints int64               `db:"story_points"`
	EmployeeId  int64               `db:"employee_id"`
	CreatedAt   time.Time           `db:"created_at"`
	CompletedAt sql.Null[time.Time] `db:"completed_at"`
}

type Todo struct {
	Id         int64  `db:"id"`
	Label      string `db:"label"`
	Priority   int    `db:"priority"`
	EmployeeId int64  `db:"employee_id"`
	Completed  bool   `db:"completed"`
}

type User struct {
	Id             int64     `db:"id"`
	Name           string    `db:"name"`
	Surname        string    `db:"surname"`
	TelegramId     int64     `db:"telegram_id"`
	OccupationId   int64     `db:"occupation_id"`
	IsAdmin        bool      `db:"is_admin"`
	Grade          int       `db:"grade"`
	StartWork      time.Time `db:"first_working_day"`
	AdaptationEnds time.Time `db:"adaptation_end_at"`
}

type Goal struct {
	Id          int64  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	EmployeeId  int64  `db:"employee_id"`
	Track       string `db:"track"`
}

type Question struct {
	Id         int64               `db:"id"`
	UserId     int64               `db:"user_id"`
	Text       string              `db:"text"`
	CreatedAt  time.Time           `db:"created_at"`
	AnsweredAt sql.Null[time.Time] `db:"answered_at"`
	AnsweredBy sql.Null[int64]     `db:"answered_by"`
	Answer     sql.Null[string]    `db:"answer"`
}

type Occupation struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	Material string `db:"material"`
}

type CommandWithMaterial struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	Message  string `db:"message"`
	ActionId int64  `db:"action_id"`
}
