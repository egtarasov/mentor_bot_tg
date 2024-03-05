package daemons

import (
	"context"
	"telegrambot_new_emploee/internal/config"
	container "telegrambot_new_emploee/internal/di-container"
	"telegrambot_new_emploee/internal/models"
	"telegrambot_new_emploee/internal/views"
)

type feedbackDaemon struct {
}

func NewFeedbackDaemon(ctx context.Context) {
	startDaemon(ctx, config.Cfg.Feedback.Duration, &feedbackDaemon{})
}

func (f *feedbackDaemon) start(ctx context.Context) error {
	users, err := container.Container.UserRepo().GetUsersOnAdaptation(ctx)
	if err != nil {
		return err
	}
	for _, user := range users {
		if sendFeedbackForm(ctx, &user) != nil {
			return err
		}
	}
	return nil
}

func sendFeedbackForm(ctx context.Context, user *models.User) error {
	return container.Container.Bot().
		SendMessage(
			ctx,
			models.NewMessageWithPhotoPath(
				views.FeedBackForm(config.Cfg.Feedback.Form),
				user.TelegramId,
				config.Cfg.Feedback.PhotoPath,
			),
		)
}
