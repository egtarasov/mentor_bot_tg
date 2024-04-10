package commands

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	container "telegrambot_new_emploee/internal/di-container"
	"telegrambot_new_emploee/internal/models"
)

const (
	CancelMessage = "отмена"
)

var (
	ErrCanceled = fmt.Errorf("the command was canceled by the user")
)

func getNumberWithMessage(ctx context.Context, job *Job, limit int, message string) (int, error) {
	if err := container.Container.Bot().SendMessage(ctx, models.NewMessage(message, job.GetChatId())); err != nil {
		return 0, err
	}
	return getNumber(ctx, job, limit)
}

func getStringWithMessage(ctx context.Context, job *Job, message string) (string, error) {
	if err := container.Container.Bot().SendMessage(ctx, models.NewMessage(message, job.GetChatId())); err != nil {
		return "", err
	}
	return getString(ctx, job)
}

// getNumber waits for a number from 1 to limit from the user. If the limit less or equal zero, there is no upper bound.
// Assuming that the message with the request for a number has been already sent. If 'Отмена' will be sent, cancel the
// processing with the corresponding error.
func getNumber(ctx context.Context, job *Job, limit int) (number int, err error) {
	for {
		update := job.Queue.WaitForUpdate()
		if isCanceled(update.Message) {
			return 0, ErrCanceled
		}
		number, err = strconv.Atoi(update.Message)
		if err != nil {
			err = container.Container.Bot().SendMessage(
				ctx,
				models.NewMessage(
					"Это не число, попробуйте снова!",
					job.GetChatId()))
			if err != nil {
				return 0, err
			}
			continue
		}

		if (limit > 0 && number > limit) || number <= 0 {
			err = container.Container.Bot().SendMessage(
				ctx,
				models.NewMessage(
					"Введенное число находится вне допустимого диапазона, попробуйет снова!",
					job.GetChatId()))
			if err != nil {
				return 0, err
			}
			continue
		}
		break
	}

	return number, err
}

func isCanceled(msg string) bool {
	return strings.ToLower(msg) == CancelMessage
}

func getString(ctx context.Context, job *Job) (string, error) {
	update := job.Queue.WaitForUpdate()
	if update == nil || isCanceled(update.Message) {
		return "", ErrCanceled
	}

	return update.Message, nil
}
