package search

import "golang.org/x/exp/constraints"

// Binary searches for the target (t) element in the sorted list (l) and returns the element and
// true if the element is found. If the element is not found a zero value is returned and false is
// returned indicating that the element does not exist in the list.
func Binary[T constraints.Ordered](l []T, t T) (elem T, exists bool) {
	low := 0
	high := len(l) - 1

	for low <= high {
		mid := (low + high) / 2

		if t < l[mid] {
			// search left sub list
			high = mid - 1
		} else if t > l[mid] {
			// search right sub list
			low = mid + 1
		} else {
			// we have found the element
			return l[mid], true
		}
	}

	var notFound T
	return notFound, false
}
