package fibonacci_test

import (
	"testing"

	"github.com/bcdxn/dsa-go/a/fibonacci"
	"github.com/stretchr/testify/assert"
)

func TestFibonacciBruteForce(t *testing.T) {
	t.Run("should calculate base case 0", func(t *testing.T) {
		assert.Equal(t, 0, fibonacci.BruteForce(0))
	})

	t.Run("should calculate base case 1", func(t *testing.T) {
		assert.Equal(t, 1, fibonacci.BruteForce(1))
	})

	t.Run("should calculate fibonacci number correctly", func(t *testing.T) {
		assert.Equal(t, 5, fibonacci.BruteForce(5))
		assert.Equal(t, 13, fibonacci.BruteForce(7))
		assert.Equal(t, 55, fibonacci.BruteForce(10))
	})
}

func TestFibonacciMemoization(t *testing.T) {
	t.Run("should calculate base case 0", func(t *testing.T) {
		assert.Equal(t, 0, fibonacci.Memoization(0))
	})

	t.Run("should calculate base case 1", func(t *testing.T) {
		assert.Equal(t, 1, fibonacci.Memoization(1))
	})

	t.Run("should calculate fibonacci number correctly", func(t *testing.T) {
		assert.Equal(t, 5, fibonacci.Memoization(5))
		assert.Equal(t, 13, fibonacci.Memoization(7))
		assert.Equal(t, 55, fibonacci.Memoization(10))
	})
}

func TestFibonacciBottomUp(t *testing.T) {
	t.Run("should calculate base case 0", func(t *testing.T) {
		assert.Equal(t, 0, fibonacci.BottomUp(0))
	})

	t.Run("should calculate base case 1", func(t *testing.T) {
		assert.Equal(t, 1, fibonacci.BottomUp(1))
	})

	t.Run("should calculate fibonacci number correctly", func(t *testing.T) {
		assert.Equal(t, 5, fibonacci.BottomUp(5))
		assert.Equal(t, 13, fibonacci.BottomUp(7))
		assert.Equal(t, 55, fibonacci.BottomUp(10))
	})
}
