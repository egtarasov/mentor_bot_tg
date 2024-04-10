package repository

import (
	"context"
	"fmt"
	"telegrambot_new_emploee/internal/models"
)

type UpdateFaq struct {
	SectionName string
	Question    string
	Answer      string
}

var ErrNoSection = fmt.Errorf("no section faq")

type AdminRepository interface {
	GetFAQSections(ctx context.Context) ([]string, error)
	UpdateFAQ(ctx context.Context, faq *UpdateFaq) error
	GetNotifications(ctx context.Context) ([]models.Notification, error)
}
