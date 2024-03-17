package bot

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sync"
	"telegrambot_new_emploee/internal/models"
	"time"
)

const maxTextSize = 4096
const maxCaptionSize = 1024

type telegramBot struct {
	bot       *tgbotapi.BotAPI
	parseMode string
	ch        <-chan tgbotapi.Update
	updates   chan *models.Update

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

func (b *telegramBot) newPhoto(message *models.Message, markUp any) ([]tgbotapi.Chattable, bool) {
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
	photo.ParseMode = b.parseMode
	photo.ReplyMarkup = markUp
	photo.Caption = message.Message[:min(maxCaptionSize, len(message.Message))]
	chattables := []tgbotapi.Chattable{photo}
	return b.parseText(message.Message[min(maxCaptionSize, len(message.Message)):], message.ChatId, chattables), flag
}

func (b *telegramBot) parseText(text string, chatId int64, chattables []tgbotapi.Chattable) []tgbotapi.Chattable {
	for len(text) > 0 {
		msg := tgbotapi.NewMessage(chatId, text[:min(maxTextSize, len(text))])
		msg.ParseMode = b.parseMode
		chattables = append(chattables, msg)
		text = text[min(maxTextSize, len(text)):]
	}
	return chattables
}

func (b *telegramBot) newMessage(message *models.Message, markUp any) []tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(message.ChatId, message.Message[:min(maxTextSize, len(message.Message))])
	msg.ParseMode = b.parseMode
	msg.ReplyMarkup = markUp
	chattables := []tgbotapi.Chattable{msg}
	return b.parseText(message.Message[min(maxTextSize, len(message.Message)):], message.ChatId, chattables)
}

func (b *telegramBot) newMediaGroup(message *models.Message) []tgbotapi.Chattable {
	var group []any
	for i, id := range message.PhotoIds {
		photo := tgbotapi.NewInputMediaPhoto(tgbotapi.FileID(id))
		if i == 0 {
			photo.Caption = message.Message[:min(maxCaptionSize, len(message.Message))]
		}
		group = append(group, photo)
	}
	chattables := []tgbotapi.Chattable{tgbotapi.NewMediaGroup(message.ChatId, group)}
	return b.parseText(message.Message[min(maxCaptionSize, len(message.Message)):], message.ChatId, chattables)
}

func (b *telegramBot) newChattable(message *models.Message, markup any) ([]tgbotapi.Chattable, bool) {
	if message.PhotoBytes != nil || message.PhotoPath != nil {
		return b.newPhoto(message, markup)
	}
	if message.PhotoIds != nil {
		return b.newMediaGroup(message), true
	}
	return b.newMessage(message, markup), true
}

func (b *telegramBot) Start(ctx context.Context) <-chan *models.Update {
	b.updates = make(chan *models.Update)

	// Retrieve all updates and convert them to standard format.
	go func() {
		cfg := tgbotapi.NewUpdate(0)
		cfg.Timeout = 30
		b.ch = b.bot.GetUpdatesChan(cfg)
		for {
			select {
			case <-ctx.Done():
				return
			case update := <-b.ch:
				b.telegramToUpdate(ctx, &update)
			}
		}
	}()

	return b.updates
}

func (b *telegramBot) getMediaGroup(ctx context.Context, update *models.Update, groupId string) *tgbotapi.Update {
	for {
		ctx, cancel := context.WithTimeout(ctx, time.Millisecond*5)
		defer cancel()

		select {
		case <-ctx.Done():
			return nil
		case u := <-b.ch:
			if u.Message == nil {
				return nil
			}
			if u.Message.MediaGroupID != groupId {
				return &u
			}
			if u.Message.Photo != nil {
				update.PhotoIds = append(update.PhotoIds, u.Message.Photo[0].FileID)
			}
		}
	}
}

func defaultUpdate(update *tgbotapi.Update) *models.Update {
	// Create an update.
	u := &models.Update{
		UpdateUserId: update.Message.From.ID,
		ChatId:       update.Message.Chat.ID,
		Message:      update.Message.Text,
	}

	// Process a Photo.
	if update.Message.Caption != "" {
		u.Message = update.Message.Caption
	}
	var photoIds []string
	if update.Message.Photo != nil {
		photoIds = append(photoIds, update.Message.Photo[0].FileID)
	}
	u.PhotoIds = photoIds

	return u
}

func (b *telegramBot) telegramToUpdate(ctx context.Context, update *tgbotapi.Update) {
	if update == nil || update.Message == nil {
		return
	}
	u := defaultUpdate(update)
	if update.Message.MediaGroupID != "" {
		newUpdate := b.getMediaGroup(ctx, u, update.Message.MediaGroupID)
		b.sendUpdate(u)
		b.telegramToUpdate(ctx, newUpdate)
		return
	}

	b.sendUpdate(u)
}

func (b *telegramBot) sendUpdate(update *models.Update) {
	// Last validation of the update.
	if update.Message == "" && len(update.PhotoIds) == 0 {
		return
	}

	b.updates <- update
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

type Message struct {
	tgbotapi.Message
}

func (b *telegramBot) send(message *models.Message, markup any) error {
	messages, ok := b.newChattable(message, markup)
	for i, msg := range messages {
		m, err := b.bot.Send(msg)
		if err != nil {
			return err
		}
		if !ok && i == 0 {
			b.storeFileId(message.FilePath(), getFileId(&m))
		}
	}
	return nil
}
