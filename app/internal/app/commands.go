package app

import (
	"context"
	"telegrambot_new_emploee/internal/bot"
)

type Cmd interface {
	Execute(ctx context.Context, job *job) error
}

type getMaterialCmd struct {
	app *app
}

type subDirCmd struct {
	app *app
}

type showTodoListCmd struct {
	app *app
}

func NewGetDataCmd(app *app) Cmd {
	return &getMaterialCmd{app: app}
}

func NewSubDirCmd(app *app) Cmd {
	return &subDirCmd{app: app}
}

func NewShowTodoListCmd(app *app) Cmd {
	return &showTodoListCmd{app: app}
}

func (c *getMaterialCmd) Execute(ctx context.Context, job *job) error {
	material, err := c.app.commandRepo.GetMaterials(ctx, job.command.Id)
	if err != nil {
		return err
	}

	message := bot.Message{
		Message: material.Message,
		ChatId:  job.update.ChatId,
	}
	return c.app.bot.SendMessage(ctx, message)
}

func (c *subDirCmd) Execute(ctx context.Context, job *job) error {
	buttons, err := c.app.commandRepo.GetCommands(ctx, job.command.Id)
	if err != nil {
		return err
	}

	material, err := c.app.commandRepo.GetMaterials(ctx, job.command.Id)
	if err != nil {
		return err
	}

	return c.app.bot.SendButtons(ctx, ToButtons(job.update.ChatId, buttons, material.Message))
}

func (c *showTodoListCmd) Execute(ctx context.Context, job *job) error {
	return nil
}
