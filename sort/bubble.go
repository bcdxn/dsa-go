package sort

import "golang.org/x/exp/constraints"

// Bubble returns a sorted copy of the list using the [Bubble Sort alogrithm][0]. This function does
// not alter the given list.
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
// [0]: https://en.wikipedia.org/wiki/Bubble_sort
func Bubble[T constraints.Ordered](list []T) []T {
	cpy := make([]T, len(list))
	copy(cpy, list)

	swap := true

	for swap {
		swap = false

		for i := 0; i < len(cpy)-1; i++ {
			if cpy[i] > cpy[i+1] {
				cpy[i], cpy[i+1] = cpy[i+1], cpy[i]
				swap = true
			}
		}

	}

	return cpy
}

// Bubble sorts the given list using the [Bubble Sort alogrithm][0]. This function alters the given
// list.
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
// [0]: https://en.wikipedia.org/wiki/Bubble_sort
func InPlaceBubble[T constraints.Ordered](list []T) []T {
	swap := true

	for swap {
		swap = false

		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swap = true
			}
		}
	}

	return list
}
