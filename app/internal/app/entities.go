package app

import (
	"telegrambot_new_emploee/internal/bot"
	"telegrambot_new_emploee/internal/database"
)

type Update struct {
	UserTag string
	ChatId  int64
	Message string
}

type User struct {
	UserId       int64
	Tag          string
	Name         string
	OccupationId int64
}

type Action string

const (
	GetDataCmd        Action = "get data"
	GetSubsectionsCmd Action = "show subsections"
	ComplexCmd        Action = "complex"
)

type Command struct {
	Id     int64
	Name   string
	Action Action
}

type Todo struct {
	Id         int64
	Label      string
	Priority   int
	EmployeeId int64
	Completed  bool
}

func ToTodo(todo *database.Todo) *Todo {
	return &Todo{
		Id:         todo.Id,
		Label:      todo.Label,
		Priority:   todo.Priority,
		EmployeeId: todo.EmployeeId,
		Completed:  todo.Completed,
	}
}

func ToUpdate(update *bot.Update) *Update {
	return &Update{
		UserTag: update.User.Tag,
		ChatId:  update.User.ChatId,
		Message: update.Message,
	}
}

func ToUser(user *database.User) *User {
	return &User{
		UserId:       user.Id,
		Tag:          user.TelegramTag,
		Name:         user.Name,
		OccupationId: user.OccupationId,
	}
}

func ToCommand(command *database.Command) *Command {
	return &Command{
		Id:     command.Id,
		Name:   command.Name,
		Action: Action(command.Action),
	}
}

func ToButtons(chatId int64, commands []database.Command, message string) bot.Buttons {
	res := bot.Buttons{
		ChatId:  chatId,
		Buttons: make([]bot.Button, 0, len(commands)),
		Message: message,
	}

	for _, command := range commands {
		res.Buttons = append(res.Buttons, bot.Button(command.Name))
	}

	return res
}
