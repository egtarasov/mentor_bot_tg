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

func IntToAction(i int) Action {
	values := map[int]Action{
		1: GetDataCmd,
		2: GetSubsectionsCmd,
		3: ComplexCmd,
	}
	return values[i]
}

type Command struct {
	Id       int64
	Name     string
	ActionId int
	ParentId *int64
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
	PhotoPath  *string
	PhotoBytes []byte
	Message    string
	ChatId     int64
}

type Button string

type Buttons struct {
	Message *Message
	Buttons [][]Button
}

type Track string

const (
	DefaultTrack Track = "default"
)

type Goal struct {
	Id          int64
	Name        string
	Description string
	EmployeeId  int64
	Track       Track
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

type Occupation struct {
	Id       int64
	Name     string
	Material string
}

type CommandWithMaterial struct {
	Id       int64
	Name     string
	Message  string
	ActionId int64
}

type Meeting struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	StartTime   time.Duration `json:"start_time"`
}

func NewMessage(msg string, chatID int64) *Message {
	return &Message{
		PhotoPath:  nil,
		PhotoBytes: nil,
		Message:    msg,
		ChatId:     chatID,
	}
}

func NewMessageWithPhotoPath(msg string, chatID int64, photoPath *string) *Message {
	return &Message{
		PhotoPath:  photoPath,
		PhotoBytes: nil,
		Message:    msg,
		ChatId:     chatID,
	}
}

func NewMessageWithPhotoBytes(msg string, chatID int64, photoBytes []byte) *Message {
	return &Message{
		PhotoPath:  nil,
		PhotoBytes: photoBytes,
		Message:    msg,
		ChatId:     chatID,
	}
}

func NewQuestion(text string, userId int64) *Question {
	return &Question{
		UserId: userId,
		Text:   text,
	}
}
