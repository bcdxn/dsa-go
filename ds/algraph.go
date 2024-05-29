package ds

import "errors"

type AlGraph struct {
	// We can mimick an adjacency list implementation with a map of slices (as long as all node keys
	// are unique)
	m map[string][]string
}

func NewAlGraph(edges [][]string) (*AlGraph, error) {
	m := make(map[string][]string)
	for _, edge := range edges {
		if len(edge) != 2 {
			return nil, errors.New("invalid edge format; must be a vertex couple")
		}

		s := edge[0]
		d := edge[1]

		if _, exists := m[s]; !exists {
			m[s] = []string{d}
		} else {
			m[s] = append(m[s], d)
		}
	}

	return &AlGraph{m}, nil
}

func (g *AlGraph) GetMap() map[string][]string {
	return g.m
}

func (g *AlGraph) NumPaths(start, end string) int {
	visited := make(map[string]struct{})
	return g.numPaths(start, end, &visited)
}

func (g *AlGraph) numPaths(curr, end string, visited *map[string]struct{}) int {
	if curr == end {
		return 1
	}

	if _, hasVisited := (*visited)[curr]; hasVisited {
		return 0
	}

	(*visited)[curr] = struct{}{}
	count := 0
	// Run numPaths on all neighbors of the current node
	for _, n := range g.m[curr] {
		count += g.numPaths(n, end, visited)
	}
	// Remove node from visited list to allow for other paths "higher up" the recursive stack
	delete(*visited, curr)

	return count
}
