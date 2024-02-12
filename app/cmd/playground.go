package main

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"telegrambot_new_emploee/internal/config"
	"telegrambot_new_emploee/internal/repository/postgres"
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
		msg := tgbotapi.NewMessage(id, "Whta's up")
		msg.ReplyMarkup = key
		bot.Send(msg)
	}
}

func images() {
	err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, config.Cfg.ConnStr)
	if err != nil {
		log.Fatal(err)
	}

	rep := postgres.NewCommandPostgres(pool)
	path, err := rep.GetImagePath(ctx, 1)

	bot, _ := tgbotapi.NewBotAPI(config.Cfg.TgToken)
	fmt.Println(os.Getwd())

	cfg := tgbotapi.NewUpdate(0)
	cfg.Timeout = 60
	ch := bot.GetUpdatesChan(cfg)
	for update := range ch {
		//msg := tgbotapi.NewMessage(update.Message.From.ID, "hello")
		photo := tgbotapi.NewPhoto(update.Message.From.ID, tgbotapi.FilePath(path))
		photo.Caption = "hello my friend"
		up, err := bot.Send(photo)
		if err != nil {
			log.Fatal(err)
		}
		_, _ = up, err
	}
}

func Play() {
	images()
}
