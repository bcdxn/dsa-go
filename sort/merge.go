package sort

// Merge returns a sorted copy of the given list using the [Merge Sort algorithm][0]. This
// function does not alter the given list.
//
// **Time Complexity**
//
// ```
// | Scenario     | Big-O  |
// |:-------------|:-------|
// | Average Case | O(n*log(n))  |
// | Best Case    | O(n*log(n))  |
// | Worst Case   | O(n*log(n))  |
// ```
//
// Additional Notes:
// - Space complexity - O(n)
//
// [0]: https://en.wikipedia.org/wiki/Merge_sort
func Merge(list []int) []int {
	// Copy the list to ensure there are no side effects on the given list
	cpy := make([]int, len(list))
	copy(cpy, list)

	return mergeHelper(cpy)
}

// InPlaceMerge sorts the given list using the [Merge Sort algorithm][0]. This function alters the
// given list.
//
// **Time Complexity**
//
// ```
// | Scenario     | Big-O  |
// |:-------------|:-------|
// | Average Case | O(n*log(n))  |
// | Best Case    | O(n*log(n))  |
// | Worst Case   | O(n*log(n))  |
// ```
//
// Additional Notes:
// - Space complexity - O(n)
//
// [0]: https://en.wikipedia.org/wiki/Merge_sort
func InPlaceMerge(list []int) []int {
	return mergeHelper(list)
}

func mergeHelper(list []int) []int {
	// Base Case
	// The list has size of 1 (or is empty) it is emplicitly sorted
	if len(list) < 2 {
		return list
	}

	midpoint := len(list) / 2

	// Recursive Case
	sortedLeft := mergeHelper(list[0:midpoint])
	sortedRight := mergeHelper(list[midpoint:])

	// Return the merged sorted lists
	return mergeSortedLists(sortedLeft, sortedRight)
}

func mergeSortedLists(leftList []int, rightList []int) []int {
	sorted := make([]int, len(leftList)+len(rightList))

	i := 0
	leftIndex := 0
	rightIndex := 0

	// Iteratively insert the smallest element from left and right sorted lists
	for leftIndex < len(leftList) && rightIndex < len(rightList) {
		if leftList[leftIndex] <= rightList[rightIndex] {
			sorted[i] = leftList[leftIndex]
			leftIndex++
		} else {
			sorted[i] = rightList[rightIndex]
			rightIndex++
		}
		i++
	}
	// add any remaining elements from the left sorted list to the merged list
	for leftIndex < len(leftList) {
		sorted[i] = leftList[leftIndex]
		leftIndex++
		i++
	}
	// add any remaining elements from the right sorted list to the merged list
	for rightIndex < len(rightList) {
		sorted[i] = rightList[rightIndex]
		rightIndex++
		i++
	}

	return sorted
}
