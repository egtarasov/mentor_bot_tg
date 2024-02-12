package models

type Command struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	Action   string `db:"action"`
	ParentId int64  `db:"parent_id"`
}

type Material struct {
	Id        int64  `db:"id"`
	Message   string `db:"message"`
	CommandId int64  `db:"command_id"`
}

type Task struct {
	Id          int64  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	StoryPoints int64  `db:"story_points"`
	Completed   bool   `db:"completed"`
	EmployeeId  int64  `db:"employee_id"`
}

type Todo struct {
	Id         int64  `db:"id"`
	Label      string `db:"label"`
	Priority   int    `db:"priority"`
	EmployeeId int64  `db:"employee_id"`
	Completed  bool   `db:"completed"`
}

type User struct {
	Id           int64  `db:"id"`
	Name         string `db:"name"`
	Surname      string `db:"surname"`
	TelegramId   int64  `db:"telegram_id"`
	OccupationId int64  `db:"occupation_id"`
}

type Goal struct {
	Id          int64  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	EmployeeId  int64  `db:"employee_id"`
	Track       string `db:"track"`
}
