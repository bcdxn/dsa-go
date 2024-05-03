package ds_test

import (
	"testing"

	"github.com/bcdxn/dsa-go/ds"
	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	t.Run("enqueue into an empty queue", func(t *testing.T) {
		q := ds.NewQueue[int]()

		assert.Equalf(t, 0, q.Depth(), "Queue should start out empty")
	})
}

func TestQueueEnqueue(t *testing.T) {
	t.Run("enqueue into an empty queue", func(t *testing.T) {
		q := ds.NewQueue[int]()

		assert.Equalf(t, 0, q.Depth(), "Queue should start out empty")
		q.Enqueue(10)
		assert.Equal(t, 1, q.Depth(), "Queue depth should have grown by 1")
	})

	t.Run("enqueue into a queue with existing elements", func(t *testing.T) {
		q := ds.NewQueue[int]()

		assert.Equalf(t, 0, q.Depth(), "Queue should start out empty")
		q.Enqueue(10)
		q.Enqueue(10)
		q.Enqueue(10)
		q.Enqueue(10)
		assert.Equal(t, 4, q.Depth(), "Enqueuing 4 items should grow the depth by 4")
	})
}

func TestQueueDequeue(t *testing.T) {
	t.Run("deuque from an empty queue", func(t *testing.T) {
		q := ds.NewQueue[int]()

		assert.Equalf(t, 0, q.Depth(), "Queue should start out empty")
		_, err := q.Dequeue()
		assert.NotNil(t, err, "Dequeuing from an empty queue should return an error")
	})

	t.Run("dequeue from a queue with existing elements", func(t *testing.T) {
		q := ds.NewQueue[int]()

		assert.Equalf(t, 0, q.Depth(), "Queue should start out empty")
		q.Enqueue(10)
		q.Enqueue(5)
		q.Enqueue(8)
		q.Enqueue(100)
		assert.Equal(t, 4, q.Depth(), "Enqueuing 4 items should grow the depth by 4")
		elem, err := q.Dequeue()
		assert.Nilf(t, err, "Dequeuing from a queue with items should not result in an error")
		assert.Equal(t, 10, elem, "Queue should be FIFO")
		assert.Equal(t, 3, q.Depth(), "Dequeuing 1 item should decrease the depth by 1")
		elem, err = q.Dequeue()
		assert.Nilf(t, err, "Dequeuing from a queue with items should not result in an error")
		assert.Equal(t, 5, elem, "Queue should be FIFO")
		assert.Equal(t, 2, q.Depth(), "Dequeuing 1 item should decrease the depth by 1")
		elem, err = q.Dequeue()
		assert.Nilf(t, err, "Dequeuing from a queue with items should not result in an error")
		assert.Equal(t, 8, elem, "Queue should be FIFO")
		assert.Equal(t, 1, q.Depth(), "Dequeuing 1 item should decrease the depth by 1")
		elem, err = q.Dequeue()
		assert.Nilf(t, err, "Dequeuing from a queue with items should not result in an error")
		assert.Equal(t, 100, elem, "Queue should be FIFO")
		assert.Equal(t, 0, q.Depth(), "Dequeuing 1 item should decrease the depth by 1")
	})
}
