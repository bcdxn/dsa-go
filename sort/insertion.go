package sort

import "golang.org/x/exp/constraints"

// Insertion returns a sorted copy of the given list using the [Insertion Sort algorithm][0]. This
// function does not alter the given list.
//
// **Time Complexity**
//
// ```
// | Scenario     | Big-O  |
// |:-------------|:-------|
// | Average Case | O(n^2) |
// | Best Case    | O(n)   |
// | Worst Case   | O(n^2) |
// ```
//
// [0]: https://en.wikipedia.org/wiki/Insertion_sort
func Insertion[T constraints.Ordered](list []T) []T {
	// Copy the list to ensure there are no side effects on the given list
	cpy := make([]T, len(list))
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

// InPlaceInsertion sorts the given list using the [Insertion Sort algorithm][0]. This function
// alters the given list.
//
// **Time Complexity**
//
// ```
// | Scenario     | Big-O  |
// |:-------------|:-------|
// | Average Case | O(n^2) |
// | Best Case    | O(n)   |
// | Worst Case   | O(n^2) |
// ```
//
// [0]: https://en.wikipedia.org/wiki/Insertion_sort
func InPlaceInsertion[T constraints.Ordered](list []T) []T {
	for i := 1; i < len(list); i++ {
		j := i
		for j > 0 && list[j-1] > list[j] {
			list[j-1], list[j] = list[j], list[j-1]
			j--
		}
	}

	return list
}
