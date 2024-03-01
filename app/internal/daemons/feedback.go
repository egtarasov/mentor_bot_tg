package daemons

import (
	"context"
	"telegrambot_new_emploee/internal/config"
	container "telegrambot_new_emploee/internal/di-container"
	"telegrambot_new_emploee/internal/models"
	"telegrambot_new_emploee/internal/views"
	"time"
)

func NewFeedbackDaemon(ctx context.Context) {
	startDaemon(ctx, config.Cfg.Feedback.Duration, feedback)
}

type daemonWork func(ctx context.Context) error

func startDaemon(ctx context.Context, duration time.Duration, work daemonWork) {
	err := work(ctx)
	if err != nil {
		// TODO Log error
	}
	ticker := time.NewTicker(duration)
	for {
		select {
		// TODO graceful shutdown.
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := work(ctx)
			if err != nil {
				// TODO Log error
			}
		}
	}
}

func feedback(ctx context.Context) error {
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
