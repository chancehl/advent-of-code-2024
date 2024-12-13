package ds

import (
	"slices"
	"testing"
)

func TestAdjacencyList(t *testing.T) {
	var tests = []struct {
		list     DirectedGraph[int]
		expected []int
	}{
		{
			list:     NewDirectedGraph[int](DefaultComparator),
			expected: []int{},
		},
		{
			list: NewDirectedGraphFromMap(map[int][]int{
				0: {},
				1: {},
				2: {3},
				3: {1},
				4: {0, 1},
				5: {2, 0},
			}, DefaultComparator),
			expected: []int{5, 4, 2, 3, 1, 0},
		},
	}

	for _, test := range tests {
		t.Run("topological sort", func(t *testing.T) {
			actual := test.list.TopologicalSort()

			if !slices.Equal(test.expected, actual) {
				t.Errorf("test failed (expected=%d, actual=%d)\n", test.expected, actual)
			}
		})
	}

}
