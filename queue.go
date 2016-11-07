package linenotcat

import (
	"sync"
)

type queue struct {
	queue []string
	lock  sync.RWMutex
}

func newQueue() *queue {
	return &queue{
		queue: []string{},
		lock:  sync.RWMutex{},
	}
}

func (q *queue) add(line string) {
	q.lock.Lock()
	q.queue = append(q.queue, line)
	q.lock.Unlock()
}

func (q *queue) isEmpty() bool {
	return (len(q.queue) < 1)
}

func (q *queue) flush() []string {
	q.lock.Lock()
	items := q.queue
	q.queue = []string{}
	q.lock.Unlock()
	return items
}
