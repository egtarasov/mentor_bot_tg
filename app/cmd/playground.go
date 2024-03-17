package main

import (
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
	bot, err := tgbotapi.NewBotAPI(config.Cfg.TgToken)
	if err != nil {
		panic(err)
	}
	bot.StopReceivingUpdates()
}

func Play() {
	b()

	time.Now()
}
