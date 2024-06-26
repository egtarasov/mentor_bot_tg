package commands

import (
	"context"
	"errors"
	"log"
	container "telegrambot_new_emploee/internal/di-container"
	"telegrambot_new_emploee/internal/models"
	"telegrambot_new_emploee/internal/repository"
	"telegrambot_new_emploee/internal/services"
	"telegrambot_new_emploee/internal/views"
)

type getUnansweredQuestionsCmd struct {
	service *services.QuestionService
}

func NewGetUnansweredQuestionsCmd() Command {
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

func NewAnswerQuestionCmd() Command {
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

func NewAddQuestionToFAQCmd() Command {
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
	question, err := getStringWithMessage(ctx, job, "Введи вопрос или 'отмена'")
	if err != nil {
		return err
	}
	answer, err := getStringWithMessage(ctx, job, "Введи ответ или 'отмена'")
	if err != nil {
		return err
	}

	err = container.Container.FAQRepo().UpdateFAQ(ctx, &repository.UpdateFaq{
		SectionName: section,
		Question:    question,
		Answer:      answer,
	})
	if errors.Is(err, repository.ErrNoSection) {
		return container.Container.Bot().SendMessage(ctx, models.NewMessage("Ты ввел не секцию FAQ", job.GetChatId()))
	}
	if err != nil {
		return err
	}

	return container.Container.Bot().SendMessage(ctx, models.NewMessage("Успех!", job.GetChatId()))
}

type sendMessageStringCmd struct {
}

func NewSendMessageCmd() Command {
	return &sendMessageStringCmd{}
}

func (c *sendMessageStringCmd) Execute(ctx context.Context, job *Job) error {
	err := container.Container.Bot().SendMessage(ctx, models.NewMessage("Введи сообщение для рассылки", job.GetChatId()))
	if err != nil {
		return err
	}

	update := job.Queue.WaitForUpdate()
	message := models.NewMessageWithPhotoIds(update.Message, -1, update.PhotoIds)

	users, err := container.Container.UserRepo().GetUsersOnAdaptation(ctx)
	if err != nil {
		return err
	}

	for _, user := range users {
		message.ChatId = user.TelegramId
		err := container.Container.Bot().SendMessage(ctx, message)
		if err != nil {
			log.Println(err)
		}
	}

	return container.Container.Bot().SendMessage(ctx, models.NewMessage("Успех!", job.GetChatId()))
}
