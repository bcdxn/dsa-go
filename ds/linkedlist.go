package ds

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type LinkedListNode[T constraints.Ordered] struct {
	Elem T
	Next *LinkedListNode[T]
}

func newLinkedListNode[T constraints.Ordered](elem T) *LinkedListNode[T] {
	return &LinkedListNode[T]{
		Elem: elem,
		Next: nil,
	}
}

type LinkedList[T constraints.Ordered] struct {
	Head *LinkedListNode[T]
	size int
}

func NewLinkedList[T constraints.Ordered]() *LinkedList[T] {
	return &LinkedList[T]{
		Head: nil,
		size: 0,
	}
}

// Size returns the number of elements in the list.
func (l LinkedList[T]) Size() int {
	return l.size
}

// Add adds an item to the head of the list.
func (l *LinkedList[T]) Add(elem T) {
	node := newLinkedListNode(elem)
	node.Next = l.Head
	l.Head = node
	l.size++
}

// Remove removes the first occurence of the specified element from the list.
func (l *LinkedList[T]) Remove(elem T) (T, error) {
	node := l.Head
	var prevNode *LinkedListNode[T] = nil

	for node != nil {
		if node.Elem == elem {
			// we've found the element in the list and can delete it
			if prevNode == nil {
				// if prevNode is nil then we're at the head of the list
				l.Head = l.Head.Next
			} else {
				// else we're in the middle of the list and need to set the previous node's pointer to the
				// next node to effectively remove the current node from the list
				prevNode.Next = node.Next
			}
			// decrement the size since we remoed a node
			l.size--
			// return the deleted element
			return elem, nil
		}
		// iterate
		prevNode = node
		node = node.Next
	}

	var e T
	return e, errors.New("the specified element does not exist in the list")
}

// Reverse reverses the order of the linked list
func (l *LinkedList[T]) Reverse() {
	node := l.Head
	var prevNode *LinkedListNode[T] = nil

	for node != nil {
		// store next before reconfiguring pointers
		var tmp *LinkedListNode[T] = node.Next
		// reconfigure pointers
		node.Next = prevNode
		l.Head = node
		// iterate
		prevNode = node
		node = tmp
	}
}

func (l *LinkedList[T]) ReverseRecursive() {
	reverseRecursive(l, l.Head)
}

func reverseRecursive[T constraints.Ordered](l *LinkedList[T], node *LinkedListNode[T]) {
	// base case: A list of length 1 is implicitly reversed
	if node.Next == nil {
		l.Head = node
		return
	}

	// Recursive Case
	reverseRecursive(l, node.Next)
	// Everything 'before' Node's next is already reversed so we just need to reverse the connection
	// of next node's next pointer
	node.Next.Next = node
	node.Next = nil
}
