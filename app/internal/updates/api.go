package updates

import "telegrambot_new_emploee/internal/models"

type Queue interface {
	AddUpdate(update *models.Update)
	GetUpdate() *models.Update
	WaitForUpdate() *models.Update
}
