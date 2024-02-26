package views

import "fmt"

const (
	AskQuestionRequest = "Напиши вопрос, который ты хотел бы задать, или 'Отмена', чтобы отменить"
)

func AskQuestionSuccess(id int64, text string) string {
	return fmt.Sprintf(
		"Ваш вопрос получил номер #%d\n"+
			"```\n"+
			"%s\n"+
			"```\n"+
			"Впорос был передан HR-отедлу, дожидайтесь ответа", id, text)
}
