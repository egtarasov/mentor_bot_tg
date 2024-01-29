package bot

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegrambot_new_emploee/internal/models"
)

type telegramBot struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramBot(token string) (Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &telegramBot{
		bot: bot,
	}, nil
}

func (b *telegramBot) Start(ctx context.Context) <-chan *models.Update {
	updates := make(chan *models.Update)

	// Retrieve all updates and convert them to standard format.
	go func() {
		cfg := tgbotapi.NewUpdate(0)
		cfg.Timeout = 60
		ch := b.bot.GetUpdatesChan(cfg)
		for {
			select {
			case <-ctx.Done():
				return
			case update := <-ch:
				u := telegramToUpdate(&update)
				if u != nil {
					updates <- u
				}
			}
		}
	}()

	return updates
}

func telegramToUpdate(update *tgbotapi.Update) *models.Update {
	if update.Message == nil {
		return nil
	}

	return &models.Update{
		UpdateUserId: update.Message.From.ID,
		ChatId:       update.Message.Chat.ID,
		Message:      update.Message.Text,
	}
}

func (b *telegramBot) SendMessage(ctx context.Context, message models.Message) error {
	msg := tgbotapi.NewMessage(message.ChatId, message.Message)
	_, err := b.bot.Send(msg)
	if err != nil {
		return ErrMessageSend
	}

	return nil
}

func (b *telegramBot) SendButtons(ctx context.Context, buttons models.Buttons) error {
	var keywordButtons [][]tgbotapi.KeyboardButton
	for _, button := range buttons.Buttons {
		keywordButtons = append(keywordButtons,
			[]tgbotapi.KeyboardButton{
				tgbotapi.NewKeyboardButton(string(button)),
			})
	}

	msg := tgbotapi.NewMessage(buttons.ChatId, buttons.Message)
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(keywordButtons...)
	_, err := b.bot.Send(msg)

	return err
}
