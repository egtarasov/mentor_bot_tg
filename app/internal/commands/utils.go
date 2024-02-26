package commands

import (
	"context"
	"fmt"
	"strconv"
	container "telegrambot_new_emploee/internal/di-container"
	"telegrambot_new_emploee/internal/models"
)

const (
	CancelMessage = "Отмена"
)

var (
	ErrCanceled = fmt.Errorf("the command was canceled by the user")
)

// getNumber waits for a number from 1 to limit from the user. If the limit less or equal zero, there is no upper bound.
// Assuming that the message with the request for a number has been already sent. If 'Отмена' will be sent, cancel the
// processing with the corresponding error.
func getNumber(ctx context.Context, job *Job, limit int) (number int, err error) {
	for {
		update := job.Queue.WaitForUpdate()
		if update.Message == CancelMessage {
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

		if (limit > 0 && number >= limit) || number < 0 {
			err = container.Container.Bot().SendMessage(
				ctx,
				models.NewMessage(
					"Введенное число превышает допустимое значение, попробуйет снова!",
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

func getString(ctx context.Context, job *Job) (string, error) {
	update := job.Queue.WaitForUpdate()
	if update == nil {
		return "", ErrCanceled
	}

	return update.Message, nil
}
