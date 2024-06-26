package ds

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// ListNode is an implementation of the building block of a Doubly Linked List.
type ListNode[T constraints.Ordered] struct {
	Elem T
	Prev *ListNode[T]
	Next *ListNode[T]
}

func newListNode[T constraints.Ordered](elem T) *ListNode[T] {
	return &ListNode[T]{
		Elem: elem,
		Prev: nil,
		Next: nil,
	}
}

// List is an implementation of the Doubly Linked List data structure.
type List[T constraints.Ordered] struct {
	Head *ListNode[T]
	Tail *ListNode[T]
	len  int
}

// NewList creates and returns an empty doubly linked list.
func NewList[T constraints.Ordered]() *List[T] {
	return &List[T]{
		Head: nil,
		Tail: nil,
		len:  0,
	}
}

// Len returns the length of the list.
func (l List[T]) Len() int {
	return l.len
}

// AddTail adds an element to the Head of the list.
func (l *List[T]) AddHead(elem T) {
	node := newListNode(elem)

	if l.Len() < 1 {
		l.Head = node
	} else {
		node.Next = l.Head
		l.Head.Prev = node
		l.Head = node
	}

	l.len++

	if l.Len() < 2 {
		l.Tail = l.Head
	}
}

// AddTail adds an element to the Tail of the list.
func (l *List[T]) AddTail(elem T) {
	node := newListNode(elem)

	if l.Len() < 1 {
		l.Tail = node
	} else {
		node.Prev = l.Tail
		l.Tail.Next = node
		l.Tail = node
	}

	l.len++

	if l.Len() < 2 {
		l.Head = l.Tail
	}
}

// RemoveHead removes the first element in the list
func (l *List[T]) RemoveHead() (T, error) {
	var elem T
	node := l.Head
	// Ensure we're not removing from an empty list
	if node == nil {
		return elem, errors.New("attempted to remove from an empty list")
	}
	// store the elem that is being removed so we can return it
	elem = node.Elem
	if node.Next == nil {
		// We're dealing with a list of length 1
		l.Head = nil
		l.Tail = nil
	} else {
		// We're dealing with a list of length > 1
		// update list pointers
		node.Next.Prev = nil
		l.Head = node.Next
		node.Next = nil // clear to help GC
	}

	l.len--

	return elem, nil
}

// RemoveTail removes the tail element in the list
func (l *List[T]) RemoveTail() (T, error) {
	var elem T
	node := l.Tail
	// Ensure we're not removing from an empty list
	if node == nil {
		return elem, errors.New("attempted to remove from an empty list")
	}
	// store the elem that is being removed so we can return it
	elem = node.Elem
	if node.Prev == nil {
		// We're dealing with a list of length 1
		l.Head = nil
		l.Tail = nil
	} else {
		// We're dealing with a list of length > 1
		// update list pointers
		node.Prev.Next = nil
		l.Tail = node.Prev
		node.Prev = nil // clear to help GC
	}

	l.len--

	return elem, nil
}

// Remove finds the specified element and removes it from the list; if the element is not found an
// error is returned indicating as such.
func (l *List[T]) Remove(elem T) (T, error) {
	node := l.Head
	var retElem T

	for node != nil {
		if node.Elem == elem {
			retElem = node.Elem

			if node.Prev == nil && node.Next == nil {
				// We're removing the only element in the list
				l.Head = nil
				l.Tail = nil
			} else {
				if node.Prev == nil {
					// we're removing the Head
					l.Head = node.Next
					node.Next.Prev = nil
				} else if node.Next == nil {
					// we're removing the Tail
					l.Tail = node.Prev
					l.Tail.Next = nil
				} else {
					// Removing from somewhere in the middle of the list; we can simply configure the
					// previous node to point to the current node's next and the next node's previous
					// to the current node's previous, effectively 'skipping' the node
					node.Prev.Next = node.Next
					node.Next.Prev = node.Prev
				}
				// clear pointers of the node being deleted to aid in GC
				node.Next = nil
				node.Prev = nil
			}
			// decrement len
			l.len--

			return retElem, nil
		}
		node = node.Next
	}

	return retElem, errors.New("cannot find element")
}

// Contains returns true if the specified element is in the list and false if the specified element
// is not found in the list.
func (l List[T]) Contains(elem T) bool {
	node := l.Head

	for node != nil {
		if node.Elem == elem {
			return true
		}
		node = node.Next
	}

	return false
}

// Reverse reverses the linked list in place.
func (l *List[T]) Reverse() {
	node := l.Head

	for node != nil {
		// store Next before any pointer reconfig so we can iterate
		next := node.Next
		// Reconfigure pointers

		if node.Prev == nil && node.Next == nil {
			// we have list of length 1 which is implicitly reversed
			return
		}

		if node.Prev == nil {
			// We're at the head, which becomes the tail
			l.Tail = node
			node.Next = nil
		} else {
			node.Next = node.Prev
		}

		if next == nil {
			// we're at the tail, which becomes the head
			l.Head = node
		} else {
			node.Prev = next
		}

		// Iterate
		node = next
	}
}
