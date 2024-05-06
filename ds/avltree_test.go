package ds_test

import (
	"testing"

	"github.com/bcdxn/dsa-go/ds"
	"github.com/stretchr/testify/assert"
)

func TestNewAvlTree(t *testing.T) {
	t.Run("should initialize a tree with a nil root", func(t *testing.T) {
		tree := ds.NewAVLTree[int]()
		assert.Nilf(t, tree.Root, "Newly initialized AVLTree Root should be nil")
	})

	t.Run("should initialize a tree with sie 0", func(t *testing.T) {
		tree := ds.NewAVLTree[int]()
		assert.Equalf(t, 0, tree.Size(), "Newly intialized AVLTree should have size of 0")
	})
}

func TestAVLTreeContains(t *testing.T) {
	t.Run("Contains on an empty tree should return false", func(t *testing.T) {
		tree := ds.NewAVLTree[int]()
		contains := tree.Contains(10)
		assert.False(t, contains)
	})
}

func TestAvlTreeInsert(t *testing.T) {
	t.Run("Insert in into an empty tree", func(t *testing.T) {
		tree := ds.NewAVLTree[int]()
		err := tree.Insert(10)
		assert.Nil(t, err)
		assert.Equal(t, 1, tree.Size())
		assert.Equal(t, 10, tree.Root.Elem)
		assert.Nil(t, tree.Root.Left)
		assert.Nil(t, tree.Root.Right)
	})
	t.Run("Insert into left subtree", func(t *testing.T) {
		tree := ds.NewAVLTree[int]()

		tree.Insert(10)
		err := tree.Insert(5)
		assert.Nil(t, err)
		assert.Equal(t, 2, tree.Size())
		assert.Equal(t, 5, tree.Root.Left.Elem)
		assert.Nil(t, tree.Root.Right)
	})

	t.Run("Insert into left subtree", func(t *testing.T) {
		tree := ds.NewAVLTree[int]()

		tree.Insert(10)
		err := tree.Insert(15)
		assert.Nil(t, err)
		assert.Equal(t, 2, tree.Size())
		assert.Equal(t, 15, tree.Root.Right.Elem)
		assert.Nil(t, tree.Root.Left)
	})

	t.Run("Insert into left heavy tree", func(t *testing.T) {
		tree := ds.NewAVLTree[int]()

		tree.Insert(10)
		tree.Insert(5)
		err := tree.Insert(1)
		assert.Nil(t, err)
		assert.Equal(t, 3, tree.Size())
		assert.Equal(t, 5, tree.Root.Elem)
		assert.Equal(t, 1, tree.Root.Left.Elem)
		assert.Equal(t, 10, tree.Root.Right.Elem)
	})

	t.Run("Insert into righ heavy tree", func(t *testing.T) {
		tree := ds.NewAVLTree[int]()

		tree.Insert(10)
		tree.Insert(15)
		err := tree.Insert(17)
		assert.Nil(t, err)
		assert.Equal(t, 3, tree.Size())
		assert.Equal(t, 15, tree.Root.Elem)
		assert.Equal(t, 10, tree.Root.Left.Elem)
		assert.Equal(t, 17, tree.Root.Right.Elem)
	})

	t.Run("Binary search tree properties should be maintained with left rotation", func(t *testing.T) {
		tree := ds.NewAVLTree[int]()

		tree.Insert(10)
		tree.Insert(5)
		tree.Insert(15)
		tree.Insert(3)
		tree.Insert(8)

		assert.Equal(t, 10, tree.Root.Elem)
		assert.Equal(t, 5, tree.Root.Left.Elem)
		assert.Equal(t, 15, tree.Root.Right.Elem)
		assert.Equal(t, 3, tree.Root.Left.Left.Elem)
		assert.Equal(t, 8, tree.Root.Left.Right.Elem)

		err := tree.Insert(1)

		assert.Nil(t, err)
		assert.Equal(t, 5, tree.Root.Elem)
		assert.Equal(t, 3, tree.Root.Left.Elem)
		assert.Equal(t, 1, tree.Root.Left.Left.Elem)
		assert.Equal(t, 10, tree.Root.Right.Elem)
		assert.Equal(t, 8, tree.Root.Right.Left.Elem)
		assert.Equal(t, 15, tree.Root.Right.Right.Elem)
	})

	t.Run("Binary search tree properties should be maintained with right rotation", func(t *testing.T) {
		tree := ds.NewAVLTree[int]()

		tree.Insert(10)
		tree.Insert(5)
		tree.Insert(15)
		tree.Insert(11)
		tree.Insert(18)

		assert.Equal(t, 10, tree.Root.Elem)
		assert.Equal(t, 5, tree.Root.Left.Elem)
		assert.Equal(t, 15, tree.Root.Right.Elem)
		assert.Equal(t, 11, tree.Root.Right.Left.Elem)
		assert.Equal(t, 18, tree.Root.Right.Right.Elem)

		err := tree.Insert(20)

		assert.Nil(t, err)
		assert.Equal(t, 15, tree.Root.Elem)
		assert.Equal(t, 10, tree.Root.Left.Elem)
		assert.Equal(t, 5, tree.Root.Left.Left.Elem)
		assert.Equal(t, 11, tree.Root.Left.Right.Elem)
		assert.Equal(t, 18, tree.Root.Right.Elem)
		assert.Equal(t, 20, tree.Root.Right.Right.Elem)
	})
}

func TestAvlTreeRemove(t *testing.T) {
	t.Run("remove an element from an empty tree", func(t *testing.T) {
		tree := ds.NewAVLTree[int]()
		err := tree.Remove(10)
		assert.NotNilf(t, err, "Remove should return an error if the element is not found")
		assert.Equalf(t, 0, tree.Size(), "AVLTree size should not change if item is not removed")
	})

	t.Run("remove an element from a tree of size 1", func(t *testing.T) {
		tree := ds.NewAVLTree[int]()
		assert.Equalf(t, 0, tree.Size(), "Newly intialized AVLTree should have size of 0")
		tree.Insert(10)
		assert.Equalf(t, 1, tree.Size(), "AVLTree size should be incremented after insert")
		err := tree.Remove(10)
		assert.Nilf(t, err, "Remove should not return an error if the element is  found")
		assert.Equalf(t, 0, tree.Size(), "AVLTree size should be decremented after remove")
	})

	t.Run("remove a leaf node", func(t *testing.T) {
		tree := ds.NewAVLTree[int]()

		tree.Insert(10)
		tree.Insert(5)
		tree.Insert(15)
		tree.Insert(7)
		tree.Insert(3)
		tree.Insert(24)
		tree.Insert(13)

		assert.Equal(t, 7, tree.Size())

		err := tree.Remove(7)

		if err != nil {
			t.Errorf("Removal of leaf node should not have resulted in error")
		}
		assert.Equal(t, 6, tree.Size())
		assert.Nilf(t, tree.Root.Left.Right, "Node should be removed")

		err = tree.Remove(13)
		if err != nil {
			t.Errorf("Removal of leaf node should not have resulted in error")
		}
		assert.Equal(t, 5, tree.Size())
		assert.Nilf(t, tree.Root.Right.Left, "Node should be removed")
	})

	t.Run("remove a node with only a left child", func(t *testing.T) {
		tree := ds.NewAVLTree[int]()

		tree.Insert(10)
		tree.Insert(5)
		tree.Insert(15)
		tree.Insert(3)

		assert.Equal(t, 4, tree.Size())

		err := tree.Remove(5)

		if err != nil {
			t.Errorf("Removal of node should not have resulted in error")
		}
		assert.Equal(t, 3, tree.Size())
		assert.Equal(t, 3, tree.Root.Left.Elem, "Left child should be promoted")
	})

	t.Run("remove a node with only a right child", func(t *testing.T) {
		tree := ds.NewAVLTree[int]()

		tree.Insert(10)
		tree.Insert(5)
		tree.Insert(15)
		tree.Insert(7)

		assert.Equal(t, 4, tree.Size())

		err := tree.Remove(5)

		if err != nil {
			t.Errorf("Removal of node should not have resulted in error")
		}
		assert.Equal(t, 3, tree.Size())
		assert.Equal(t, 7, tree.Root.Left.Elem, "Left child should be promoted")
	})

	t.Run("remove a node with both left and right children", func(t *testing.T) {
		tree := ds.NewAVLTree[int]()

		tree.Insert(10)
		tree.Insert(5)
		tree.Insert(15)
		tree.Insert(7)
		tree.Insert(3)

		assert.Equal(t, 5, tree.Size())

		err := tree.Remove(5)

		if err != nil {
			t.Errorf("Removal of node should not have resulted in error")
		}
		assert.Equal(t, 4, tree.Size())
		assert.Equal(t, 7, tree.Root.Left.Elem, "Left child should be promoted")
		assert.Equal(t, 3, tree.Root.Left.Left.Elem, "Left child of removed elem should be promoted node's left child")
	})

	t.Run("the AVL properties should be maintained during a delete - left rotation", func(t *testing.T) {
		tree := ds.NewAVLTree[int]()

		tree.Insert(10)
		tree.Insert(5)
		tree.Insert(15)
		tree.Insert(7)
		tree.Insert(3)
		tree.Insert(24)
		tree.Insert(13)
		tree.Insert(30)

		assert.Equal(t, 8, tree.Size())

		assert.Equal(t, 10, tree.Root.Elem)
		assert.Equal(t, 5, tree.Root.Left.Elem)
		assert.Equal(t, 3, tree.Root.Left.Left.Elem)
		assert.Equal(t, 7, tree.Root.Left.Right.Elem)
		assert.Equal(t, 15, tree.Root.Right.Elem)
		assert.Equal(t, 13, tree.Root.Right.Left.Elem)
		assert.Equal(t, 24, tree.Root.Right.Right.Elem)
		assert.Equal(t, 30, tree.Root.Right.Right.Right.Elem)

		err := tree.Remove(5)
		if err != nil {
			t.Errorf("Removal of node(5) should not have resulted in error")
		}
		err = tree.Remove(7)
		if err != nil {
			t.Errorf("Removal of node(7) should not have resulted in error")
		}

		assert.Equal(t, 15, tree.Root.Elem)
		assert.Equal(t, 10, tree.Root.Left.Elem)
		assert.Equal(t, 3, tree.Root.Left.Left.Elem)
		assert.Equal(t, 13, tree.Root.Left.Right.Elem)
		assert.Equal(t, 24, tree.Root.Right.Elem)
		assert.Equal(t, 30, tree.Root.Right.Right.Elem)
	})
}
