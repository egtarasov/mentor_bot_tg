package bot

import (
	"context"
	"fmt"
	"telegrambot_new_emploee/internal/models"
)

var (
	ErrMessageSend = fmt.Errorf("can't send the message")
)

type Bot interface {
	// Start starts processing the incoming updates to bot.
	// Returns the chan of processed updates.
	Start(ctx context.Context) <-chan *models.Update

	// SendMessage sends the given message to the user.
	SendMessage(ctx context.Context, message *models.Message) error

	// SendButtons sends the message and a new keyboard to the user.
	SendButtons(ctx context.Context, buttons *models.Buttons) error
}
