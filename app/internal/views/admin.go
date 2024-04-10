package views

import (
	"fmt"
	"strings"
	"telegrambot_new_emploee/internal/convert"
	"telegrambot_new_emploee/internal/models"
)

func UnansweredQuestionsView(questions []models.Question, chatId int64) *models.Message {
	builder := strings.Builder{}
	builder.WriteString("Вопросы сотрудников:\n\n")
	for _, question := range questions {
		builder.WriteString(questionView(question))
		builder.WriteString("\n")
	}
	return models.NewMessage(builder.String(), chatId)
}

func questionView(question models.Question) string {
	return fmt.Sprintf(
		"*Id*: %d\n"+
			"*Вопрос*: %s\n", question.Id, question.Text)
}

func ChooseSectionView(sections []string, chatId int64) *models.Buttons {
	return convert.ToButtons(sections, convert.DefaultToButtonsCfg("Выбери секцию FAQ в меню ниже или введи 'отмена'", chatId))
}
