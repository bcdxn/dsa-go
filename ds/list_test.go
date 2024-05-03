package ds_test

import (
	"testing"

	"github.com/bcdxn/dsa-go/ds"
	"github.com/stretchr/testify/assert"
)

func TestListInitializer(t *testing.T) {
	t.Run("Initializer should create an empty list", func(t *testing.T) {
		ll := ds.NewList[int]()
		assert.Equal(t, 0, ll.Size())
		assert.Nil(t, ll.Head, "empty list Head should be nil")
		assert.Nil(t, ll.Tail, "empty list Tail should be nil")
	})
}

func TestListAddHead(t *testing.T) {
	t.Run("Adding to an empty list", func(t *testing.T) {
		ll := ds.NewList[int]()
		assert.Equal(t, 0, ll.Size())
		ll.AddHead(10)
		assert.Equal(t, 1, ll.Size(), "Adding 1 element should increase the size by 1")
		ll.AddHead(10)
		ll.AddHead(10)
		ll.AddHead(10)
		assert.Equal(t, 4, ll.Size(), "Adding 3 element should increase the size by 3")
	})

	t.Run("Adding to a non-empty list", func(t *testing.T) {
		ll := ds.NewList[int]()
		assert.Equal(t, 0, ll.Size())
		ll.AddHead(10)
		ll.AddHead(5)
		ll.AddHead(8)
		ll.AddHead(19)
		assert.Equal(t, 4, ll.Size(), "Adding 4 element should increase the size by 4")
	})

	t.Run("Should be able to traverse to all of the added nodes from head", func(t *testing.T) {
		ll := ds.NewList[int]()
		elems := [4]int{10, 5, 8, 19}

		for _, elem := range elems {
			ll.AddHead(elem)
		}

		node := ll.Head
		i := len(elems) - 1

		for node != nil {
			assert.Equal(t, elems[i], node.Elem, "Element inserted from array should be in reverse order in list")
			node = node.Next
			i--
		}
	})

	t.Run("Should be able to traverse to all of the added nodes from tail", func(t *testing.T) {
		ll := ds.NewList[int]()
		elems := [4]int{10, 5, 8, 19}

		for _, elem := range elems {
			ll.AddHead(elem)
		}

		node := ll.Tail
		i := 0

		for node != nil {
			assert.Equal(t, elems[i], node.Elem, "Element inserted from array should be in reverse order in list")
			node = node.Prev
			i++
		}
	})
}

func TestListAddTail(t *testing.T) {
	t.Run("Adding to an empty list", func(t *testing.T) {
		ll := ds.NewList[int]()
		assert.Equal(t, 0, ll.Size())
		ll.AddTail(10)
		assert.Equal(t, 1, ll.Size(), "Adding 1 element should increase the size by 1")
		ll.AddTail(10)
		ll.AddTail(10)
		ll.AddTail(10)
		assert.Equal(t, 4, ll.Size(), "Adding 3 element should increase the size by 3")
	})

	t.Run("Adding to a non-empty list", func(t *testing.T) {
		ll := ds.NewList[int]()
		assert.Equal(t, 0, ll.Size())
		ll.AddTail(10)
		ll.AddTail(5)
		ll.AddTail(8)
		ll.AddTail(19)
		assert.Equal(t, 4, ll.Size(), "Adding 4 element should increase the size by 4")
	})

	t.Run("Should be able to traverse to all of the added nodes from head", func(t *testing.T) {
		ll := ds.NewList[int]()
		elems := [4]int{10, 5, 8, 19}

		for _, elem := range elems {
			ll.AddTail(elem)
		}

		node := ll.Head
		i := 0

		for node != nil {
			assert.Equal(t, elems[i], node.Elem, "Element inserted from array should be in reverse order in list")
			node = node.Next
			i++
		}
	})

	t.Run("Should be able to traverse to all of the added nodes from tail", func(t *testing.T) {
		ll := ds.NewList[int]()
		elems := [4]int{10, 5, 8, 19}

		for _, elem := range elems {
			ll.AddTail(elem)
		}

		node := ll.Tail
		i := len(elems) - 1

		for node != nil {
			assert.Equal(t, elems[i], node.Elem, "Element inserted from array should be in reverse order in list")
			node = node.Prev
			i--
		}
	})
}

func TestListRemove(t *testing.T) {
	t.Run("Removing from an empty list", func(t *testing.T) {
		ll := ds.NewList[int]()
		_, err := ll.Remove(10)

		assert.NotNil(t, err, "An error should be returned if the value is not found in the list")
	})

	t.Run("Removing an element that does not exist in the list", func(t *testing.T) {
		ll := ds.NewList[int]()
		elems := [4]int{10, 5, 8, 19}

		for _, elem := range elems {
			ll.AddHead(elem)
		}
		_, err := ll.Remove(15)

		assert.NotNil(t, err, "An error should be returned if the value is not found in the list")
	})

	t.Run("Removing from a list of length 1", func(t *testing.T) {
		ll := ds.NewList[int]()
		ll.AddHead(10)
		elem, err := ll.Remove(10)

		if err != nil {
			t.Errorf("Should not have errored - %s", err)
		}

		assert.Equal(t, elem, 10, "Should return the removed element")
		assert.Equal(t, ll.Size(), 0, "Size should be decremented")
		assert.Nil(t, ll.Head, "The list head should be nil")
		assert.Nil(t, ll.Tail, "The list tail should be nil")
	})

	t.Run("Removing an element from the middle of the list", func(t *testing.T) {
		ll := ds.NewList[int]()
		elems := [5]int{10, 5, 8, 19, 16}
		expectedList := [4]int{16, 19, 5, 10}

		for _, elem := range elems {
			ll.AddHead(elem)
		}

		assert.Equal(t, ll.Size(), 5)
		ll.Remove(8)
		assert.Equal(t, ll.Size(), 4, "Size should be decremented")

		node := ll.Head
		i := 0

		for node != nil {
			assert.Equal(t, expectedList[i], node.Elem, "List should still be traversable from head to tail")
			node = node.Next
			i++
		}

		i = 0
		for node != nil {
			assert.Equal(t, elems[i], node.Elem, "List should still be traversable from tail to head")
			node = node.Prev
			i++
		}
	})

	t.Run("Removing an element from the head of the list", func(t *testing.T) {
		ll := ds.NewList[int]()
		elems := [5]int{10, 5, 8, 19, 16}
		expectedList := [4]int{19, 8, 5, 10}

		for _, elem := range elems {
			ll.AddHead(elem)
		}

		assert.Equal(t, ll.Size(), 5)
		ll.Remove(16)
		assert.Equal(t, ll.Size(), 4, "Size should be decremented")

		node := ll.Head
		i := 0

		for node != nil {
			assert.Equal(t, expectedList[i], node.Elem, "List should still be traversable from head to tail")
			node = node.Next
			i++
		}

		i = 0
		for node != nil {
			assert.Equal(t, elems[i], node.Elem, "List should still be traversable from tail to head")
			node = node.Prev
			i++
		}
	})

	t.Run("Removing an element from the tail of the list", func(t *testing.T) {
		ll := ds.NewList[int]()
		elems := [5]int{10, 5, 8, 19, 16}
		expectedList := [4]int{16, 19, 8, 5}

		for _, elem := range elems {
			ll.AddHead(elem)
		}

		assert.Equal(t, ll.Size(), 5)
		ll.Remove(10)
		assert.Equal(t, ll.Size(), 4, "Size should be decremented")

		node := ll.Head
		i := 0

		for node != nil {
			assert.Equal(t, expectedList[i], node.Elem, "List should still be traversable from head to tail")
			node = node.Next
			i++
		}

		i = 0
		for node != nil {
			assert.Equal(t, elems[i], node.Elem, "List should still be traversable from tail to head")
			node = node.Prev
			i++
		}
	})
}

func TestListContains(t *testing.T) {
	t.Run("Check contains on an empty list", func(t *testing.T) {
		l := ds.NewList[int]()

		assert.Falsef(t, l.Contains(5), "An empty list should not contain elements")
	})

	t.Run("Check contains on a list with a single item", func(t *testing.T) {
		l := ds.NewList[int]()

		l.AddHead(10)

		assert.Falsef(t, l.Contains(5), "The element should not be contained in the list")
		assert.Truef(t, l.Contains(10), "The element should be contained in the list")
	})

	t.Run("Check contains on a list with many items", func(t *testing.T) {
		l := ds.NewList[int]()

		l.AddHead(10)
		l.AddHead(15)
		l.AddHead(20)
		l.AddHead(30)

		assert.Falsef(t, l.Contains(5), "The element should not be contained in the list")
		assert.Truef(t, l.Contains(10), "The element should be contained in the list")
		assert.Truef(t, l.Contains(15), "The element should be contained in the list")
		assert.Truef(t, l.Contains(20), "The element should be contained in the list")
		assert.Truef(t, l.Contains(30), "The element should be contained in the list")
	})
}

func TestListReverse(t *testing.T) {
	t.Run("testing the test", func(t *testing.T) {
		ll := ds.NewList[int]()
		elems := [5]int{10, 5, 8, 19, 16}

		for _, elem := range elems {
			ll.AddHead(elem)
		}

		ll.Reverse()

		node := ll.Head
		i := 0

		for node != nil {
			assert.Equal(t, elems[i], node.Elem, "List should be reversed (treaversable forwards)")
			node = node.Next
			i++
		}
		assert.Equal(t, i, len(elems))

		i = len(elems) - 1

		for node != nil {
			assert.Equal(t, elems[i], node.Elem, "List should be reversed (treaversable backwards)")
			node = node.Prev
			i--
		}
	})
}
