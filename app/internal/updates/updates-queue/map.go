package updates_queue

import (
	"sync"
	"telegrambot_new_emploee/internal/models"
	"telegrambot_new_emploee/internal/updates"
)

type updatesMap struct {
	lock  sync.Mutex
	store map[int64]updates.Queue
}

func NewMap() updates.Map {
	return &updatesMap{
		lock:  sync.Mutex{},
		store: make(map[int64]updates.Queue),
	}
}

func (m *updatesMap) GetOrCreate(key int64) (updates.Queue, bool) {
	m.lock.Lock()
	defer m.lock.Unlock()
	queue, ok := m.store[key]
	if !ok {
		queue := NewQueue()
		m.store[key] = queue
		return queue, false
	}

	return queue, true
}

func (m *updatesMap) GetUpdate(key int64, queue updates.Queue) *models.Update {
	m.lock.Lock()
	queue.Lock()
	defer m.lock.Unlock()

	if queue.Size() == 0 {
		delete(m.store, key)
		queue.Unlock()
		return nil
	}
	queue.Unlock()

	return queue.GetUpdate()
}
