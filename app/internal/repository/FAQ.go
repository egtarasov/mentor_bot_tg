package repository

import "context"

type UpdateFaq struct {
	SectionName string
	Question    string
	Answer      string
}

type FAQRepository interface {
	GetFAQSections(ctx context.Context) ([]string, error)
	UpdateFAQ(ctx context.Context, faq *UpdateFaq) error
}
