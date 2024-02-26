package commands

import (
	"context"
	container "telegrambot_new_emploee/internal/di-container"
	"telegrambot_new_emploee/internal/models"
	"telegrambot_new_emploee/internal/views"
)

type askQuestionCmd struct {
}

func NewAskQuestionCmd() Cmd {
	return &askQuestionCmd{}
}

func (c *askQuestionCmd) Execute(ctx context.Context, job *Job) error {
	err := container.Container.Bot().
		SendMessage(ctx, models.NewMessage(views.AskQuestionRequest, job.GetChatId()))
	if err != nil {
		return err
	}

	question, err := getString(ctx, job)
	if err != nil {
		return err
	}
	if question == CancelMessage {
		return ErrCanceled
	}

	id, err := container.Container.TaskRepo().
		CreateQuestion(ctx, models.NewQuestion(question, job.User.Id))
	if err != nil {
		return err
	}

	return container.Container.Bot().
		SendMessage(ctx, views.AskQuestionSuccess(id, question, job.GetChatId()))
}