package ds

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// AVLTreeNode is an implementation of a node to be used on self-balancing binary search trees.
type AVLTreeNode[T constraints.Ordered] struct {
	Elem   T
	Left   *AVLTreeNode[T]
	Right  *AVLTreeNode[T]
	height int
}

// NewAVLTreeNode returns a pointer to a new AVL Tree Node.
func NewAVLTreeNode[T constraints.Ordered](elem T) *AVLTreeNode[T] {
	return &AVLTreeNode[T]{
		Elem:   elem,
		Left:   nil,
		Right:  nil,
		height: 0,
	}
}

// GetHeight returns the height of the specified node.
func (n AVLTreeNode[T]) GetHeight() int {
	leftHeight, rightHeight := n.getChildrenHeights()
	return max(leftHeight, rightHeight) + 1
}

// GetBalanceFactor returns the balance factor of a particular node, telling us if it is balanced,
// left-heavy, or right-heavy.
func (n AVLTreeNode[T]) GetBalanceFactor() int {
	leftHeight, rightHeight := n.getChildrenHeights()
	return leftHeight - rightHeight
}

func (n AVLTreeNode[T]) getChildrenHeights() (int, int) {
	leftHeight := -1
	rightHeight := -1
	if n.Left != nil {
		leftHeight = n.Left.height
	}
	if n.Right != nil {
		rightHeight = n.Right.height
	}

	return leftHeight, rightHeight
}

// AVLTree is an implementation of a self-balancing binary search tree.
type AVLTree[T constraints.Ordered] struct {
	Root *AVLTreeNode[T]
	size int
}

// NewAVLTree returns a pointer to a new AVL Tree.
func NewAVLTree[T constraints.Ordered]() *AVLTree[T] {
	return &AVLTree[T]{
		Root: nil,
		size: 0,
	}
}

// Size returns the number of nodes in the tree.
func (t AVLTree[T]) Size() int {
	return t.size
}

// Contains returns true if the specified element is in the tree; else it returns false.
func (t AVLTree[T]) Contains(elem T) bool {
	node := t.contains(t.Root, elem)

	return node != nil
}

func (t *AVLTree[T]) Insert(elem T) error {
	if node, err := t.insert(t.Root, elem); err != nil {
		return err
	} else {
		t.Root = node
		t.size++
		return nil
	}
}

/* Private helper functions
------------------------------------------------------------------------------------------------- */

func (t AVLTree[T]) contains(root *AVLTreeNode[T], elem T) *AVLTreeNode[T] {
	if root == nil {
		return root
	}

	if elem == root.Elem {
		return root
	}

	if elem < root.Elem {
		return t.contains(root.Left, elem)
	} else {
		return t.contains(root.Right, elem)
	}
}

func (t *AVLTree[T]) insert(root *AVLTreeNode[T], elem T) (*AVLTreeNode[T], error) {
	if root == nil {
		return NewAVLTreeNode(elem), nil
	}

	if elem < root.Elem {
		if node, err := t.insert(root.Left, elem); err != nil {
			return nil, err
		} else {
			root.Left = node
		}
	} else if elem > root.Elem {
		if node, err := t.insert(root.Right, elem); err != nil {
			return nil, err
		} else {
			root.Right = node
		}
	} else {
		return nil, errors.New("cannot insert a duplicate element")
	}

	// Check the balance factor
	bFactor := root.GetBalanceFactor()

	if bFactor < -1 {
		// tree is right-heavy
		rightBFactor := root.Right.GetBalanceFactor()
		if rightBFactor > 0 {
			// the right subtree is left heavy; we must do a right-left rotation
			root.Right = t.rightRotation(root.Right)
		}
		root = t.leftRotation(root)
	} else if bFactor > 1 {
		// tree is left-heavy
		leftBFactor := root.Left.GetBalanceFactor()
		if leftBFactor < 0 {
			// the left subtree is right heavy; we must do a left-right rotation
			root.Left = t.leftRotation(root.Left)
		}
		root = t.rightRotation(root)
	}

	// Set the height
	root.height = root.GetHeight()

	return root, nil
}

// leftRotation will perform an AVL Left Rotation and return the new post-rotation root
func (t *AVLTree[T]) leftRotation(root *AVLTreeNode[T]) *AVLTreeNode[T] {
	node := root.Right
	root.Right = node.Left
	node.Left = root
	// Recalculate heights for the affected nodes
	root.height = root.GetHeight()
	node.height = node.GetHeight()
	// return the new root
	return node
}

// rightRotation will perform an AVL Right Rotation and return the new post-rotation root
func (t *AVLTree[T]) rightRotation(root *AVLTreeNode[T]) *AVLTreeNode[T] {
	node := root.Left
	root.Left = node.Right
	node.Right = root
	// Recalculate heights for the affected nodes
	root.height = root.GetHeight()
	node.height = node.GetHeight()
	// return the new root
	return node
}
