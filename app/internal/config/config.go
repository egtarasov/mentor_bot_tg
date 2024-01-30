package config

import (
	"fmt"
	"os"
)

var Cfg *Config

const (
	tgToken = "TELEGRAM_TOKEN"
)

type Config struct {
	ConnStr string
	TgToken string
}

func connString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_SSLMODE"))
}

func NewConfig() error {
	val := os.Getenv("TELEGRAM_TOKEN")
	_ = val
	Cfg = &Config{
		ConnStr: connString(),
		TgToken: os.Getenv(tgToken),
	}
	return nil
}
