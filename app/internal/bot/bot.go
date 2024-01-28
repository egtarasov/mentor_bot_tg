package bot

import (
	"context"
	"fmt"
)

var (
	ErrMessageSend = fmt.Errorf("can't send the message")
)

type Update struct {
	User    *User
	Message string
}

type User struct {
	ChatId int64
	Tag    string
}

type Message struct {
	Message string
	ChatId  int64
}

type Button string

type Buttons struct {
	ChatId  int64
	Buttons []Button
	Message string
}

type Bot interface {
	// Start starts processing the incoming updates to bot.
	// Returns the chan of processed updates.
	Start(ctx context.Context) <-chan *Update

	SendMessage(ctx context.Context, message Message) error

	SendButtons(ctx context.Context, buttons Buttons) error
}
