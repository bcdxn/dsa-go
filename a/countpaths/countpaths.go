package countpaths

// BruteForce calculates the number of unique paths from the start to finish row, column in the
// given grid using brute-force DFS. a path can only move to down and to the right of the start
// row, column.
func BruteForce(r, c, destR, destC int, grid [][]int) int {
	if r >= len(grid) || c >= len(grid[r]) {
		// out of bounds
		return 0
	}
	if r == destR || c == destC {
		// we've reached the destination and found a valid path
		return 1
	}
	return BruteForce(r+1, c, destR, destC, grid) + BruteForce(r, c+1, destR, destC, grid)
}

func TopDown(r, c, destR, destC int, grid [][]int) int {
	cache := make([][]int, len(grid))
	for i := range len(cache) {
		cache[i] = make([]int, len(grid[i]))
	}
	return topDown(r, c, destR, destC, grid, cache)
}

func topDown(r, c, destR, destC int, grid [][]int, cache [][]int) int {
	if r >= len(grid) || c >= len(grid[r]) {
		// out of bounds
		return 0
	}
	if r == destR && c == destC {
		// we've reached the destination and found a valid path
		return 1
	}
	if cache[r][c] > 0 {
		return cache[r][c]
	}
	cache[r][c] = topDown(r+1, c, destR, destC, grid, cache) + topDown(r, c+1, destR, destC, grid, cache)
	return cache[r][c]
}
