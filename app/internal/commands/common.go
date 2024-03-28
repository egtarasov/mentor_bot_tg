package commands

import (
	"context"
	"errors"
	"telegrambot_new_emploee/internal/convert"
	container "telegrambot_new_emploee/internal/di-container"
	"telegrambot_new_emploee/internal/models"
	"telegrambot_new_emploee/internal/repository"
)

type getMaterialCmd struct {
}

type subDirCmd struct {
}

func NewGetDataCmd() Command {
	return &getMaterialCmd{}
}

func NewSubDirCmd() Command {
	return &subDirCmd{}
}

func (c *getMaterialCmd) Execute(ctx context.Context, job *Job) error {
	material, err := container.Container.CmdRepo().GetMaterials(ctx, job.Command.Id)
	if err != nil {
		return err
	}
	photoPath, err := container.Container.CmdRepo().GetImagePath(ctx, job.Command.Id)
	if err != nil && !errors.Is(err, repository.ErrNoImage) {
		return err
	}

	msg := models.NewMessage(material.Message, job.GetChatId())
	msg.PhotoPath = photoPath

	return container.Container.Bot().SendMessage(ctx, msg)
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

	photoPath, err := container.Container.CmdRepo().GetImagePath(ctx, job.Command.Id)
	if err != nil && !errors.Is(err, repository.ErrNoImage) {
		return err
	}

	buttons := convert.CommandsToButtons(commands, convert.DefaultToButtonsCfg(material.Message, job.GetChatId()))
	buttons.Message.PhotoPath = photoPath

	return container.Container.Bot().SendButtons(ctx, buttons)
}
