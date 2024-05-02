package sort_test

import (
	"slices"
	"testing"

	"github.com/bcdxn/dsa-go/sort"
)

// Test suite 'borrowed' from - https://github.com/TheAlgorithms/Go/blob/master/sort/sorts_test.go

type SortTest struct {
	input       []int
	expected    []int
	description string
	willChange  bool
}

// A function that runs through a set of common sorting-oriented tests for a given sort function
func runTests(t *testing.T, sortFunction func([]int) []int, inPlace bool) {
	tests := []SortTest{
		//Sorted slice
		{
			input:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			description: "Sorted Unsigned",
			willChange:  false,
		},
		//Reversed slice
		{
			input:       []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			description: "Reversed Unsigned",
			willChange:  true,
		},
		//Sorted slice
		{
			input:       []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected:    []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			description: "Sorted Signed",
			willChange:  false,
		},
		//Reversed slice
		{
			input:       []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
			expected:    []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			description: "Reversed Signed",
			willChange:  true,
		},
		//Reversed slice, even length
		{
			input:       []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
			expected:    []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			description: "Reversed Signed #2",
			willChange:  true,
		},
		//Random order with repetitions
		{
			input:       []int{-5, 7, 4, -2, 6, 5, 8, 3, 2, -7, -1, 0, -3, 9, -6, -4, 10, 9, 1, -8, -9, -10},
			expected:    []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 10},
			description: "Random order Signed",
			willChange:  true,
		},
		//Single-entry slice
		{
			input:       []int{1},
			expected:    []int{1},
			description: "Singleton",
			willChange:  false,
		},
		// Empty slice
		{
			input:       []int{},
			expected:    []int{},
			description: "Empty Slice",
			willChange:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			actual := sortFunction(test.input)
			isSorted := false

			if inPlace {
				isSorted = slices.Equal(test.input, test.expected)
				if !slices.Equal(test.input, test.expected) {
					t.Errorf("test '%s' - FAILED (due to missing side effects)\n", test.description)
				}
			} else {
				isSorted = slices.Equal(actual, test.expected)
				if test.willChange && slices.Equal(test.input, test.expected) {
					t.Errorf("test '%s' - FAILED (due to unwanted side effects)\n", test.description)
				}
			}

			if !isSorted {
				t.Errorf("test '%s' - FAILED (due to incorrect sort order)\n", test.description)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	runTests(t, sort.Insertion, false)
}

func TestInPlaceInsertionSort(t *testing.T) {
	runTests(t, sort.InPlaceInsertion, true)
}

func TestBubbleSort(t *testing.T) {
	runTests(t, sort.Bubble, false)
}

func TestInPlaceBubbleSort(t *testing.T) {
	runTests(t, sort.InPlaceBubble, true)
}

func TestMergeSort(t *testing.T) {
	runTests(t, sort.Merge, false)
}

func TestQuickSort(t *testing.T) {
	runTests(t, sort.Quick, false)
}
