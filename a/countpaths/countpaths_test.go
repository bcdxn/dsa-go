package countpaths_test

import (
	"testing"

	"github.com/bcdxn/dsa-go/a/countpaths"
	"github.com/stretchr/testify/assert"
)

func TestCountPathsBruteForce(t *testing.T) {
	t.Run("should return the number of paths", func(t *testing.T) {
		grid := [][]int{
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
		}

		assert.Equal(t, 2, countpaths.BruteForce(2, 2, 3, 3, grid))
		assert.Equal(t, 6, countpaths.BruteForce(1, 1, 3, 3, grid))
		assert.Equal(t, 20, countpaths.BruteForce(0, 0, 3, 3, grid))
	})
}

func TestCountPathsTopDown(t *testing.T) {
	t.Run("should return the number of paths", func(t *testing.T) {
		grid := [][]int{
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
		}

		assert.Equal(t, 2, countpaths.TopDown(2, 2, 3, 3, grid))
		assert.Equal(t, 6, countpaths.TopDown(1, 1, 3, 3, grid))
		assert.Equal(t, 20, countpaths.TopDown(0, 0, 3, 3, grid))
	})
}

func TestCountPathsBottomUp(t *testing.T) {
	t.Run("should return the number of paths", func(t *testing.T) {
		grid := [][]int{
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
		}

		assert.Equal(t, 2, countpaths.BottomUp(2, 2, 3, 3, grid))
		assert.Equal(t, 6, countpaths.BottomUp(1, 1, 3, 3, grid))
		assert.Equal(t, 20, countpaths.BottomUp(0, 0, 3, 3, grid))
	})
}
