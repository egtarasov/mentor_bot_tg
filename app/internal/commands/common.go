package commands

import (
	"context"
	"telegrambot_new_emploee/internal/convert"
	container "telegrambot_new_emploee/internal/di-container"
	"telegrambot_new_emploee/internal/models"
)

type getMaterialCmd struct {
}

type subDirCmd struct {
}

func NewGetDataCmd() Cmd {
	return &getMaterialCmd{}
}

func NewSubDirCmd() Cmd {
	return &subDirCmd{}
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
