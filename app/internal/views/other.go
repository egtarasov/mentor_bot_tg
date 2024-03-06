package views

import (
	"fmt"
	"strings"
	"telegrambot_new_emploee/internal/models"
	"time"
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

func CalendarMessage(meetings []models.Meeting, chatId int64) *models.Message {
	msg := strings.Builder{}
	msg.WriteString("Встречи на сегодня:\n\n")
	for _, meeting := range meetings {
		meetingView(&msg, &meeting)
		msg.WriteString("\n\n")
	}
	return models.NewMessage(msg.String(), chatId)
}

func meetingView(msg *strings.Builder, meeting *models.Meeting) {
	msg.WriteString(fmt.Sprintf("*%s*\nОписание: %s\nНачало: %s",
		meeting.Name, meeting.Description, fmtDuration(meeting.StartTime)))
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%02d:%02d", h, m)
}
