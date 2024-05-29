package ds_test

import (
	"testing"

	"github.com/bcdxn/dsa-go/ds"
	"github.com/stretchr/testify/assert"
)

func TestNewAlGraph(t *testing.T) {
	t.Run("should initialize the data structure properly", func(t *testing.T) {
		g, err := ds.NewAlGraph([][]string{
			{"A", "B"},
			{"B", "C"},
			{"B", "E"},
			{"C", "E"},
			{"E", "D"},
		})

		m := map[string][]string{
			"A": {"B"},
			"B": {"C", "E"},
			"C": {"E"},
			"E": {"D"},
		}

		assert.Nil(t, err)
		assert.Equal(t, m, g.GetMap())
	})

	t.Run("should return an error if the edge list is malformed", func(t *testing.T) {
		_, err := ds.NewAlGraph([][]string{
			{"A", "B", "A", "E"},
			{"B", "C"},
			{"B", "E"},
			{"C", "E"},
			{"E", "D"},
		})
		assert.NotNil(t, err)
	})
}

func TestAlGraphNumPaths(t *testing.T) {
	t.Run("should find all paths from start to destination vertices", func(t *testing.T) {
		g, err := ds.NewAlGraph([][]string{
			{"A", "B"},
			{"B", "C"},
			{"B", "E"},
			{"C", "E"},
			{"E", "D"},
		})

		assert.Nil(t, err)
		assert.Equal(t, g.NumPaths("A", "E"), 2)
	})
}
