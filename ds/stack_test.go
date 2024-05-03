package ds_test

import (
	"testing"

	"github.com/bcdxn/dsa-go/ds"
	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	t.Run("enqueue into an empty stack", func(t *testing.T) {
		s := ds.NewStack[int]()

		assert.Equalf(t, 0, s.Height(), "Stack should start out empty")
	})
}

func TestStackPush(t *testing.T) {
	t.Run("push into an empty stack", func(t *testing.T) {
		s := ds.NewStack[int]()

		assert.Equalf(t, 0, s.Height(), "Stack should start out empty")
		s.Push(10)
		assert.Equal(t, 1, s.Height(), "Stack height should have grown by 1")
	})

	t.Run("enqueue into a stack with existing elements", func(t *testing.T) {
		s := ds.NewStack[int]()

		assert.Equalf(t, 0, s.Height(), "Stack should start out empty")
		s.Push(10)
		s.Push(10)
		s.Push(10)
		s.Push(10)
		assert.Equal(t, 4, s.Height(), "Pushing 4 items should grow the height by 4")
	})
}

func TestStackPop(t *testing.T) {
	t.Run("deuque from an empty stack", func(t *testing.T) {
		s := ds.NewStack[int]()

		assert.Equalf(t, 0, s.Height(), "Stack should start out empty")
		_, err := s.Pop()
		assert.NotNil(t, err, "Popping from an empty stack should return an error")
	})

	t.Run("dequeue from a stack with existing elements", func(t *testing.T) {
		s := ds.NewStack[int]()

		assert.Equalf(t, 0, s.Height(), "Stack should start out empty")
		s.Push(10)
		s.Push(5)
		s.Push(8)
		s.Push(100)
		assert.Equal(t, 4, s.Height(), "Pushing 4 items should grow the height by 4")
		elem, err := s.Pop()
		assert.Nilf(t, err, "Popping from a stack with items should not result in an error")
		assert.Equal(t, 100, elem, "Stack should be LIFO")
		assert.Equal(t, 3, s.Height(), "Popping 1 item should decrease the height by 1")
		elem, err = s.Pop()
		assert.Nilf(t, err, "Popping from a stack with items should not result in an error")
		assert.Equal(t, 8, elem, "Stack should be LIFO")
		assert.Equal(t, 2, s.Height(), "Popping 1 item should decrease the height by 1")
		elem, err = s.Pop()
		assert.Nilf(t, err, "Popping from a stack with items should not result in an error")
		assert.Equal(t, 5, elem, "Stack should be LIFO")
		assert.Equal(t, 1, s.Height(), "Popping 1 item should decrease the height by 1")
		elem, err = s.Pop()
		assert.Nilf(t, err, "Popping from a stack with items should not result in an error")
		assert.Equal(t, 10, elem, "Stack should be LIFO")
		assert.Equal(t, 0, s.Height(), "Popping 1 item should decrease the height by 1")
	})
}
