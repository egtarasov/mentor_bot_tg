package views

import (
	"fmt"
	"telegrambot_new_emploee/internal/models"
)

const (
	AskQuestionRequest = "Напиши вопрос, который ты хотел бы задать, или 'Отмена', чтобы отменить"
)

func AskQuestionSuccess(id int64, text string, chatId int64) *models.Message {
	msg := fmt.Sprintf(
		"Ваш вопрос получил номер #%d\n"+
			"```\n"+
			"%s\n"+
			"```\n"+
			"Впорос был передан HR-отедлу, дожидайтесь ответа", id, text)
	return models.NewMessage(msg, chatId)
}

func QuestionResponse(question *models.Question, chatId int64) *models.Message {
	msg := fmt.Sprintf(
		"Ответ на ваш вопрос номер #%d\n"+
			"```\n"+
			"%s\n"+
			"```", question.Id, *question.Answer)
	return models.NewMessage(msg, chatId)
}
