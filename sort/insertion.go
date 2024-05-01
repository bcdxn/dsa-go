package sort

// Return a sorted copy of the given list using the [insertion sort algorithm]. This function
// does not alter the given list.
//
// **Time Complexity**
//
// | Scenario     | Big-O  |
// |:-------------|:-------|
// | Average Case | O(n^2) |
// | Best Case    | O(n)   |
// | Worst Case   | O(n^2) |
//
// [insertion sort algorithm]: https://en.wikipedia.org/wiki/Insertion_sort
func Insertion(list []int) []int {
	// Copy the list to ensure there are no side effects on the given list
	cpy := make([]int, len(list))
	copy(cpy, list)

	for i := 1; i < len(cpy); i++ {
		j := i
		for j > 0 && cpy[j-1] > cpy[j] {
			cpy[j-1], cpy[j] = cpy[j], cpy[j-1]
			j--
		}
	}

	return cpy
}

// Sort the given list using the [insertion sort algorithm]. This function alters the given list.
// **Time Complexity**
//
// | Scenario     | Big-O  |
// |:-------------|:-------|
// | Average Case | O(n^2) |
// | Best Case    | O(n)   |
// | Worst Case   | O(n^2) |
//
// [insertion sort algorithm]: https://en.wikipedia.org/wiki/Insertion_sort
func InPlaceInsertion(list []int) []int {
	for i := 1; i < len(list); i++ {
		j := i
		for j > 0 && list[j-1] > list[j] {
			list[j-1], list[j] = list[j], list[j-1]
			j--
		}
	}

	return list
}
