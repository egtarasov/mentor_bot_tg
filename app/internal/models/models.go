package models

type Update struct {
	UpdateUserId int64
	User         *User
	ChatId       int64
	Message      string
}

type User struct {
	Id           int64
	TelegramId   int64
	Name         string
	OccupationId int64
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
	Completed   bool
	EmployeeId  int64
}

type Message struct {
	Message string
	ChatId  int64
}

type Button string

type Buttons struct {
	ChatId  int64
	Buttons []Button
	Message string
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

func NewMessage(msg string, chatID int64) *Message {
	return &Message{
		Message: msg,
		ChatId:  chatID,
	}
}
