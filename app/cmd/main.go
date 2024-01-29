package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"telegrambot_new_emploee/internal/api"
)

const (
	TgBotToken = "6601547584:AAHrN_-ejSi4xeEFhLBAx9k2im-h9OYqJ4E"
)

var keyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("first"),
		tgbotapi.NewKeyboardButton("second")))

func main() {
	// TODO better configuration
	err := godotenv.Load("/Users/egtarasov/University/Projects/telegrambot_ne_employe/deploy/.env")
	if err != nil {
		log.Fatal(err)
	}

	api.Run()
}
