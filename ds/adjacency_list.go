package ds

import (
	"fmt"
	"slices"
)

type AdjacencyList[T comparable] map[T][]T

func NewAdjacencyList[T comparable]() AdjacencyList[T] {
	return make(map[T][]T)
}

func (list AdjacencyList[T]) Size() int {
	return len(list)
}

func (list AdjacencyList[T]) Vertices() []T {
	vertices := make([]T, 0, len(list))
	for key := range list {
		vertices = append(vertices, key)
	}
	return vertices
}

func (list AdjacencyList[T]) Get(value T) []T {
	return list[value]
}

func (list AdjacencyList[T]) Insert(key T, value T) error {
	if list[key] != nil && slices.Contains(list[key], value) {
		return fmt.Errorf("value already exists in adjacency list (nodes=%v)", list[key])
	}

	updated := append(list[key], value)
	list[key] = updated

	return nil
}

func (list AdjacencyList[T]) TopologicalSort() []T {
	sorted := []T{}

	visited := make(map[T]bool)
	stack := Stack[T]{}

	for _, vertex := range list.Vertices() {
		if !visited[vertex] {
			var dfs func(v T)

			dfs = func(v T) {
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
