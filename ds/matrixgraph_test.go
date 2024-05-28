package ds_test

import (
	"testing"

	"github.com/bcdxn/dsa-go/ds"
	"github.com/stretchr/testify/assert"
)

func TestMatrixGraphStringify(t *testing.T) {
	t.Run("Should output a string representation of the graph", func(t *testing.T) {
		g := ds.NewMatrixGraph([][]bool{
			{true, true, true, true},
			{false, false, true, true},
			{true, true, true, false},
			{true, false, true, true},
		})

		s := g.Stringify()

		expected := "1 1 1 1 \n0 0 1 1 \n1 1 1 0 \n1 0 1 1 \n"

		assert.Equal(t, expected, s)
	})
}

func TestMatrixGraphNumPaths(t *testing.T) {
	t.Run("Should return the count of all paths from the start node", func(t *testing.T) {
		g := ds.NewMatrixGraph([][]bool{
			{true, true, true, true},
			{false, false, true, true},
			{true, true, true, false},
			{true, false, true, true},
		})

		paths := g.NumPaths(0, 0, 3, 3)
		assert.Equal(t, 2, paths)

		paths = g.NumPaths(0, 2, 3, 3)
		assert.Equal(t, 2, paths)

		paths = g.NumPaths(1, 3, 3, 3)
		assert.Equal(t, 2, paths)

		paths = g.NumPaths(1, 2, 3, 3)
		assert.Equal(t, 1, paths)
	})

	t.Run("Should return 0 if the start node is outside of the matrix boundaries", func(t *testing.T) {
		g := ds.NewMatrixGraph([][]bool{
			{true, true, true, true},
			{false, false, true, true},
			{true, true, true, false},
			{true, false, true, true},
		})

		paths := g.NumPaths(-1, -1, 3, 3)
		assert.Equal(t, 0, paths)
	})

	t.Run("Should return 0 if the destination node is outside of the matrix boundaries", func(t *testing.T) {
		g := ds.NewMatrixGraph([][]bool{
			{true, true, true, true},
			{false, false, true, true},
			{true, true, true, false},
			{true, false, true, true},
		})

		paths := g.NumPaths(0, 0, 10, 10)
		assert.Equal(t, 0, paths)
	})

	t.Run("Should return 0 the start node is non traversable", func(t *testing.T) {
		g := ds.NewMatrixGraph([][]bool{
			{true, true, true, true},
			{false, false, true, true},
			{true, true, true, false},
			{true, false, true, true},
		})

		paths := g.NumPaths(1, 1, 10, 10)
		assert.Equal(t, 0, paths)
	})
}

func TestMatrixGraphNumIslands(t *testing.T) {
	t.Run("should count the correct number of islands", func(t *testing.T) {
		g := ds.NewMatrixGraph([][]bool{
			{false, true, true, true, false},
			{false, true, false, true, false},
			{true, true, false, false, false},
			{false, false, false, false, false},
		})

		c := g.NumIslands()

		assert.Equal(t, 1, c)
	})

	t.Run("should count the correct number of islands", func(t *testing.T) {
		g := ds.NewMatrixGraph([][]bool{
			{true, true, false, false, true},
			{true, true, false, false, true},
			{false, false, true, false, false},
			{false, false, false, true, true},
		})

		c := g.NumIslands()

		assert.Equal(t, 4, c)
	})
}

func TestMatrixGraphShortestPath(t *testing.T) {
	t.Run("should return the shortest path length from start to finish", func(t *testing.T) {
		g := ds.NewMatrixGraph([][]bool{
			{true, true, true, true},
			{false, false, true, true},
			{true, true, true, false},
			{true, false, true, true},
		})

		l, err := g.ShortestPath(0, 0, 3, 3)

		assert.Nil(t, err)
		assert.Equal(t, 6, l)
	})

	t.Run("should return an error if thre is no valid path from start to destination", func(t *testing.T) {
		g := ds.NewMatrixGraph([][]bool{
			{true, true, false, false},
			{false, false, false, false},
			{true, true, true, false},
			{true, false, true, true},
		})

		_, err := g.ShortestPath(0, 0, 3, 3)

		assert.NotNil(t, err)
	})
}
