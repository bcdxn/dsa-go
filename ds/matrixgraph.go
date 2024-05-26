package ds

import "fmt"

type MatrixGraph struct {
	matrix [][]bool
}

type MatrixGraphNode struct {
	Row    int
	Column int
}

// NewMatrixGraph returns a new instance of an undirected graph implemented using a matrix.
func NewMatrixGraph(matrix [][]bool) *MatrixGraph {
	return &MatrixGraph{
		matrix,
	}
}

// Stringify returns a string representation of the graph, showing the underlying matrix as a
// series of rows and columns made up of 1s and 0s.
func (g *MatrixGraph) Stringify() string {
	str := ""
	for _, row := range g.matrix {
		for _, cell := range row {
			if cell {
				str += "1 "
			} else {
				str += "0 "
			}
		}
		str += "\n"
	}

	return str
}

// NumPaths returns the number of valid paths from the start node to the end node (specified by
// row and column).
func (g *MatrixGraph) NumPaths(startRow, startCol, destRow, destCol int) int {
	visited := make(map[string]struct{})
	return g.numPaths(destRow, destCol, startRow, startCol, &visited)
}

/* Private helper functions
------------------------------------------------------------------------------------------------- */

// numPaths is a recursive helper function to count paths between two nodes in the graph using DFS.
func (g *MatrixGraph) numPaths(
	destRow int,
	destCol int,
	row int,
	col int,
	visited *map[string]struct{},
) int {
	width := len(g.matrix)
	// out of bounds (wider than row)
	if row < 0 || row >= width {
		return 0
	}
	height := len(g.matrix[row])
	// out of bounds (taller than column)
	if col < 0 || col >= height {
		return 0
	}
	// we reached a non-traversable node in the graph
	if !g.matrix[row][col] {
		return 0
	}
	// The node has already been visited
	_, hasVisited := (*visited)[nodeKey(row, col)]
	if hasVisited {
		return 0
	}
	// We've found the destination node
	if row == destRow && col == destCol {
		return 1
	}

	// The node is traversable; add it to the visited set
	(*visited)[nodeKey(row, col)] = struct{}{}

	count := 0
	// recursively check num paths at each neighboring node
	count += g.numPaths(destRow, destCol, row+1, col, visited)
	count += g.numPaths(destRow, destCol, row-1, col, visited)
	count += g.numPaths(destRow, destCol, row, col+1, visited)
	count += g.numPaths(destRow, destCol, row, col-1, visited)
	// Remove node from visited list because it could be a part of other paths
	delete(*visited, nodeKey(row, col))

	return count
}

// nodeKey calculates a unique key for a node in the graph to be stored in the visited 'set' used
// by the recursive DFS numPaths function so that we don't get stuck in cycles.
func nodeKey(row, col int) string {
	return fmt.Sprintf("%d%d", row, col)
}
