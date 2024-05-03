package ds

import (
	"errors"

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

// FindMin returns the node with the smallest element value.
func (t BST[T]) FindMin(root *TreeNode[T]) *TreeNode[T] {
	if root == nil {
		return root
	}

	if root.Left == nil {
		return root
	}

	return t.FindMin(root.Left)
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
		if root.Left == nil && root.Right == nil {
			// We're at a leave node and can safely remove it without any pointer reconfiguration
			node = nil
		} else if root.Left == nil {
			// Our node has only a single child which can be promoted
			node = root.Right
		} else if root.Right == nil {
			// Our node has only a single child which can be promoted
			node = root.Left
		} else {
			// Our node has 2 children; we should replace it with the smallest in the right subtree
			node = t.FindMin(root.Right)
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
