package bot

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sync"
	"telegrambot_new_emploee/internal/models"
)

type telegramBot struct {
	bot       *tgbotapi.BotAPI
	parseMode string

	lock          sync.RWMutex
	photosStorage map[string]string
}

func (b *telegramBot) getFileId(path string) (string, bool) {
	b.lock.RLock()
	defer b.lock.RUnlock()
	fileId, ok := b.photosStorage[path]
	return fileId, ok
}

func (b *telegramBot) storeFileId(path string, fileId string) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.photosStorage[path] = fileId
}

func NewTelegramBot(token string, parseMode string) (Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &telegramBot{
		bot:           bot,
		parseMode:     parseMode,
		photosStorage: make(map[string]string),
	}, nil
}

func (b *telegramBot) newPhoto(message *models.Message, markUp any) (*tgbotapi.PhotoConfig, bool) {
	var file tgbotapi.RequestFileData
	flag := true
	switch {
	case message.PhotoPath != nil:
		fileId, ok := b.getFileId(*message.PhotoPath)
		if !ok {
			flag = false
			file = tgbotapi.FilePath(*message.PhotoPath)
		} else {
			file = tgbotapi.FileID(fileId)
		}
	case message.PhotoBytes != nil:
		file = tgbotapi.FileBytes{
			Name:  "",
			Bytes: message.PhotoBytes,
		}
	default:
		return nil, flag
	}
	photo := tgbotapi.NewPhoto(message.ChatId, file)
	photo.Caption = message.Message
	photo.ParseMode = b.parseMode
	photo.ReplyMarkup = markUp
	return &photo, flag
}

func (b *telegramBot) newMessage(message *models.Message, markUp any) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(message.ChatId, message.Message)
	msg.ParseMode = b.parseMode
	msg.ReplyMarkup = markUp
	return &msg
}

func (b *telegramBot) newChattable(message *models.Message, markup any) (tgbotapi.Chattable, bool) {
	if message.PhotoBytes != nil || message.PhotoPath != nil {
		return b.newPhoto(message, markup)
	}
	return b.newMessage(message, markup), true
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

func getFileId(msg *tgbotapi.Message) string {
	return msg.Photo[0].FileID
}

func (b *telegramBot) SendMessage(ctx context.Context, message *models.Message) error {
	return b.send(message, nil)
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
	return b.send(buttons.Message, tgbotapi.NewReplyKeyboard(keywordButtons...))
}

func (b *telegramBot) send(message *models.Message, markup any) error {
	msg, ok := b.newChattable(message, markup)
	m, err := b.bot.Send(msg)
	if !ok {
		b.storeFileId(message.FilePath(), getFileId(&m))
	}
	if err != nil {
		return ErrMessageSend
	}
	return nil
}
