package repository

import (
	"context"
	"telegrambot_new_emploee/internal/models"
)

type UpdateFaq struct {
	SectionName string
	Question    string
	Answer      string
}

type AdminRepository interface {
	GetFAQSections(ctx context.Context) ([]string, error)
	UpdateFAQ(ctx context.Context, faq *UpdateFaq) error
	GetNotifications(ctx context.Context) ([]models.Notification, error)
}
