package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegrambot_new_emploee/internal/config"
	"time"
)

var (
	key = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("hello", "world"),
			tgbotapi.NewInlineKeyboardButtonData("hello 2", "world 2"),
			tgbotapi.NewInlineKeyboardButtonURL("you", "https://www.twitch.tv/smurf_tv"),
		),
	)
)

func b() {
	config.NewConfig()
	bot, _ := tgbotapi.NewBotAPI(config.Cfg.TgToken)

	cfg := tgbotapi.NewUpdate(0)
	cfg.Timeout = 60
	ch := bot.GetUpdatesChan(cfg)
	for update := range ch {
		id := update.Message.Chat.ID
		var files []any
		for _, photo := range update.Message.Photo {
			files = append(files, tgbotapi.NewInputMediaPhoto(tgbotapi.FileID(photo.FileID)))
		}
		media := tgbotapi.NewMediaGroup(id, files)
		m, err := bot.SendMediaGroup(media)
		if err != nil {
			fmt.Println(err.Error())
		}
		_, _ = m, err
	}
}

func Play() {
	b()

	time.Now()
}
