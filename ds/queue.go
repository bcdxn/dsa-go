package ds

import "golang.org/x/exp/constraints"

// Queue is a simple data structure that offers FIFO access to its elements.
type Queue[T constraints.Ordered] struct {
	l List[T]
}

// NewQueue creates and returns an empty queue.
func NewQueue[T constraints.Ordered]() *Queue[T] {
	return &Queue[T]{
		l: *NewList[T](),
	}
}

// Depth returns the number of elements in the queue.
func (q Queue[T]) Depth() int {
	return q.l.Len()
}

// Enqueue adds an element to the end of the queue
func (q *Queue[T]) Enqueue(elem T) {
	q.l.AddTail(elem)
}

// Dequeue removes an element from the front of the queue (FIFO)
func (q *Queue[T]) Dequeue() (T, error) {
	return q.l.RemoveHead()
}
