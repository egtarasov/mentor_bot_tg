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
	Tasks    TasksConfig
	Admin    AdminConfig
}

type AdminConfig struct {
	Port string
}

type FeedBackConfig struct {
	Form      string
	Duration  time.Duration
	PhotoPath *string
}

type TasksConfig struct {
	PhotoPathGoals *string
	PhotoPathTasks *string
	PhotoPathTodos *string
}

var Cfg *Config

const (
	tgToken = "TELEGRAM_TOKEN"
)

// Yaml configuration.
type yamlConfig struct {
	Feedback feedbackConfig `yaml:"feedback"`
	Tasks    tasksConfig    `yaml:"tasks"`
	Admin    adminConfig    `yaml:"admin"`
}

type adminConfig struct {
	Port string `yaml:"port"`
}

type feedbackConfig struct {
	Form      string  `yaml:"form"`
	Duration  int     `yaml:"duration"`
	PhotoPath *string `yaml:"photo_path"`
}

type tasksConfig struct {
	PhotoPathGoals *string `yaml:"photo_path_goals"`
	PhotoPathTasks *string `yaml:"photo_path_tasks"`
	PhotoPathTodos *string `yaml:"photo_path_todos"`
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
			Form:      cfg.Feedback.Form,
			Duration:  weeks(cfg.Feedback.Duration),
			PhotoPath: cfg.Feedback.PhotoPath,
		},
		Tasks: TasksConfig{
			PhotoPathGoals: cfg.Tasks.PhotoPathGoals,
			PhotoPathTasks: cfg.Tasks.PhotoPathTasks,
			PhotoPathTodos: cfg.Tasks.PhotoPathTodos,
		},
		Admin: AdminConfig{
			Port: cfg.Admin.Port,
		},
	}

	return nil
}
