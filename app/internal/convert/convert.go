package convert

import "telegrambot_new_emploee/internal/models"

func ToButtons(commands []models.Command, chatId int64, message string) models.Buttons {
	res := models.Buttons{
		ChatId:  chatId,
		Buttons: make([]models.Button, 0, len(commands)),
		Message: message,
	}

	for _, command := range commands {
		res.Buttons = append(res.Buttons, models.Button(command.Name))
	}

	return res
}
