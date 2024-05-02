package ds_test

import (
	"testing"

	"github.com/bcdxn/dsa-go/ds"
	"github.com/stretchr/testify/assert"
)

func TestLinkedListInitializer(t *testing.T) {
	t.Run("Initializer should create an empty list", func(t *testing.T) {
		ll := ds.NewLinkedList[int]()
		assert.Equal(t, 0, ll.Size())
	})
}

func TestLinkedListAdd(t *testing.T) {
	t.Run("Adding to an empty list", func(t *testing.T) {
		ll := ds.NewLinkedList[int]()
		assert.Equal(t, 0, ll.Size())
		ll.Add(10)
		assert.Equal(t, 1, ll.Size(), "Adding 1 element should increase the size by 1")
		ll.Add(10)
		ll.Add(10)
		ll.Add(10)
		assert.Equal(t, 4, ll.Size(), "Adding 3 element should increase the size by 3")
	})

	t.Run("Adding to a non-empty list", func(t *testing.T) {
		ll := ds.NewLinkedList[int]()
		assert.Equal(t, 0, ll.Size())
		ll.Add(10)
		ll.Add(5)
		ll.Add(8)
		ll.Add(19)
		assert.Equal(t, 4, ll.Size(), "Adding 4 element should increase the size by 4")
	})

	t.Run("Should be able to traverse to all of the added nodes", func(t *testing.T) {
		ll := ds.NewLinkedList[int]()
		elems := [4]int{10, 5, 8, 19}

		for _, elem := range elems {
			ll.Add(elem)
		}

		node := ll.Head
		i := len(elems) - 1

		for node != nil {
			assert.Equal(t, elems[i], node.Elem, "Element inserted from array should be in reverse order in list")
			node = node.Next
			i--
		}
	})
}

func TestLinkedListRemove(t *testing.T) {
	t.Run("Removing from an empty list", func(t *testing.T) {
		ll := ds.NewLinkedList[int]()
		_, err := ll.Remove(10)

		assert.NotNil(t, err, "An error should be returned if the value is not found in the list")
	})

	t.Run("Removing an element that does not exist in the list", func(t *testing.T) {
		ll := ds.NewLinkedList[int]()
		elems := [4]int{10, 5, 8, 19}

		for _, elem := range elems {
			ll.Add(elem)
		}
		_, err := ll.Remove(15)

		assert.NotNil(t, err, "An error should be returned if the value is not found in the list")
	})

	t.Run("Removing from a list of length 1", func(t *testing.T) {
		ll := ds.NewLinkedList[int]()
		ll.Add(10)
		elem, err := ll.Remove(10)

		if err != nil {
			t.Errorf("Should not have errored - %s", err)
		}

		assert.Equal(t, elem, 10, "Should return the removed element")
		assert.Equal(t, ll.Size(), 0, "Size should be decremented")
		assert.Nil(t, ll.Head, "The list should now be empty")
	})

	t.Run("Removing an element from the middle of the list", func(t *testing.T) {
		ll := ds.NewLinkedList[int]()
		elems := [5]int{10, 5, 8, 19, 16}
		expectedList := [4]int{16, 19, 5, 10}

		for _, elem := range elems {
			ll.Add(elem)
		}

		assert.Equal(t, ll.Size(), 5)
		ll.Remove(8)
		assert.Equal(t, ll.Size(), 4, "Size should be decremented")

		node := ll.Head
		i := 0

		for node != nil {
			assert.Equal(t, expectedList[i], node.Elem, "List should still be traversable")
			node = node.Next
			i++
		}
	})
}

func TestLinkedListReverse(t *testing.T) {
	t.Run("testing the test", func(t *testing.T) {
		ll := ds.NewLinkedList[int]()
		elems := [5]int{10, 5, 8, 19, 16}

		for _, elem := range elems {
			ll.Add(elem)
		}

		ll.Reverse()

		node := ll.Head
		i := 0

		for node != nil {
			assert.Equal(t, elems[i], node.Elem, "List should be reversed")
			node = node.Next
			i++
		}

		assert.Equal(t, i, len(elems))
	})
}

func TestLinkedListReverseRecursive(t *testing.T) {
	t.Run("testing the test", func(t *testing.T) {
		ll := ds.NewLinkedList[int]()
		elems := [5]int{10, 5, 8, 19, 16}

		for _, elem := range elems {
			ll.Add(elem)
		}

		ll.ReverseRecursive()

		node := ll.Head
		i := 0

		for node != nil {
			assert.Equal(t, elems[i], node.Elem, "List should be reversed")
			node = node.Next
			i++
		}

		assert.Equal(t, i, len(elems))
	})
}
