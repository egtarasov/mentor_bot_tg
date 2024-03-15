package services

import (
	"context"
	"fmt"
	"log"
	container "telegrambot_new_emploee/internal/di-container"
	"telegrambot_new_emploee/internal/models"
	"telegrambot_new_emploee/internal/views"
)

var (
	ErrQuestionAnswered = fmt.Errorf("qustion is already answered")
)

type QuestionService struct {
}

func NewQuestionService() *QuestionService {
	return &QuestionService{}
}

func (s *QuestionService) GetQuestions(ctx context.Context) ([]models.Question, error) {
	questions, err := container.Container.QuestionRepo().GetUnansweredQuestions(ctx)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (s *QuestionService) GetQuestion(ctx context.Context, questionId int64) (*models.Question, error) {
	question, err := container.Container.QuestionRepo().GetQuestionById(ctx, questionId)
	if err != nil {
		return nil, err
	}
	return question, nil
}

func (s *QuestionService) AnswerQuestion(ctx context.Context, req *AnswerQuestionRequest) error {
	// Get the question and check that there is no answer.
	question, err := container.Container.QuestionRepo().GetQuestionById(ctx, req.QuestionId)
	if err != nil {
		return err
	}
	if question.AnsweredAt != nil {
		return ErrQuestionAnswered
	}
	question.Answer = &req.Answer
	question.AnsweredBy = &req.ResponderId

	// Get the user, who asked the question.
	user, err := container.Container.UserRepo().GetUserById(ctx, question.UserId)
	if err != nil {
		return err
	}

	// Send the answer to user.
	err = container.Container.Bot().SendMessage(ctx, views.QuestionResponse(question, user.TelegramId))
	if err != nil {
		return err
	}

	// Update info about question status in db.
	err = container.Container.QuestionRepo().AnswerQuestion(ctx, question)
	return err
}

type CommandsService struct {
}

func NewCommandsService() *CommandsService {
	return &CommandsService{}
}

func (s *CommandsService) GetCommands(ctx context.Context) ([]models.CommandWithMaterial, error) {
	commands, err := container.Container.CmdRepo().GetCommandsWithMaterials(ctx)
	if err != nil {
		return nil, err
	}
	return commands, nil
}

func (s *CommandsService) UpdateCommand(ctx context.Context, req *UpdateCommandRequest) error {
	return container.Container.CmdRepo().UpdateCommand(ctx, req.Name, &models.Material{
		Message:   req.Message,
		CommandId: req.CommandId,
	})
}

func (s *CommandsService) AddCommand(ctx context.Context, req *AddCommandRequest) error {
	err := container.Container.CmdRepo().
		AddCommand(
			ctx,
			&models.Command{
				Name:     req.Name,
				ActionId: req.ActionId,
				ParentId: req.ParentId,
			},
			req.Message)
	return err
}

func SendMessage(ctx context.Context, req *SendMessageRequest) error {
	users, err := container.Container.UserRepo().GetUsersOnAdaptation(ctx)
	if err != nil {
		return err
	}

	message := models.NewMessageWithPhotoBytes(req.Message, -1, req.Photo)
	for _, user := range users {
		message.ChatId = user.TelegramId
		err := container.Container.Bot().SendMessage(ctx, message)
		if err != nil {
			log.Println(err)
			continue
		}
	}

	return nil
}
