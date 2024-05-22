package ds

import (
	"container/list"
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

type TreeNode[T constraints.Ordered] struct {
	Elem  T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func newTreeNode[T constraints.Ordered](elem T) *TreeNode[T] {
	return &TreeNode[T]{
		Elem:  elem,
		Left:  nil,
		Right: nil,
	}
}

// BST implements a Binary Search Tree (BST)
type BST[T constraints.Ordered] struct {
	Root *TreeNode[T]
	size int
}

// NewBST creates a new Binary Search Tree (BST)
func NewBST[T constraints.Ordered]() *BST[T] {
	return &BST[T]{
		Root: nil,
		size: 0,
	}
}

func (t BST[T]) Size() int {
	return t.size
}

// FindMin returns the node with the smallest element value. If the tree is empty an error is
// returned.
func (t BST[T]) FindMin() (T, error) {
	var elem T

	if t.Root == nil {
		return elem, errors.New("tree is empty")
	}

	smallestNode := t.findMin(t.Root)

	return smallestNode.Elem, nil
}

// Insert adds a node to the tree and maintains the BST properties.
func (t *BST[T]) Insert(elem T) {
	t.Root = t.insert(t.Root, elem)
}

// Remove deletes a node from the tree and maintains the BST properties.
func (t *BST[T]) Remove(elem T) error {
	root, err := t.remove(t.Root, elem)
	if err != nil {
		t.Root = root
	}

	return err
}

/* Collection of Traversals
------------------------------------------------------------------------------------------------- */

func (t BST[T]) InOrder() {
	fmt.Print("[ ")
	inOrder(t.Root)
	fmt.Print("]\n")
}

func (t BST[T]) PreOrder() {
	fmt.Print("[ ")
	preOrder(t.Root)
	fmt.Print("]\n")
}

func (t BST[T]) PostOrder() {
	fmt.Print("[ ")
	postOrder(t.Root)
	fmt.Print("]\n")
}

func (t BST[T]) BreadthFirst() {
	if t.Root == nil {
		fmt.Println("[ ]")
		return
	}

	q := list.New()
	currNode := t.Root
	q.PushBack(currNode)

	fmt.Print("[ ")
	for currNode != nil {
		currNode = nil
		l := q.Len()
		for i := 0; i < l; i++ {
			// Dequeue node to process
			if n, ok := any(q.Front().Value).(*TreeNode[T]); ok {
				currNode = n
				q.Remove(q.Front())
			} else {
				panic("invalid type in BFS queue ")
			}
			// process current node
			fmt.Print(currNode.Elem, " ")
			// Add children to queue to be processed
			if currNode.Left != nil {
				// enqueue left child
				q.PushBack(currNode.Left)
			}
			if currNode.Right != nil {
				// enqueue right child
				q.PushBack(currNode.Right)
			}
		}
	}
	fmt.Print("]\n")
}

/* Private helper functions
------------------------------------------------------------------------------------------------- */

// insert is a recursive helper function to insert an element into the tree.
func (t *BST[T]) insert(root *TreeNode[T], elem T) *TreeNode[T] {
	// base case
	if root == nil {
		// we can safely insert the element as a new leaf
		t.size++
		return newTreeNode(elem)
	}

	if elem < root.Elem {
		// The element belongs in the left subtree of the current element
		root.Left = t.insert(root.Left, elem)
	} else {
		// The element belongs in the right subtree of the current element
		root.Right = t.insert(root.Right, elem)
	}

	return root
}

// remove is a recursive helper function to remove a node from the tree.
func (t *BST[T]) remove(root *TreeNode[T], elem T) (*TreeNode[T], error) {
	if root == nil {
		// The element wasn't found in the tree
		return root, errors.New("element was not found")
	}

	var err error = nil
	var node *TreeNode[T]

	if elem < root.Elem {
		root.Left, err = t.remove(root.Left, elem)
	} else if elem > root.Elem {
		root.Right, err = t.remove(root.Right, elem)
	} else {
		// We've found the node to remove
		if root.Left == nil {
			// Our node has, at most, only a single child which can be promoted
			node = root.Right
		} else if root.Right == nil {
			// Our node has, at most, only a single child which can be promoted
			node = root.Left
		} else {
			// Our node has 2 children; we should replace it with the smallest in the right subtree
			node = t.findMin(root.Right)
			// Remove the node from the right subtree since we're going to add it at the root of the
			// current subtree
			root.Right, err = t.remove(root.Right, node.Elem)
			node.Left = root.Left
			node.Right = root.Right
			// removing the minimum from the right subtree will decrement our size but we're really just
			// moving it up in the tree to replace the original node to be removed
			t.size++
		}
		t.size--
		return node, err
	}

	return root, err
}

func (t BST[T]) findMin(root *TreeNode[T]) *TreeNode[T] {
	if root == nil {
		return root
	}

	if root.Left == nil {
		return root
	}

	return t.findMin(root.Left)
}

func inOrder[T constraints.Ordered](root *TreeNode[T]) {
	if root == nil {
		return
	}

	inOrder(root.Left)
	fmt.Print(root.Elem, " ")
	inOrder(root.Right)
}

func preOrder[T constraints.Ordered](root *TreeNode[T]) {
	if root == nil {
		return
	}

	fmt.Print(root.Elem, " ")
	preOrder(root.Left)
	preOrder(root.Right)
}

func postOrder[T constraints.Ordered](root *TreeNode[T]) {
	if root == nil {
		return
	}

	postOrder(root.Left)
	postOrder(root.Right)
	fmt.Print(root.Elem, " ")
}
