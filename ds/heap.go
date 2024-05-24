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
	pi := parentIndex(h.lastIndex)
	ci := h.lastIndex
	for pi > 0 {
		if (*h.s)[pi] < (*h.s)[ci] {
			// nodes do not satisfy the heap property and must be swapped
			(*h.s)[pi], (*h.s)[ci] = (*h.s)[ci], (*h.s)[pi]
			ci = pi
			pi = parentIndex(pi)
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
	for leftChildIndex(currIndex) <= h.lastIndex {
		left := leftChildIndex(currIndex)
		right := rightChildIndex(currIndex)

		var maxChildIndex int

		if right <= h.lastIndex && (*h.s)[right] > (*h.s)[left] {
			maxChildIndex = right
		} else if left <= h.lastIndex {
			maxChildIndex = left
		} else {
			// the node has no children and therefore satisfies the heap property
			break
		}

		if (*h.s)[currIndex] < (*h.s)[maxChildIndex] {
			// nodes do not satisfy the heap property and must be swapped
			(*h.s)[currIndex], (*h.s)[maxChildIndex] = (*h.s)[maxChildIndex], (*h.s)[currIndex]
			currIndex = maxChildIndex
		} else {
			// nodes satisfy the heap property
			break
		}
	}

	return elem, nil
}

// Heapify converts the given list into a valid heap in O(n) time.
func Heapify[T constraints.Ordered](list []T) *MaxHeap[T] {
	h := NewMaxHeap[T]()
	h.heapify(list)
	return h
}

// heapify is a helper function to heapify a given list (giving us access to the private internal
// slice that the heap is implemented on)
func (h *MaxHeap[T]) heapify(list []T) *MaxHeap[T] {
	if len(list) < 1 {
		return h
	}
	h.size = len(list)
	// Move 0th element in the list to the end to make pointer arthmetic simpler
	list = append(list, list[0])
	var empty T
	list[0] = empty
	h.s = &list
	// We only need to 'percolate' elements with children; half the nodes
	// in the heap will be leaf nodes
	currIndex := len(*h.s) / 2
	// Stop the loop when we've reached the second level of the heap (the root can't percolate up
	// any more)
	for currIndex > 0 {
		percolateIndex := currIndex
		for percolateIndex < len(*h.s)-1 {
			li := leftChildIndex(percolateIndex)
			ri := rightChildIndex(percolateIndex)
			maxChildIndex := li
			// Find the larger of the two children
			if ri < len(*h.s) && (*h.s)[ri] > (*h.s)[li] {
				maxChildIndex = ri
			}
			// swap with largest child (if largest child is greater than current node)
			if maxChildIndex < len(*h.s) && (*h.s)[percolateIndex] < (*h.s)[maxChildIndex] {
				(*h.s)[maxChildIndex], (*h.s)[percolateIndex] = (*h.s)[percolateIndex], (*h.s)[maxChildIndex]
				percolateIndex = maxChildIndex
			} else {
				// the heap property is satisfied at the current node
				break
			}
		}
		// iterate 'up' the array to the root
		currIndex--
	}

	return h
}

// leftChildIndex returns in the index of the left child of the node at the given index
func leftChildIndex(i int) int {
	return i * 2
}

// rightChildIndex returns in the index of the left child of the node at the given index
func rightChildIndex(i int) int {
	return (i * 2) + 1
}

// parentIndex returns in the index of the left child of the node at the given index
func parentIndex(i int) int {
	return i / 2
}
