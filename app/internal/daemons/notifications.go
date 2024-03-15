package daemons

import (
	"context"
	"log"
	"telegrambot_new_emploee/internal/config"
	"telegrambot_new_emploee/internal/convert"
	container "telegrambot_new_emploee/internal/di-container"
	"telegrambot_new_emploee/internal/models"
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

func StartNotificationDaemon(ctx context.Context) {
	notifications, err := container.Container.FAQRepo().GetNotifications(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, notification := range notifications {
		startNotificationDaemon(
			ctx,
			notification.NotificationTime,
			notification.DayOfWeek,
			notification.RepeatTime,
			newNotificationDaemon(notification.Message))
	}
}

func startNotificationDaemon(
	ctx context.Context,
	hour time.Duration,
	dayOfWeek int,
	repeat time.Duration,
	daemon daemon,
) {
	waitUntilTheCorrectTime(hour, dayOfWeek)
	startDaemon(ctx, repeat, daemon)
}

func waitUntilTheCorrectTime(hour time.Duration, dayOfWeek int) {
	now := time.Now()
	days := time.Duration((dayOfWeek - int(now.Weekday()) + 7) % 7)
	hours := convert.TimeToDuration(now)
	time.Sleep(days + (hour - hours))
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
