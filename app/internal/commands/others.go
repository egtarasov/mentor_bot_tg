package commands

import (
	"context"
	"strings"
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
	if strings.ToLower(question) == CancelMessage {
		return ErrCanceled
	}

	id, err := container.Container.QuestionRepo().
		CreateQuestion(ctx, models.NewQuestion(question, job.User.Id))
	if err != nil {
		return err
	}

	return container.Container.Bot().
		SendMessage(ctx, views.AskQuestionSuccess(id, question, job.GetChatId()))
}

type calendarCmd struct {
}

func NewCalendarCmd() Cmd {
	return &calendarCmd{}
}

func (c *calendarCmd) Execute(ctx context.Context, job *Job) error {
	meetings, err := container.Container.Calendar().GetMeetingsById(job.User)
	if err != nil {
		return err
	}

	return container.Container.Bot().SendMessage(
		ctx,
		views.CalendarMessage(meetings, job.GetChatId()),
	)
}
