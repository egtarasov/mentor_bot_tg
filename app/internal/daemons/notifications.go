package daemons

import (
	"context"
	"telegrambot_new_emploee/internal/config"
	container "telegrambot_new_emploee/internal/di-container"
	"telegrambot_new_emploee/internal/models"
	"telegrambot_new_emploee/internal/views"
	"time"
)

type notificationDaemon struct {
	notificationMsg string
}

func newNotificationDaemon(message string) daemon {
	return &notificationDaemon{notificationMsg: message}
}

func (n *notificationDaemon) start(ctx context.Context) error {
	users, err := container.Container.UserRepo().GetUsersOnAdaptation(ctx)
	if err != nil {
		return err
	}
	for _, user := range users {
		if sendNotification(ctx, &user, n.notificationMsg) != nil {
			return err
		}
	}
	return nil
}

func NewTrainingDaemon(ctx context.Context) {
	startNotificationDaemon(
		ctx,
		config.Cfg.Notifications.TrainingHour,
		config.Cfg.Notifications.TrainingDayOfWeek,
		config.Cfg.Notifications.TrainingRepeat,
		newNotificationDaemon(views.TrainingNotification()),
	)
}

func NewHrMeetupDaemon(ctx context.Context) {
	startNotificationDaemon(
		ctx,
		config.Cfg.Notifications.HrMeetupHour,
		config.Cfg.Notifications.HrMeetupDayOfWeek,
		config.Cfg.Notifications.HrMeetupRepeat,
		newNotificationDaemon(views.HrMeetupNotification()),
	)
}

func NewMentorMeetupDaemon(ctx context.Context) {
	startNotificationDaemon(
		ctx,
		config.Cfg.Notifications.MentorMeetupHour,
		config.Cfg.Notifications.MentorMeetupDayOfWeek,
		config.Cfg.Notifications.MentorMeetupRepeat,
		newNotificationDaemon(views.MentorMeetupNotification()),
	)
}

func startNotificationDaemon(
	ctx context.Context,
	hour int,
	dayOfWeek int,
	repeat time.Duration,
	daemon daemon,
) {
	waitUntilTheCorrectTime(hour, dayOfWeek)
	startDaemon(ctx, repeat, daemon)
}

func waitUntilTheCorrectTime(hour int, dayOfWeek int) {
}

func sendNotification(ctx context.Context, user *models.User, message string) error {
	return container.Container.Bot().
		SendMessage(
			ctx,
			models.NewMessageWithPhotoPath(
				message,
				user.TelegramId,
				config.Cfg.Notifications.PhotoPath,
			),
		)
}
