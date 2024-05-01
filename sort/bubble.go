package sort

// Return a sorted copy of the list using the [Bubble Sort alogrithm]. This function does not alter
// the given list
//
// **Time Complexity**
//
// | Scenario     | Big-O  |
// |:-------------|:-------|
// | Average Case | O(n^2) |
// | Best Case    | O(n)   |
// | Worst Case   | O(n^2) |
//
// [Bubble Sort alogrithm]: https://en.wikipedia.org/wiki/Bubble_sort
func Bubble(list []int) []int {
	cpy := make([]int, len(list))
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

func InPlaceBubble(list []int) []int {
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
