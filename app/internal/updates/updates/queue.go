package updates

import (
	"fmt"
	"sync"
	"telegrambot_new_emploee/internal/models"
	"telegrambot_new_emploee/internal/updates"
)

type updatesQueue struct {
	lock sync.Mutex

	user *models.User

	waiting bool
	sync    chan *models.Update
	updates []*models.Update
}

func NewQueue(user *models.User) updates.Queue {
	return &updatesQueue{
		lock:    sync.Mutex{},
		user:    user,
		waiting: false,
		sync:    make(chan *models.Update),
		updates: make([]*models.Update, 0, 1),
	}
}

// AddUpdate adds an update to the queue.
func (q *updatesQueue) AddUpdate(update *models.Update) {
	q.lock.Lock()
	defer q.lock.Unlock()

	// If another goroutine is waiting for an update, synchronizes with this goroutine.
	if q.waiting == true {
		fmt.Println("Waiting for an update")
		q.waiting = false
		q.sync <- update
		return
	}

	q.addUpdate(update)
}

// GetUpdate gets an update from the queue. If there is no updates, returns nil.
func (q *updatesQueue) GetUpdate() *models.Update {
	q.lock.Lock()
	defer q.lock.Unlock()
	update := q.getUpdate()
	return update
}

func (q *updatesQueue) WaitForUpdate() *models.Update {
	// Determine whether there is an update in the queue.
	q.lock.Lock()
	update := q.getUpdate()
	// If there is an update, just return it.
	if update != nil {
		q.lock.Unlock()
		return update
	}
	// If there is no update, mark the waiting flag and wait for an update.
	q.waiting = true
	q.lock.Unlock()
	update = <-q.sync

	return update
}

func (q *updatesQueue) addUpdate(update *models.Update) {
	q.updates = append(q.updates, update)
}

func (q *updatesQueue) getUpdate() *models.Update {
	if len(q.updates) == 0 {
		return nil
	}
	update := q.updates[0]
	q.updates = q.updates[1:]

	return update
}
