package sort

import "math/rand"

// Merge returns a sorted copy of the given list using the [Merge Sort algorithm][0]. This
// function does not alter the given list.
//
// **Time Complexity**
//
// ```
// | Scenario     | Big-O       |
// |:-------------|:------------|
// | Average Case | O(n*log(n)) |
// | Best Case    | O(n*log(n)) |
// | Worst Case   | O(n^2)      |
// ```
//
// Additional Notes:
// - Space complexity - O(log(n))
//
// [0]: https://en.wikipedia.org/wiki/Quicksort
func Quick(list []int) []int {
	// Copy the list to ensure there are no side effects on the given list
	cpy := make([]int, len(list))
	copy(cpy, list)

	quickHelper(cpy, 0, len(cpy)-1)

	return cpy
}

// parition is a private helper function to create a semi-sorted sort list with relation to a rand
// pivot value. All values to the left of the pivot are less than the pivot and all values to the
// right of the pivot are greater than or equal to the pivot.
func partition(list []int, low int, high int) int {
	// Pick a random pivot
	pivotIndex := rand.Intn(high-low+1) + low
	pivotValue := list[pivotIndex]
	// move the pivot value to the end of the array
	if pivotIndex != high {
		list[pivotIndex], list[high] = list[high], list[pivotIndex]
	}
	// keep track of the first element that is larger than the pivot value
	j := low
	// order the list comparing to the pivot
	for i := low; i < high; i++ {
		if list[i] < pivotValue {
			list[i], list[j] = list[j], list[i]
			j++
		}
	}
	// at this point everything to the left of 'j' is less than the pivot and everything to the right
	// of 'j' is greater than the pivot.

	// swap j and the pivot value (high index) to complete the fully semi-sorted partition
	list[j], list[high] = list[high], list[j]
	// return the pivot so we can recursively quick sort to the left and right of the pivot
	return j
}

// quickHelper is a recursive helper function to implement Quick Sort.
func quickHelper(list []int, low int, high int) {
	// base case
	if high <= low {
		return
	}
	// partition
	pivot := partition(list, low, high)
	// recursive case
	quickHelper(list, low, pivot-1)
	quickHelper(list, pivot+1, high)
}
