package models

import (
	"time"
)

type Update struct {
	UpdateUserId int64
	User         *User
	ChatId       int64
	Message      string
}

type User struct {
	Id             int64
	TelegramId     int64
	Name           string
	OccupationId   int64
	StartWork      time.Time
	AdaptationEnds time.Time
}

type Action string

const (
	GetDataCmd        Action = "get data"
	GetSubsectionsCmd Action = "show subsections"
	ComplexCmd        Action = "complex"
)

type Command struct {
	Id       int64
	Name     string
	Action   Action
	ParentId int64
}

type Todo struct {
	Id         int64
	Label      string
	Priority   int
	EmployeeId int64
	Completed  bool
}

type Material struct {
	Id        int64
	Message   string
	CommandId int64
}

type Task struct {
	Id          int64
	Name        string
	Description string
	StoryPoints int64
	EmployeeId  int64
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Message struct {
	PhotoPath *string
	Message   string
	ChatId    int64
}

type Button string

type Buttons struct {
	Message *Message
	Buttons [][]Button
}

type GoalTrack string

const (
	DefaultTrack GoalTrack = "default"
)

type Goal struct {
	Id          int64
	Name        string
	Description string
	EmployeeId  int64
	Track       GoalTrack
}

type Question struct {
	Id         int64
	UserId     int64
	Text       string
	CreatedAt  time.Time
	AnsweredAt *time.Time
	AnsweredBy *int64
	Answer     *string
}

func NewMessage(msg string, chatID int64) *Message {
	return &Message{
		Message: msg,
		ChatId:  chatID,
	}
}

func NewMessageWithPhoto(msg string, chatID int64, photoPath *string) *Message {
	return &Message{
		PhotoPath: photoPath,
		Message:   msg,
		ChatId:    chatID,
	}
}

func NewQuestion(text string, userId int64) *Question {
	return &Question{
		UserId: userId,
		Text:   text,
	}
}
