package ds

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// MaxHeap implements a max heap data structure, ensuring a complete tree that satisfies the max
// heap property
type MaxHeap[T constraints.Ordered] struct {
	s         *[]T
	lastIndex int
	size      int
}

// NewMaxHeap returns a new empty heap
func NewMaxHeap[T constraints.Ordered]() *MaxHeap[T] {
	s := make([]T, 5)
	return &MaxHeap[T]{
		s:         &s,
		lastIndex: 0,
		size:      0,
	}
}

// Size returns the number of elements currently in the heap
func (h MaxHeap[T]) Size() int {
	return h.size
}

// Peek returns the element at the top of the heap without removing it
func (h MaxHeap[T]) Peek() (T, error) {
	if h.Size() == 0 {
		var empty T
		return empty, errors.New("cannot peek an empty heap")
	}
	return (*h.s)[1], nil
}

// Push adds an element to the heap and maintains the max heap property
func (h *MaxHeap[T]) Push(elem T) {
	if h.Size() == (len(*h.s) - 1) {
		// increase the size of the underlying slice and copy the values over to the new slice
		s := make([]T, len(*h.s)*2)
		copy(s, *h.s)
		h.s = &s
	}
	// Add the element to the end of the underlying slice and increment size
	h.lastIndex++
	(*h.s)[h.lastIndex] = elem
	h.size++
	// Ensure heap property (percolate up)
	parentIndex := h.parentIndex(h.lastIndex)
	childIndex := h.lastIndex
	for parentIndex > 0 {
		if (*h.s)[parentIndex] < (*h.s)[childIndex] {
			// nodes do not satisfy the heap property and must be swapped
			(*h.s)[parentIndex], (*h.s)[childIndex] = (*h.s)[childIndex], (*h.s)[parentIndex]
			childIndex = parentIndex
			parentIndex = h.parentIndex(parentIndex)
		} else {
			// nodes satisfy the heap property
			break
		}
	}
}

// Pop removes the element at the top of the heap and maintains the max heap property
func (h *MaxHeap[T]) Pop() (T, error) {
	var empty T
	if h.Size() < 1 {
		return empty, errors.New("cannot pop an empty tree")
	}
	// Remove the element at the root of the heap
	elem := (*h.s)[1]
	h.size--
	// move the last child to the root to maintain a complete tree
	(*h.s)[1] = (*h.s)[h.lastIndex]
	(*h.s)[h.lastIndex] = empty
	h.lastIndex--
	// Ensure heap property (percolate down)
	currIndex := 1
	for h.leftChildIndex(currIndex) <= h.lastIndex {
		left := h.leftChildIndex(currIndex)
		right := h.rightChildIndex(currIndex)

		var minChildIndex int

		if right <= h.lastIndex && (*h.s)[right] > (*h.s)[left] {
			minChildIndex = right
		} else if left <= h.lastIndex {
			minChildIndex = left
		} else {
			// the node has no children and therefore satisfies the heap property
			break
		}

		if (*h.s)[currIndex] < (*h.s)[minChildIndex] {
			// nodes do not satisfy the heap property and must be swapped
			(*h.s)[currIndex], (*h.s)[minChildIndex] = (*h.s)[minChildIndex], (*h.s)[currIndex]
			currIndex = minChildIndex
		} else {
			// nodes satisfy the heap property
			break
		}
	}

	return elem, nil
}

// leftChildIndex returns in the index of the left child of the node at the given index
func (h MaxHeap[T]) leftChildIndex(i int) int {
	return i * 2
}

// rightChildIndex returns in the index of the left child of the node at the given index
func (h MaxHeap[T]) rightChildIndex(i int) int {
	return (i * 2) + 1
}

// parentIndex returns in the index of the left child of the node at the given index
func (h MaxHeap[T]) parentIndex(i int) int {
	return i / 2
}
