package updates

import (
	"sync"
	"telegrambot_new_emploee/internal/models"
)

type Queue interface {
	sync.Locker
	// Size returns the size of the queue. This method is not concurrent safe, so make sure to use Lock before calling.
	Size() int
	AddUpdate(update *models.Update)
	GetUpdate() *models.Update
	WaitForUpdate() *models.Update
}

type Map interface {
	GetOrCreate(key int64) (Queue, bool)
	GetUpdate(key int64, queue Queue) *models.Update
}
