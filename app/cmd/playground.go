package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegrambot_new_emploee/internal/config"
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

func Play() {
	config.NewConfig()
	bot, _ := tgbotapi.NewBotAPI(config.Cfg.TgToken)

	cfg := tgbotapi.NewUpdate(0)
	cfg.Timeout = 60
	ch := bot.GetUpdatesChan(cfg)
	for update := range ch {
		id := update.Message.Chat.ID
		msg := tgbotapi.NewMessage(id, "Whta's up")
		msg.ReplyMarkup = key
		bot.Send(msg)
	}
}
