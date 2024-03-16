package daemons

import (
	"context"
	"log"
	container "telegrambot_new_emploee/internal/di-container"
	"telegrambot_new_emploee/internal/models"
	"telegrambot_new_emploee/internal/views"
	"time"
)

const day = time.Hour * 24

type taskReminderDaemon struct {
}

func StartTasksNotificationDaemon(ctx context.Context) {
	startDaemon(ctx, day, &taskReminderDaemon{})
}

func (t *taskReminderDaemon) start(ctx context.Context) error {
	users, err := container.Container.UserRepo().GetUsersOnAdaptation(ctx)
	if err != nil {
		return err
	}
	for _, user := range users {
		tasks, err := container.Container.TaskRepo().GetTasksById(ctx, user.Id)
		if err != nil {
			return err
		}
		var taskToNotify []models.Task
		for _, task := range tasks {
			if task.Deadline == nil || task.CompletedAt != nil || (task.Deadline.Sub(time.Now()) > day) {
				continue
			}
			taskToNotify = append(taskToNotify, task)
		}
		if len(taskToNotify) == 0 {
			continue
		}
		err = container.Container.Bot().SendMessage(ctx, views.TasksDeadlineNotification(taskToNotify, user.TelegramId))
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
