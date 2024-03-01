package convert

import (
	"telegrambot_new_emploee/internal/models"
)

var (
	returnButton = models.Command{
		Name: "В меню",
	}
)

type ToButtonsCfg struct {
	ChatId       int64
	Message      string
	ButtonsInRow int
	ReturnButton bool
}

func DefaultToButtonsCfg(message string, chatId int64) *ToButtonsCfg {
	return &ToButtonsCfg{
		ChatId:       chatId,
		Message:      message,
		ButtonsInRow: 2,
		ReturnButton: true,
	}
}

func ToButtons(commands []models.Command, cfg *ToButtonsCfg) *models.Buttons {
	size := len(commands)
	if cfg.ReturnButton {
		size++
		commands = append(commands, returnButton)
	}
	res := &models.Buttons{
		Message: models.NewMessage(cfg.Message, cfg.ChatId),
		Buttons: make([][]models.Button, 0, size),
	}

	i := 0
	for i < len(commands) {
		row := make([]models.Button, 0, cfg.ButtonsInRow)
		for j := 0; j < cfg.ButtonsInRow && i < len(commands); j++ {
			row = append(row, models.Button(commands[i].Name))
			i++
		}
		res.Buttons = append(res.Buttons, row)
	}

	return res
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
