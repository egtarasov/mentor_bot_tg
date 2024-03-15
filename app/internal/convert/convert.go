package convert

import (
	"telegrambot_new_emploee/internal/models"
	"time"
)

var (
	returnButton = "В меню"
)

type ToButtonsCfg struct {
	ChatId       int64
	Message      string
	ButtonsInRow int
	ReturnButton bool
}

func TimeToDuration(t time.Time) time.Duration {
	hours := time.Duration(t.Hour())
	minutes := time.Duration(t.Minute())
	seconds := time.Duration(t.Second())
	return time.Hour*hours + time.Minute*minutes + time.Second*seconds
}

func DefaultToButtonsCfg(message string, chatId int64) *ToButtonsCfg {
	return &ToButtonsCfg{
		ChatId:       chatId,
		Message:      message,
		ButtonsInRow: 2,
		ReturnButton: true,
	}
}

func ToButtons(buttonLabels []string, cfg *ToButtonsCfg) *models.Buttons {
	size := len(buttonLabels)
	if cfg.ReturnButton {
		size++
		buttonLabels = append(buttonLabels, returnButton)
	}
	res := &models.Buttons{
		Message: models.NewMessage(cfg.Message, cfg.ChatId),
		Buttons: make([][]models.Button, 0, size),
	}

	i := 0
	for i < len(buttonLabels) {
		row := make([]models.Button, 0, cfg.ButtonsInRow)
		for j := 0; j < cfg.ButtonsInRow && i < len(buttonLabels); j++ {
			row = append(row, models.Button(buttonLabels[i]))
			i++
		}
		res.Buttons = append(res.Buttons, row)
	}

	return res
}

func CommandsToButtons(commands []models.Command, cfg *ToButtonsCfg) *models.Buttons {
	buttonLabels := make([]string, 0, len(commands))
	for _, command := range commands {
		buttonLabels = append(buttonLabels, command.Name)
	}

	return ToButtons(buttonLabels, cfg)
}

func UncompletedTodo(todos []models.Todo) []models.Todo {
	uncompletedTodos := make([]models.Todo, 0, len(todos))
	for _, todo := range todos {
		if todo.Completed {
			continue
		}
		uncompletedTodos = append(uncompletedTodos, todo)
	}

	return uncompletedTodos
}
