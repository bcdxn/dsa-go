package ds

import "errors"

type LinkedListNode struct {
	Elem int
	Next *LinkedListNode
}

func newLinkedListNode(elem int) *LinkedListNode {
	return &LinkedListNode{
		Elem: elem,
		Next: nil,
	}
}

type LinkedList struct {
	Head *LinkedListNode
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		size: 0,
	}
}

// Size returns the number of elements in the list.
func (l LinkedList) Size() int {
	return l.size
}

// Add adds an item to the head of the list.
func (l *LinkedList) Add(elem int) {
	node := newLinkedListNode(elem)
	node.Next = l.Head
	l.Head = node
	l.size++
}

// Remove removes the first occurence of the specified element from the list.
func (l *LinkedList) Remove(elem int) (int, error) {
	node := l.Head
	var prevNode *LinkedListNode = nil

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

	return 0, errors.New("the specified element does not exist in the list")
}

// Reverse reverses the order of the linked list
func (l *LinkedList) Reverse() {
	node := l.Head
	var prevNode *LinkedListNode = nil

	for node != nil {
		// store next before reconfiguring pointers
		var tmp *LinkedListNode = node.Next
		// reconfigure pointers
		node.Next = prevNode
		l.Head = node
		// iterate
		prevNode = node
		node = tmp
	}
}

func (l *LinkedList) ReverseRecursive() {
	reverseRecursive(l, l.Head)
}

func reverseRecursive(l *LinkedList, node *LinkedListNode) {
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
