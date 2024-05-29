package ds

import (
	"container/list"
	"errors"
)

// A Queue implemented using the go standard library container.List as the underlying data structure.
type QueueN struct {
	l *list.List
}

// NewQueue creates and returns an empty queue.
func NewQueueN() *QueueN {
	return &QueueN{
		l: list.New(),
	}
}

// Depth returns the number of elements in the queue.
func (q QueueN) Depth() int {
	return q.l.Len()
}

// Enqueue adds an element to the end of the queue
func (q *QueueN) Enqueue(elem any) {
	q.l.PushBack(elem)
}

// Dequeue removes an element from the front of the queue (FIFO)
func (q *QueueN) Dequeue() (any, error) {
	if q.l.Len() < 1 {
		return nil, errors.New("cannot Dequeue from an empty queue")
	}
	return q.l.Remove(q.l.Front()), nil
}
