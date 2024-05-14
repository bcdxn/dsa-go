package search_test

import (
	"testing"

	"github.com/bcdxn/dsa-go/search"
	"github.com/stretchr/testify/assert"
)

func TestSearchBinary(t *testing.T) {
	t.Run("Search an empty list", func(t *testing.T) {
		l := make([]int, 0)

		f, e := search.Binary(l, 5)

		assert.Zero(t, f)
		assert.False(t, e)
	})

	t.Run("Search a list of size 1", func(t *testing.T) {
		l := []int{1}

		f, e := search.Binary(l, 5)
		assert.Zero(t, f)
		assert.False(t, e)

		f, e = search.Binary(l, 1)
		assert.Equal(t, f, 1)
		assert.True(t, e)
	})

	t.Run("Search a list of size 10", func(t *testing.T) {
		l := []int{1, 5, 30, 31, 32, 30, 45, 80, 87, 100}

		f, e := search.Binary(l, -1)
		assert.Zero(t, f)
		assert.False(t, e)
		// Found with first guess
		f, e = search.Binary(l, 32)
		assert.Equal(t, f, 32)
		assert.True(t, e)
		// search left sub list
		f, e = search.Binary(l, 5)
		assert.Equal(t, f, 5)
		assert.True(t, e)
		// search left sub list
		f, e = search.Binary(l, 1)
		assert.Equal(t, f, 1)
		assert.True(t, e)
		// search right sub list
		f, e = search.Binary(l, 87)
		assert.Equal(t, f, 87)
		assert.True(t, e)
		// search right sub list
		f, e = search.Binary(l, 100)
		assert.Equal(t, f, 100)
		assert.True(t, e)
	})
}
