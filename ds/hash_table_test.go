package ds_test

import (
	"testing"

	"github.com/bcdxn/dsa-go/ds"
	"github.com/stretchr/testify/assert"
)

func TestDbj2Hash(t *testing.T) {
	t.Run("Should allow hashing of strings", func(t *testing.T) {
		hash := ds.Dbj2Hash("Hello")
		assert.Equal(t, uint(210676686969), hash)
	})

	t.Run("Should return the same hash for the same string", func(t *testing.T) {
		hash := ds.Dbj2Hash("Hello")
		hash2 := ds.Dbj2Hash("Hello")
		assert.Equal(t, hash, hash2)
	})

	t.Run("Should return different hashes for different strings", func(t *testing.T) {
		hash := ds.Dbj2Hash("Hello!")
		assert.Equal(t, uint(6952330670010), hash)
	})
}

func TestNewHashTable(t *testing.T) {
	t.Run("Should create an empty hash table", func(t *testing.T) {
		h := ds.NewHashTable()
		assert.Equal(t, 0, h.Len())
	})
}

func TestHashTableSet(t *testing.T) {
	t.Run("Should grow in size when an element is added", func(t *testing.T) {
		h := ds.NewHashTable()
		assert.Equal(t, 0, h.Len())

		h.Add("a test")
		assert.Equal(t, 1, h.Len())
	})

	t.Run("Adding elements with the same hash value should grow the size", func(t *testing.T) {
		h := ds.NewHashTable()
		assert.Equal(t, 0, h.Len())

		h.Add("a test")
		h.Add("a test")
		h.Add("a test")
		assert.Equal(t, 3, h.Len())
	})

	t.Run("Adding elements with differe hash values should grow the size", func(t *testing.T) {
		h := ds.NewHashTable()
		assert.Equal(t, 0, h.Len())

		h.Add("a test")
		h.Add("a test again")
		h.Add("a test and another")
		assert.Equal(t, 3, h.Len())
	})
}

func TestHashTableGet(t *testing.T) {
	t.Run("Get on an empty table", func(t *testing.T) {
		h := ds.NewHashTable()
		hash := ds.Dbj2Hash("Test")
		s, err := h.Get(hash)

		assert.Zero(t, s)
		assert.NotNil(t, err)
	})

	t.Run("Get non existent hash on an empty table", func(t *testing.T) {
		h := ds.NewHashTable()
		h.Add("a test element")
		h.Add("a test element or two")
		h.Add("a test element or three")
		hash := ds.Dbj2Hash("Test")
		s, err := h.Get(hash)

		assert.Zero(t, s)
		assert.NotNil(t, err)
	})

	t.Run("Get existing elements from table", func(t *testing.T) {
		h := ds.NewHashTable()
		h.Add("a test element")
		h.Add("a test element or two")
		h.Add("a test element or three")

		hash1 := ds.Dbj2Hash("a test element")
		hash2 := ds.Dbj2Hash("a test element or two")
		hash3 := ds.Dbj2Hash("a test element or three")
		s1, err := h.Get(hash1)
		assert.Equal(t, s1, "a test element")
		assert.Nil(t, err)
		s2, err := h.Get(hash2)
		assert.Equal(t, s2, "a test element or two")
		assert.Nil(t, err)
		s3, err := h.Get(hash3)
		assert.Equal(t, s3, "a test element or three")
		assert.Nil(t, err)
	})
}
