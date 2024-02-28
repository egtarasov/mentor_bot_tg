package bot

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegrambot_new_emploee/internal/models"
)

type telegramBot struct {
	bot       *tgbotapi.BotAPI
	parseMode string
}

func NewTelegramBot(token string, parseMode string) (Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &telegramBot{
		bot:       bot,
		parseMode: parseMode,
	}, nil
}

func (b *telegramBot) newPhoto(message *models.Message, markUp any) *tgbotapi.PhotoConfig {
	photo := tgbotapi.NewPhoto(message.ChatId, tgbotapi.FilePath(*message.PhotoPath))
	photo.Caption = message.Message
	photo.ParseMode = b.parseMode
	photo.ReplyMarkup = markUp
	return &photo
}

func (b *telegramBot) newMessage(message *models.Message, markUp any) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(message.ChatId, message.Message)
	msg.ParseMode = b.parseMode
	msg.ReplyMarkup = markUp
	return &msg
}

func (b *telegramBot) newChattable(message *models.Message, markup any) tgbotapi.Chattable {
	if message.PhotoPath != nil {
		return b.newPhoto(message, markup)
	}
	return b.newMessage(message, markup)
}

func (b *telegramBot) Start(ctx context.Context) <-chan *models.Update {
	updates := make(chan *models.Update)

	// Retrieve all updates and convert them to standard format.
	go func() {
		cfg := tgbotapi.NewUpdate(0)
		cfg.Timeout = 30
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

func (b *telegramBot) SendMessage(ctx context.Context, message *models.Message) error {
	msg := b.newChattable(message, nil)
	_, err := b.bot.Send(msg)
	if err != nil {
		return ErrMessageSend
	}
	return nil
}

func (b *telegramBot) SendButtons(ctx context.Context, buttons *models.Buttons) error {
	var keywordButtons [][]tgbotapi.KeyboardButton
	for _, row := range buttons.Buttons {
		keyBoardRow := make([]tgbotapi.KeyboardButton, 0, len(row))
		for _, button := range row {
			keyBoardRow = append(keyBoardRow, tgbotapi.NewKeyboardButton(string(button)))
		}
		keywordButtons = append(keywordButtons, keyBoardRow)
	}
	msg := b.newChattable(buttons.Message, tgbotapi.NewReplyKeyboard(keywordButtons...))
	_, err := b.bot.Send(msg)

	return err
}
