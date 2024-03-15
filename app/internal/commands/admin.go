package commands

import (
	"context"
	container "telegrambot_new_emploee/internal/di-container"
	"telegrambot_new_emploee/internal/models"
	"telegrambot_new_emploee/internal/repository"
	"telegrambot_new_emploee/internal/services"
	"telegrambot_new_emploee/internal/views"
)

type getUnansweredQuestionsCmd struct {
	service *services.QuestionService
}

func NewGetUnansweredQuestionsCmd() Cmd {
	return &getUnansweredQuestionsCmd{
		service: services.NewQuestionService(),
	}
}

func (c *getUnansweredQuestionsCmd) Execute(ctx context.Context, job *Job) error {
	questions, err := c.service.GetQuestions(ctx)
	if err != nil {
		return err
	}
	return container.Container.Bot().SendMessage(ctx, views.UnansweredQuestionsView(questions, job.GetChatId()))
}

type answerQuestionCmd struct {
	service *services.QuestionService
}

func NewAnswerQuestionCmd() Cmd {
	return &answerQuestionCmd{
		service: services.NewQuestionService(),
	}
}

func (c *answerQuestionCmd) Execute(ctx context.Context, job *Job) error {
	questionId, err := getNumberWithMessage(ctx, job, -1, "Введи id вопроса, на который ты хочешь ответить")
	if err != nil {
		return err
	}
	answer, err := getStringWithMessage(ctx, job, "Введи ответ на сообщение")

	err = c.service.AnswerQuestion(ctx, &services.AnswerQuestionRequest{
		QuestionId:  int64(questionId),
		Answer:      answer,
		ResponderId: job.User.Id,
	})
	if err != nil {
		return err
	}

	return container.Container.Bot().SendMessage(ctx, models.NewMessage("Успех!", job.GetChatId()))
}

type addQuestionToFAQCmd struct {
}

func NewAddQuestionToFAQCmd() Cmd {
	return &addQuestionToFAQCmd{}
}

func (c *addQuestionToFAQCmd) Execute(ctx context.Context, job *Job) error {
	sections, err := container.Container.FAQRepo().GetFAQSections(ctx)
	if err != nil {
		return err
	}
	err = container.Container.Bot().SendButtons(ctx, views.ChooseSectionView(sections, job.GetChatId()))
	if err != nil {
		return err
	}

	section, err := getString(ctx, job)
	if err != nil {
		return err
	}
	question, err := getStringWithMessage(ctx, job, "Введи вопрос")
	if err != nil {
		return err
	}
	answer, err := getStringWithMessage(ctx, job, "Введи ответ")
	if err != nil {
		return err
	}

	err = container.Container.FAQRepo().UpdateFAQ(ctx, &repository.UpdateFaq{
		SectionName: section,
		Question:    question,
		Answer:      answer,
	})
	if err != nil {
		return err
	}

	return container.Container.Bot().SendMessage(ctx, models.NewMessage("Успех!", job.GetChatId()))
}
