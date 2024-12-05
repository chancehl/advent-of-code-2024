package ds

import (
	"fmt"
	"slices"
)

type AdjacencyList map[int][]int

func (list AdjacencyList) Size() int {
	return len(list)
}

func (list AdjacencyList) Vertices() []int {
	vertices := make([]int, 0, len(list))
	for key := range list {
		vertices = append(vertices, key)
	}
	return vertices
}

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

func (list AdjacencyList) TopologicalSort() []int {
	sorted := []int{}

	visited := make(map[int]bool)
	stack := Stack{}

	for _, vertex := range list.Vertices() {
		if !visited[vertex] {
			var dfs func(v int)

			dfs = func(v int) {
				// mark as visited
				visited[v] = true

				// visit neighbors
				for _, neighbor := range list.Get(v) {
					if !visited[neighbor] {
						dfs(neighbor)
					}
				}

				// push onto stack
				stack.Push(v)
			}

			dfs(vertex)
		}
	}

	for !stack.IsEmpty() {
		sorted = append(sorted, *stack.Pop())
	}

	return sorted
}
