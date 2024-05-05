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
