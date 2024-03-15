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
	ConnStr       string
	TgToken       string
	Feedback      FeedBackConfig
	Tasks         TasksConfig
	Admin         AdminConfig
	Notifications NotificationsConfig
	CalendarUrl   *string
}

type NotificationsConfig struct {
	PhotoPath *string

	TrainingRepeat    time.Duration
	TrainingDayOfWeek int
	TrainingHour      int

	HrMeetupRepeat    time.Duration
	HrMeetupDayOfWeek int
	HrMeetupHour      int

	MentorMeetupRepeat    time.Duration
	MentorMeetupDayOfWeek int
	MentorMeetupHour      int
}

type AdminConfig struct {
	Port           string
	MaxPhotoSize   int64
	PhotoFormKey   string
	MessageFormKey string
	FAQCommandName string
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
	BarCount       int
}

var Cfg *Config

const (
	tgToken = "TELEGRAM_TOKEN"
)

// Yaml configuration.
type yamlConfig struct {
	Feedback      feedbackConfig     `yaml:"feedback"`
	Tasks         tasksConfig        `yaml:"tasks"`
	Admin         adminConfig        `yaml:"admin"`
	Notifications notificationConfig `yaml:"notifications"`
	Calendar      calendarConfig     `yaml:"calendar"`
}

type calendarConfig struct {
	Url *string `yaml:"url"`
}

type notificationConfig struct {
	PhotoPath *string `yaml:"photo_path"`

	TrainingRepeat    int    `yaml:"training_repeat"`
	TrainingDayOfWeek string `yaml:"training_day_of_week"`
	TrainingHour      int    `yaml:"training_hour"`

	HrMeetupRepeat    int    `yaml:"hr_meetup_repeat"`
	HrMeetupDayOfWeek string `yaml:"hr_meetup_day_of_week"`
	HrMeetupHour      int    `yaml:"hr_meetup_hour"`

	MentorMeetupRepeat    int    `yaml:"mentor_meetup_repeat"`
	MentorMeetupDayOfWeek string `yaml:"mentor_meetup_day_of_week"`
	MentorMeetupHour      int    `yaml:"mentor_meetup_hour"`
}

type adminConfig struct {
	Port           string `yaml:"port"`
	MaxPhotoSize   int64  `yaml:"max_photo_size"`
	PhotoFormKey   string `yaml:"photo_form_key"`
	MessageFormKey string `yaml:"message_form_key"`
	FAQCommandName string `yaml:"faq_command_name"`
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
		Feedback: FeedBackConfig{
			Form:      cfg.Feedback.Form,
			Duration:  weeks(cfg.Feedback.Duration),
			PhotoPath: cfg.Feedback.PhotoPath,
		},
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
		Notifications: NotificationsConfig{
			PhotoPath:             cfg.Notifications.PhotoPath,
			TrainingRepeat:        weeks(cfg.Notifications.TrainingRepeat),
			TrainingDayOfWeek:     parseDayOfWeek(cfg.Notifications.TrainingDayOfWeek),
			TrainingHour:          cfg.Notifications.TrainingHour,
			HrMeetupRepeat:        weeks(cfg.Notifications.HrMeetupRepeat),
			HrMeetupDayOfWeek:     parseDayOfWeek(cfg.Notifications.HrMeetupDayOfWeek),
			HrMeetupHour:          cfg.Notifications.HrMeetupHour,
			MentorMeetupRepeat:    weeks(cfg.Notifications.MentorMeetupRepeat),
			MentorMeetupDayOfWeek: parseDayOfWeek(cfg.Notifications.MentorMeetupDayOfWeek),
			MentorMeetupHour:      cfg.Notifications.MentorMeetupHour,
		},
		CalendarUrl: cfg.Calendar.Url,
	}

	return nil
}
