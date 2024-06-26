package ds

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

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

// ShortestPath returns the length of the shortest path from the given start row, column to the
// given destination row, column
func (g *MatrixGraph) ShortestPath(startRow, startCol, destRow, destCol int) (int, error) {
	visited := make(map[string]struct{})
	q := NewQueue[string]()
	q.Enqueue(nodeKey(startRow, startCol))
	visited[nodeKey(startRow, startCol)] = struct{}{}
	length := 0

	// Continue looping while the queue is not empty
	for q.Depth() > 0 {
		depth := q.Depth()

		for range depth {
			key, _ := q.Dequeue()
			r, c := decodeKey(key)

			if r == destRow && c == destCol {
				// we've arrived at the destination
				return length, nil
			}
			// Add all valid neighbors of the current node in the graph; check validity and add to queue
			if g.isValidNeighbor(r-1, c, visited) {
				q.Enqueue(nodeKey(r-1, c))
				visited[nodeKey(r-1, c)] = struct{}{}
			}
			if g.isValidNeighbor(r+1, c, visited) {
				q.Enqueue(nodeKey(r+1, c))
				visited[nodeKey(r+1, c)] = struct{}{}
			}
			if g.isValidNeighbor(r, c-1, visited) {
				q.Enqueue(nodeKey(r, c-1))
				visited[nodeKey(r, c-1)] = struct{}{}
			}
			if g.isValidNeighbor(r, c+1, visited) {
				q.Enqueue(nodeKey(r, c+1))
				visited[nodeKey(r, c+1)] = struct{}{}
			}
		}

		length += 1
	}

	return 0, errors.New("no path from start to destination found")
}

// NumIslands returns the number of groupings of 'true' values in the graph using DFS.
// e.g., there are 2 'islands' in the graph below:
// 1 1 0 0
// 1 0 0 1
// 1 0 0 1
func (g *MatrixGraph) NumIslands() int {
	visited := make(map[string]struct{})
	islandCount := 0

	for r := range len(g.matrix) {
		for c := range len(g.matrix[r]) {
			_, hasVisited := visited[nodeKey(r, c)]
			if g.matrix[r][c] && !hasVisited {
				islandCount++
				g.visitIslandNodes(r, c, &visited)
			}
		}
	}

	return islandCount
}

// NumIslands returns the number of groupings of 'true' values in the graph using BFS.
func (g *MatrixGraph) NumIslandsBfs() int {
	if len(g.matrix) < 1 && len(g.matrix[0]) < 1 {
		// empty matrix cannot have islands
		return 0
	}

	count := 0
	visited := make(map[string]struct{})

	for r := range len(g.matrix) {
		for c := range len(g.matrix[r]) {
			if _, hasVisited := visited[nodeKey(r, c)]; !hasVisited && g.matrix[r][c] {
				count++
				g.numIslandsBfs(r, c, &visited)
			}
		}
	}

	return count
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
	// out of bounds (wider than row)
	if row < 0 || row >= len(g.matrix) {
		return 0
	}
	// out of bounds (taller than column)
	if col < 0 || col >= len(g.matrix[row]) {
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

// A depth first search helper function that does not remove nodes from the path (useful when
// counting groups or 'islands' within the graph).
func (g *MatrixGraph) visitIslandNodes(row, col int, visited *map[string]struct{}) {
	if row < 0 || row >= len(g.matrix) {
		return // out of bounds
	}
	if col < 0 || col >= len(g.matrix[row]) {
		return // out of bounds
	}
	if !g.matrix[row][col] {
		return // hit water
	}
	_, hasVisited := (*visited)[nodeKey(row, col)]
	if hasVisited {
		return // already a part of an island
	}

	// Add the node to the visited set
	(*visited)[nodeKey(row, col)] = struct{}{}

	g.visitIslandNodes(row+1, col, visited)
	g.visitIslandNodes(row-1, col, visited)
	g.visitIslandNodes(row, col+1, visited)
	g.visitIslandNodes(row, col-1, visited)
}

func (g *MatrixGraph) isValidNeighbor(row, col int, visited map[string]struct{}) bool {
	if row < 0 || row >= len(g.matrix) {
		return false
	}
	if col < 0 || col >= len(g.matrix[row]) {
		return false
	}
	_, ok := visited[nodeKey(row, col)]
	return !ok && g.matrix[row][col]
}

// nodeKey calculates a unique key for a node in the graph to be stored in the visited 'set' used
// by the recursive DFS numPaths function so that we don't get stuck in cycles.
func nodeKey(row, col int) string {
	return fmt.Sprintf("%d:%d", row, col)
}

func decodeKey(key string) (int, int) {
	rc := strings.Split(key, ":")
	r, err := strconv.Atoi(rc[0])
	if err != nil {
		panic("invalid row in key")
	}
	c, err := strconv.Atoi(rc[1])
	if err != nil {
		panic("invalid column key")
	}
	return r, c
}

func (g *MatrixGraph) numIslandsBfs(startRow, startCol int, v *map[string]struct{}) {
	q := NewQueue[string]()
	(*v)[nodeKey(startRow, startCol)] = struct{}{}
	q.Enqueue(nodeKey(startRow, startCol))

	for q.Depth() > 0 {
		l := q.Depth()

		for range l {
			rc, err := q.Dequeue()

			if err != nil {
				panic("invalid queue operation; dequing empty queue")
			}

			r, c := decodeKey(rc)

			if g.isValidNeighbor(r-1, c, *v) {
				q.Enqueue(nodeKey(r-1, c))
				(*v)[nodeKey(r-1, c)] = struct{}{}
			}
			if g.isValidNeighbor(r+1, c, *v) {
				q.Enqueue(nodeKey(r+1, c))
				(*v)[nodeKey(r+1, c)] = struct{}{}
			}
			if g.isValidNeighbor(r, c-1, *v) {
				q.Enqueue(nodeKey(r, c-1))
				(*v)[nodeKey(r, c-1)] = struct{}{}
			}
			if g.isValidNeighbor(r, c+1, *v) {
				q.Enqueue(nodeKey(r, c+1))
				(*v)[nodeKey(r, c+1)] = struct{}{}
			}
		}
	}
}
