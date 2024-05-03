package ds

import "golang.org/x/exp/constraints"

// Stack is a simple data structure that offers LIFO access to its elements.
type Stack[T constraints.Ordered] struct {
	l List[T]
}

// NewStack returns an empty stack
func NewStack[T constraints.Ordered]() *Stack[T] {
	l := NewList[T]()

	return &Stack[T]{
		l: *l,
	}
}

// Height returns the number of elements in the stack.
func (s Stack[T]) Height() int {
	return s.l.Len()
}

// Push adds an item to the top of the stack
func (s *Stack[T]) Push(elem T) {
	s.l.AddHead(elem)
}

// Pop removes an item from the top of the stack
func (s *Stack[T]) Pop() (T, error) {
	return s.l.RemoveHead()
}
