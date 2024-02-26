package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Config struct {
	ConnStr  string
	TgToken  string
	Feedback FeedBackConfig
}

type FeedBackConfig struct {
	Form      string
	Duration  time.Duration
	PhotoPath *string
}

var Cfg *Config

const (
	tgToken = "TELEGRAM_TOKEN"
)

// Yaml configuration.
type yamlConfig struct {
	Feedback *feedbackConfig `yaml:"feedback"`
}

type feedbackConfig struct {
	Form      string  `yaml:"form"`
	Duration  int     `yaml:"duration"`
	PhotoPath *string `yaml:"photo_path"`
}

func weeks(n int) time.Duration {
	return time.Duration(n) * 24 * 7 * time.Hour
}

func loadYamlConfig() (cfg *yamlConfig, err error) {
	file, err := os.Open("./config.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.NewDecoder(file).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, err
}

// Env configuration
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
	cfg, err := loadYamlConfig()
	if err != nil {
		return err
	}

	Cfg = &Config{
		ConnStr: connString(),
		TgToken: os.Getenv(tgToken),
		Feedback: FeedBackConfig{
			Form:     cfg.Feedback.Form,
			Duration: weeks(cfg.Feedback.Duration),
		},
	}
	return nil
}
