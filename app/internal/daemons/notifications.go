package daemons

import (
	"context"
	"log"
	"telegrambot_new_emploee/internal/convert"
	container "telegrambot_new_emploee/internal/di-container"
	"telegrambot_new_emploee/internal/models"
	"time"
)

type notificationDaemon struct {
	notificationMsg string
	photoPath       *string
}

func newNotificationDaemon(message string, photoPath *string) daemon {
	return &notificationDaemon{
		notificationMsg: message,
		photoPath:       photoPath,
	}
}

func (n *notificationDaemon) start(ctx context.Context) error {
	users, err := container.Container.UserRepo().GetUsersOnAdaptation(ctx)
	if err != nil {
		return err
	}
	for _, user := range users {
		if n.sendNotification(ctx, &user) != nil {
			return err
		}
	}
	return nil
}

func (n *notificationDaemon) sendNotification(ctx context.Context, user *models.User) error {
	return container.Container.Bot().
		SendMessage(
			ctx,
			models.NewMessageWithPhotoPath(
				n.notificationMsg,
				user.TelegramId,
				n.photoPath,
			),
		)
}

func StartNotificationDaemon(ctx context.Context) {
	notifications, err := container.Container.FAQRepo().GetNotifications(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, notification := range notifications {
		go func(notification *models.Notification) {
			startNotificationDaemon(
				ctx,
				notification.NotificationTime,
				notification.DayOfWeek,
				notification.RepeatTime,
				newNotificationDaemon(notification.Message, notification.PhotoPath))
		}(&notification)
	}
}

func startNotificationDaemon(
	ctx context.Context,
	hour time.Duration,
	dayOfWeek int,
	repeat time.Duration,
	daemon daemon,
) {
	log.Printf("Create a notification:\n\t[hour:%v]\n\t[dayOfWeek:%v]\n\t[repeat:%v]", hour, dayOfWeek, repeat)
	waitUntilTheCorrectTime(hour, dayOfWeek)
	startDaemon(ctx, repeat, daemon)
}

func waitUntilTheCorrectTime(hour time.Duration, dayOfWeek int) {
	now := time.Now()
	days := time.Duration((dayOfWeek-int(now.Weekday())+7)%7) * day
	hours := convert.TimeToDuration(now)
	time.Sleep(days + (hour - hours))
}
