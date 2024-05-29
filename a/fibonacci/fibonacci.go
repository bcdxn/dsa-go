package fibonacci

// BruteForce calculates the nth fibonacci recursively with no optimizations
func BruteForce(n int) int {
	if n <= 1 {
		return n
	}

	return BruteForce(n-1) + BruteForce(n-2)
}

// Memoization calculates the nth fibonacci number recurisvely optmizing with memoization
func Memoization(n int) int {
	cache := make([]int, n+1)
	return memoization(n, cache)
}

func memoization(n int, cache []int) int {
	if n <= 1 {
		return n
	}

	if cache[n] > 0 {
		return cache[n]
	}

	cache[n] = memoization(n-1, cache) + memoization(n-2, cache)
	return cache[n]
}

func BottomUp(n int) int {
	if n <= 1 {
		return n
	}
	// use an array to hold the previous values
	s := [2]int{0, 1}
	i := 2

	for i <= n {
		tmp := s[1]
		s[1] = s[0] + s[1]
		s[0] = tmp
		i++
	}

	return s[1]
}
