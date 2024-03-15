package commands

import (
	"context"
	container "telegrambot_new_emploee/internal/di-container"
	"telegrambot_new_emploee/internal/models"
	"telegrambot_new_emploee/internal/updates"
)

type Job struct {
	Ctx     context.Context
	Command *models.Command
	Update  *models.Update
	User    *models.User
	Queue   updates.Queue
}

func NewJob(
	ctx context.Context,
	queue updates.Queue,
	update *models.Update,
	user *models.User,
) (*Job, bool) {
	job := &Job{
		Ctx:    ctx,
		Update: update,
		User:   user,
		Queue:  queue,
	}

	if ok := job.getCommand(); !ok {
		return nil, false
	}

	return job, true
}

func (j *Job) GetChatId() int64 {
	return j.Update.ChatId
}

func (j *Job) getCommand() bool {
	command, err := container.Container.CmdRepo().GetCommand(j.Ctx, j.Update.Message)
	if err != nil {
		// TODO logging
		return false
	}
	if command.IsAdmin && !j.User.IsAdmin {
		return false
	}

	j.Command = command

	return true
}

type Cmd interface {
	Execute(ctx context.Context, job *Job) error
}
