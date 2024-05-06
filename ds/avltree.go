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

// Insert adds an element into the tree and maintains the binary search and AVL balance properties
func (t *AVLTree[T]) Insert(elem T) error {
	if node, err := t.insert(t.Root, elem); err != nil {
		return err
	} else {
		t.Root = node
		t.size++
		return nil
	}
}

// Remove removes an element from the tree and maintains the binary search and AVL balance
// properties
func (t *AVLTree[T]) Remove(elem T) error {
	if node, err := t.remove(t.Root, elem); err != nil {
		return err
	} else {
		t.Root = node
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
	// Calculate current height of subtree before balancing
	root.height = root.GetHeight()
	// Ensure tree is still balanced
	root = t.balance(root)
	// Set the height after balance
	root.height = root.GetHeight()

	return root, nil
}

func (t *AVLTree[T]) remove(root *AVLTreeNode[T], elem T) (*AVLTreeNode[T], error) {
	if root == nil {
		return nil, errors.New("element not found in the tree")
	}

	if elem < root.Elem {
		if node, err := t.remove(root.Left, elem); err != nil {
			return root, err
		} else {
			root.Left = node
			root.height = root.GetHeight()
		}
	} else if elem > root.Elem {
		if node, err := t.remove(root.Right, elem); err != nil {
			return root, err
		} else {
			root.Right = node
			root.height = root.GetHeight()
		}
	} else {
		// we've found the node to be removed
		if root.Left == nil {
			root = root.Right
		} else if root.Right == nil {
			root = root.Left
		} else {
			// Promote the minimum child in the right subtree
			promotionNode := t.findMin(root.Right)
			if node, err := t.remove(root.Right, promotionNode.Elem); err != nil {
				return root, err
			} else {
				promotionNode.Left = root.Left
				promotionNode.Right = node
				root = promotionNode
				// The remove call above will have decremented size, but we're really just moving it up in
				// the tree; let's add back the decremented size
				t.size++
			}
		}
		t.size--
	}

	if root != nil {
		root.height = root.GetHeight()
		// Ensure tree is still balanced
		root = t.balance(root)
		root.height = root.GetHeight()
	}

	return root, nil
}

// balance ensures the subtree at the given root is balanced in accordance with the properties of an
// AVL tree. This function may mutate the structure of the tree if it is unbalanced.
func (t *AVLTree[T]) balance(root *AVLTreeNode[T]) *AVLTreeNode[T] {
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

	return root
}

func (t *AVLTree[T]) findMin(root *AVLTreeNode[T]) *AVLTreeNode[T] {
	if root == nil {
		return root
	}

	if root.Left != nil {
		return t.findMin(root.Left)
	} else {
		return root
	}
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
