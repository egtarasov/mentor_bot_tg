package repository

import (
	"context"
	"fmt"
	"telegrambot_new_emploee/internal/models"
)

var (
	ErrQuestionNotExist = fmt.Errorf("question does not exist or already anwered")
)

type QuestionRepo interface {
	CreateQuestion(ctx context.Context, question *models.Question) (int64, error)
	GetUnansweredQuestions(ctx context.Context) ([]models.Question, error)
	AnswerQuestion(ctx context.Context, question *models.Question) error
	GetQuestionById(ctx context.Context, questionId int64) (*models.Question, error)
}
