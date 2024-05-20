package ds_test

import (
	"testing"

	"github.com/bcdxn/dsa-go/ds"
	"github.com/stretchr/testify/assert"
)

func TestNewBst(t *testing.T) {
	t.Run("should initialize a tree with a nil root", func(t *testing.T) {
		tree := ds.NewBST[int]()
		assert.Nilf(t, tree.Root, "Newly initialized BST Root should be nil")
	})

	t.Run("should initialize a tree with sie 0", func(t *testing.T) {
		tree := ds.NewBST[int]()
		assert.Equalf(t, 0, tree.Size(), "Newly intialized BST should have size of 0")
	})
}

func TestBstInsert(t *testing.T) {
	t.Run("should increase the size by 1 when 1 element is added to the tree", func(t *testing.T) {
		tree := ds.NewBST[int]()

		assert.Equalf(t, 0, tree.Size(), "Newly intialized BST should have size of 0")
		tree.Insert(10)
		assert.Equalf(t, 1, tree.Size(), "BST size should be incremented after insert")
		tree.Insert(10)
		tree.Insert(10)
		tree.Insert(10)
		assert.Equalf(t, 4, tree.Size(), "BST size should be incremented after insert")
	})
}

func TestBstFindMin(t *testing.T) {
	t.Run("should return the minimum element in the true", func(t *testing.T) {
		tree := ds.NewBST[int]()

		tree.Insert(10)
		tree.Insert(5)
		tree.Insert(7)
		tree.Insert(17)
		tree.Insert(3)
		tree.Insert(24)
		tree.Insert(13)
		tree.Insert(15)
		tree.Insert(6)
		min, err := tree.FindMin()
		assert.Nil(t, err, "Should not return error if the tree is not empty")
		assert.Equalf(t, 3, min, "3 is the minimum value in this tree")
	})
}

func TestBstRemove(t *testing.T) {
	t.Run("remove an element from an empty tree", func(t *testing.T) {
		tree := ds.NewBST[int]()
		err := tree.Remove(10)
		assert.NotNilf(t, err, "Remove should return an error if the element is not found")
		assert.Equalf(t, 0, tree.Size(), "BST size should not change if item is not removed")
	})

	t.Run("remove an element from a tree of size 1", func(t *testing.T) {
		tree := ds.NewBST[int]()
		assert.Equalf(t, 0, tree.Size(), "Newly intialized BST should have size of 0")
		tree.Insert(10)
		assert.Equalf(t, 1, tree.Size(), "BST size should be incremented after insert")
		err := tree.Remove(10)
		assert.Nilf(t, err, "Remove should not return an error if the element is  found")
		assert.Equalf(t, 0, tree.Size(), "BST size should be decremented after remove")
	})

	t.Run("remove a leaf node", func(t *testing.T) {
		tree := ds.NewBST[int]()

		tree.Insert(10)
		tree.Insert(5)
		tree.Insert(7)
		tree.Insert(17)
		tree.Insert(3)
		tree.Insert(24)
		tree.Insert(13)
		tree.Insert(15)
		tree.Insert(6)

		// root's left subtree
		assert.Equal(t, 5, tree.Root.Left.Elem)
		assert.Equal(t, 7, tree.Root.Left.Right.Elem)
		assert.Equal(t, 3, tree.Root.Left.Left.Elem)
		assert.Equal(t, 3, tree.Root.Left.Left.Elem)
		assert.Equal(t, 6, tree.Root.Left.Right.Left.Elem)
		// root's right subtree
		assert.Equal(t, 17, tree.Root.Right.Elem)
		assert.Equal(t, 13, tree.Root.Right.Left.Elem)
		assert.Equal(t, 24, tree.Root.Right.Right.Elem)
		assert.Equal(t, 15, tree.Root.Right.Left.Right.Elem)

		assert.Equal(t, 9, tree.Size())

		err := tree.Remove(6)

		if err != nil {
			t.Errorf("Removal of leaf node should not have resulted in error")
		}
		assert.Equal(t, 8, tree.Size())
		assert.Nilf(t, tree.Root.Left.Right.Left, "Node should be removed")

		err = tree.Remove(24)
		if err != nil {
			t.Errorf("Removal of leaf node should not have resulted in error")
		}
		assert.Equal(t, 7, tree.Size())
		assert.Nilf(t, tree.Root.Left.Right.Left, "Node should be removed")
	})

	t.Run("remove a node with only a left child", func(t *testing.T) {
		tree := ds.NewBST[int]()

		tree.Insert(10)
		tree.Insert(5)
		tree.Insert(7)
		tree.Insert(17)
		tree.Insert(3)
		tree.Insert(24)
		tree.Insert(13)
		tree.Insert(15)
		tree.Insert(6)

		// root's left subtree
		assert.Equal(t, 5, tree.Root.Left.Elem)
		assert.Equal(t, 7, tree.Root.Left.Right.Elem)
		assert.Equal(t, 3, tree.Root.Left.Left.Elem)
		assert.Equal(t, 3, tree.Root.Left.Left.Elem)
		assert.Equal(t, 6, tree.Root.Left.Right.Left.Elem)
		// root's right subtree
		assert.Equal(t, 17, tree.Root.Right.Elem)
		assert.Equal(t, 13, tree.Root.Right.Left.Elem)
		assert.Equal(t, 24, tree.Root.Right.Right.Elem)
		assert.Equal(t, 15, tree.Root.Right.Left.Right.Elem)

		assert.Equal(t, 9, tree.Size())

		err := tree.Remove(7)

		if err != nil {
			t.Errorf("Removal of node should not have resulted in error")
		}
		assert.Equal(t, 8, tree.Size())
		assert.Equal(t, 6, tree.Root.Left.Right.Elem, "Left child should be promoted")
	})

	t.Run("remove a node with only a right child", func(t *testing.T) {
		tree := ds.NewBST[int]()

		tree.Insert(10)
		tree.Insert(5)
		tree.Insert(7)
		tree.Insert(17)
		tree.Insert(3)
		tree.Insert(24)
		tree.Insert(13)
		tree.Insert(15)
		tree.Insert(6)

		// root's left subtree
		assert.Equal(t, 5, tree.Root.Left.Elem)
		assert.Equal(t, 7, tree.Root.Left.Right.Elem)
		assert.Equal(t, 3, tree.Root.Left.Left.Elem)
		assert.Equal(t, 3, tree.Root.Left.Left.Elem)
		assert.Equal(t, 6, tree.Root.Left.Right.Left.Elem)
		// root's right subtree
		assert.Equal(t, 17, tree.Root.Right.Elem)
		assert.Equal(t, 13, tree.Root.Right.Left.Elem)
		assert.Equal(t, 24, tree.Root.Right.Right.Elem)
		assert.Equal(t, 15, tree.Root.Right.Left.Right.Elem)

		assert.Equal(t, 9, tree.Size())

		err := tree.Remove(13)

		if err != nil {
			t.Errorf("Removal of node should not have resulted in error")
		}
		assert.Equal(t, 8, tree.Size())
		assert.Equal(t, 15, tree.Root.Right.Left.Elem, "Left child should be promoted")
	})

	t.Run("remove a node with both left and right child", func(t *testing.T) {
		tree := ds.NewBST[int]()

		tree.Insert(10)
		tree.Insert(5)
		tree.Insert(7)
		tree.Insert(17)
		tree.Insert(3)
		tree.Insert(24)
		tree.Insert(13)
		tree.Insert(15)
		tree.Insert(6)

		// root's left subtree
		assert.Equal(t, 5, tree.Root.Left.Elem)
		assert.Equal(t, 7, tree.Root.Left.Right.Elem)
		assert.Equal(t, 3, tree.Root.Left.Left.Elem)
		assert.Equal(t, 3, tree.Root.Left.Left.Elem)
		assert.Equal(t, 6, tree.Root.Left.Right.Left.Elem)
		// root's right subtree
		assert.Equal(t, 17, tree.Root.Right.Elem)
		assert.Equal(t, 13, tree.Root.Right.Left.Elem)
		assert.Equal(t, 24, tree.Root.Right.Right.Elem)
		assert.Equal(t, 15, tree.Root.Right.Left.Right.Elem)

		assert.Equal(t, 9, tree.Size())

		err := tree.Remove(17)

		if err != nil {
			t.Errorf("Removal of node should not have resulted in error")
		}
		assert.Equal(t, 8, tree.Size())
		assert.Equal(t, 24, tree.Root.Right.Elem, "Left child should be promoted")
		assert.Equal(t, 13, tree.Root.Right.Left.Elem, "Left child should be promoted")
		assert.Equal(t, 15, tree.Root.Right.Left.Right.Elem, "Left child should be promoted")

		err = tree.Remove(5)
		if err != nil {
			t.Errorf("Removal of node should not have resulted in error")
		}
		assert.Equal(t, 7, tree.Size())
		assert.Equal(t, 6, tree.Root.Left.Elem, "Left child should be promoted")
	})
}

func TestBstInOrder(t *testing.T) {
	t.Run("should print the nodes in order", func(t *testing.T) {
		tree := ds.NewBST[int]()

		tree.Insert(10)
		tree.Insert(5)
		tree.Insert(7)
		tree.Insert(17)
		tree.Insert(3)
		tree.Insert(24)
		tree.Insert(13)
		tree.Insert(15)
		tree.Insert(6)

		tree.InOrder()
	})
}

func TestBstPreOrder(t *testing.T) {
	t.Run("should print in pre-order", func(t *testing.T) {
		tree := ds.NewBST[int]()

		tree.Insert(10)
		tree.Insert(5)
		tree.Insert(7)
		tree.Insert(17)
		tree.Insert(3)
		tree.Insert(24)
		tree.Insert(13)
		tree.Insert(15)
		tree.Insert(6)

		tree.PreOrder()
	})
}

func TestBstPostOrder(t *testing.T) {
	t.Run("should print in post-order", func(t *testing.T) {
		tree := ds.NewBST[int]()

		tree.Insert(10)
		tree.Insert(5)
		tree.Insert(7)
		tree.Insert(17)
		tree.Insert(3)
		tree.Insert(24)
		tree.Insert(13)
		tree.Insert(15)
		tree.Insert(6)

		tree.PostOrder()
	})
}
