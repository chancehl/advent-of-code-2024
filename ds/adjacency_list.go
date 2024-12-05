package ds

import (
	"fmt"
	"slices"
)

type AdjacencyList map[int][]int

func (list AdjacencyList) Get(value int) []int {
	return list[value]
}

func (list AdjacencyList) Insert(value int, new int) error {
	if list[value] != nil && slices.Contains(list[value], new) {
		return fmt.Errorf("value %d already exists in adjacency list (nodes=%v)", value, list[value])
	}

	updated := append(list[value], new)
	list[value] = updated

	return nil
}
