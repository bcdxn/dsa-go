package ds_test

import (
	"testing"

	"github.com/bcdxn/dsa-go/ds"
	"github.com/stretchr/testify/assert"
)

func TestNewMaxHeap(t *testing.T) {
	t.Run("Should create a new heap of size 0", func(t *testing.T) {
		h := ds.NewMaxHeap[int]()

		assert.Equal(t, 0, h.Size())
		_, err := h.Peek()
		assert.NotNil(t, err)
	})
}

func TestMaxHeapPush(t *testing.T) {
	t.Run("Push into empty heap", func(t *testing.T) {
		h := ds.NewMaxHeap[int]()
		assert.Equal(t, 0, h.Size())
		h.Push(10)
		assert.Equal(t, 1, h.Size())
	})

	t.Run("Push into non empty heap", func(t *testing.T) {
		h := ds.NewMaxHeap[int]()
		h.Push(10)
		h.Push(5)
		h.Push(2)
		h.Push(100)
		assert.Equal(t, 4, h.Size())
	})

	t.Run("Push into non empty heap should maintain heap property", func(t *testing.T) {
		h := ds.NewMaxHeap[int]()
		h.Push(10)
		p, _ := h.Peek()
		assert.Equal(t, 10, p)
		h.Push(15)
		p, _ = h.Peek()
		assert.Equal(t, 15, p)
		h.Push(30)
		p, _ = h.Peek()
		assert.Equal(t, 30, p)
		h.Push(100)
		p, _ = h.Peek()
		assert.Equal(t, 100, p)
	})

	t.Run("Underlying slice should resize when needed", func(t *testing.T) {
		h := ds.NewMaxHeap[int]()
		h.Push(10)
		h.Push(15)
		h.Push(30)
		h.Push(100)
		h.Push(5)
		h.Push(3)
		p, _ := h.Peek()
		assert.Equal(t, 100, p)
		assert.Equal(t, 6, h.Size())
	})
}

func TestMaxHeapPop(t *testing.T) {
	t.Run("Pop from an empty heap", func(t *testing.T) {
		h := ds.NewMaxHeap[int]()
		_, err := h.Pop()
		assert.NotNil(t, err)
	})

	t.Run("Pop from a heap of size 1", func(t *testing.T) {
		h := ds.NewMaxHeap[int]()
		h.Push(10)
		elem, err := h.Pop()
		assert.Nil(t, err)
		assert.Equal(t, 10, elem)
	})

	t.Run("Pop from a heap with many elements", func(t *testing.T) {
		h := ds.NewMaxHeap[int]()
		h.Push(10)
		h.Push(5)
		h.Push(2)
		h.Push(100)
		h.Push(50)
		assert.Equal(t, 5, h.Size())
		elem, err := h.Pop()
		assert.Nil(t, err)
		assert.Equal(t, 100, elem)
		p, _ := h.Peek()
		assert.Equal(t, 50, p)

		elem, err = h.Pop()
		assert.Nil(t, err)
		assert.Equal(t, 50, elem)
		p, _ = h.Peek()
		assert.Equal(t, 10, p)

		elem, err = h.Pop()
		assert.Nil(t, err)
		assert.Equal(t, 10, elem)
		p, _ = h.Peek()
		assert.Equal(t, 5, p)

		elem, err = h.Pop()
		assert.Nil(t, err)
		assert.Equal(t, 5, elem)
		p, _ = h.Peek()
		assert.Equal(t, 2, p)

		elem, err = h.Pop()
		assert.Nil(t, err)
		assert.Equal(t, 2, elem)

		assert.Equal(t, 0, h.Size())
	})
}

func TestHeapify(t *testing.T) {
	t.Run("Heapify should return a valid heap", func(t *testing.T) {
		list := []int{12, 1, 10, 5, 6, 3, 9, 11}

		h := ds.Heapify(list)

		elem, err := h.Pop()
		assert.Nil(t, err)
		assert.Equal(t, 12, elem)
	})
}
