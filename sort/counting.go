package sort

// Counting sort returns a sorted copy of the list using the [Counting Sort alogrithm][0]. This
// function does not alter the given list.
//
// **Time Complexity**
//
// ```
// | Scenario     | Big-O    |
// |:-------------|:---------|
// | Average Case | O(n + k) |
// | Best Case    | O(n + k) |
// | Worst Case   | O(n + k) |
// ```
//
// [0]: https://en.wikipedia.org/wiki/Counting_sort
func Counting(list []uint) []uint {
	cpy := make([]uint, len(list))
	copy(cpy, list)

	counting(cpy)

	return cpy
}

// Counting sort returns the sorted list using the [Counting Sort alogrithm][0]. This function
// alters the given list.
//
// **Time Complexity**
//
// ```
// | Scenario     | Big-O  |
// |:-------------|:-------|
// | Average Case | O(n)   |
// | Best Case    | O(n)   |
// | Worst Case   | O(n)   |
// ```
//
// [0]: https://en.wikipedia.org/wiki/Counting_sort
func InPlaceCounting(list []uint) []uint {
	return counting(list)
}

func counting(list []uint) []uint {
	// create the bucket
	bucket := make([]uint, len(list))
	// count the occurences of each number
	for i := 0; i < len(list); i++ {
		bucket[list[i]]++
	}
	// add the numbers back
	i := 0
	for elem, count := range bucket {
		for j := count; j > 0; j-- {
			list[i] = uint(elem)
			i++
		}
	}

	return list
}
