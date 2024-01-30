package commands

// This file contains all the commands, which are available for a bot. The commands are responsible for processing
// the data

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"telegrambot_new_emploee/internal/convert"
	"telegrambot_new_emploee/internal/di-container"
	"telegrambot_new_emploee/internal/models"
)

type getMaterialCmd struct {
}

type subDirCmd struct {
}

type showTodoListCmd struct {
}

func NewGetDataCmd() Cmd {
	return &getMaterialCmd{}
}

func NewSubDirCmd() Cmd {
	return &subDirCmd{}
}

func NewShowTodoListCmd() Cmd {
	return &showTodoListCmd{}
}

func (c *getMaterialCmd) Execute(ctx context.Context, job *Job) error {
	material, err := container.Container.CmdRepo().GetMaterials(ctx, job.Command.Id)
	if err != nil {
		return err
	}

	return container.Container.Bot().
		SendMessage(
			ctx,
			models.NewMessage(material.Message, job.Update.ChatId),
		)
}

func (c *subDirCmd) Execute(ctx context.Context, job *Job) error {
	commands, err := container.Container.CmdRepo().GetCommands(ctx, job.Command.Id)
	if err != nil {
		return err
	}

	material, err := container.Container.CmdRepo().GetMaterials(ctx, job.Command.Id)
	if err != nil {
		return err
	}

	return container.Container.Bot().SendButtons(
		ctx,
		convert.ToButtons(commands, job.Update.ChatId, material.Message))
}

func (c *showTodoListCmd) Execute(ctx context.Context, job *Job) error {
	todos := container.Container.TaskRepo().GetTodoList(ctx, job.User.UserId)
	uncompletedTodos := make([]models.Todo, 0, len(todos))
	for _, todo := range todos {
		if todo.Completed {
			continue
		}
		uncompletedTodos = append(uncompletedTodos, todo)
	}

	msg := todoMessage(uncompletedTodos)

	return container.Container.Bot().SendMessage(ctx, models.NewMessage(msg, job.Update.ChatId))
}

func todoMessage(todos []models.Todo) string {
	var message strings.Builder
	message.WriteString("Список задач в твоем чек-листе:\n")

	sort.Slice(todos, func(i, j int) bool {
		return todos[i].Priority > todos[j].Priority
	})
	for i, todo := range todos {
		message.WriteString(fmt.Sprintf("%v. %s\n", i+1, todo.Label))
	}

	return message.String()
}
