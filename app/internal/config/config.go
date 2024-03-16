package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

func parseDayOfWeek(day string) int {
	dayToInt := map[string]int{
		"Mon": 1,
		"Tue": 2,
		"Wed": 3,
		"Thu": 4,
		"Fri": 5,
		"Sat": 6,
		"Sun": 7,
	}

	dayOfWeek, ok := dayToInt[day]
	if !ok {
		panic(fmt.Sprintf("invalid day of the week in configuration: [%s]", day))
	}
	return dayOfWeek
}

type Config struct {
	ConnStr     string
	TgToken     string
	Tasks       TasksConfig
	Admin       AdminConfig
	CalendarUrl *string
}
type AdminConfig struct {
	Port           string
	MaxPhotoSize   int64
	PhotoFormKey   string
	MessageFormKey string
	FAQCommandName string
}

type TasksConfig struct {
	PhotoPathGoals *string
	PhotoPathTasks *string
	PhotoPathTodos *string
	BarCount       int
}

var Cfg *Config

const (
	tgToken = "TELEGRAM_TOKEN"
)

// Yaml configuration.
type yamlConfig struct {
	Tasks    tasksConfig    `yaml:"tasks"`
	Admin    adminConfig    `yaml:"admin"`
	Calendar calendarConfig `yaml:"calendar"`
}

type calendarConfig struct {
	Url *string `yaml:"url"`
}

type adminConfig struct {
	Port           string `yaml:"port"`
	MaxPhotoSize   int64  `yaml:"max_photo_size"`
	PhotoFormKey   string `yaml:"photo_form_key"`
	MessageFormKey string `yaml:"message_form_key"`
	FAQCommandName string `yaml:"faq_command_name"`
}

type tasksConfig struct {
	PhotoPathGoals *string `yaml:"photo_path_goals"`
	PhotoPathTasks *string `yaml:"photo_path_tasks"`
	PhotoPathTodos *string `yaml:"photo_path_todos"`
	BarCount       int     `yaml:"bar_count"`
}

func weeks(n int) time.Duration {
	return time.Duration(n) * 24 * 7 * time.Hour
}

func toMegabytes(n int64) int64 {
	return n << 32
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
		Tasks: TasksConfig{
			PhotoPathGoals: cfg.Tasks.PhotoPathGoals,
			PhotoPathTasks: cfg.Tasks.PhotoPathTasks,
			PhotoPathTodos: cfg.Tasks.PhotoPathTodos,
			BarCount:       cfg.Tasks.BarCount,
		},
		Admin: AdminConfig{
			Port:           cfg.Admin.Port,
			MaxPhotoSize:   toMegabytes(cfg.Admin.MaxPhotoSize),
			PhotoFormKey:   cfg.Admin.PhotoFormKey,
			MessageFormKey: cfg.Admin.MessageFormKey,
			FAQCommandName: cfg.Admin.FAQCommandName,
		},
		CalendarUrl: cfg.Calendar.Url,
	}

	return nil
}
