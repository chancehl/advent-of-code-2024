package ds

import (
	"fmt"
	"slices"
)

type DirectedGraph[T comparable] map[T][]T

func NewDirectedGraph[T comparable]() DirectedGraph[T] {
	return make(map[T][]T)
}

func (g DirectedGraph[T]) Size() int {
	return len(g)
}

func (g DirectedGraph[T]) Vertices() []T {
	vertices := make([]T, 0, len(g))
	for key := range g {
		vertices = append(vertices, key)
	}
	return vertices
}

func (g DirectedGraph[T]) Get(key T) []T {
	return g[key]
}

func (g DirectedGraph[T]) AddKey(key T) {
	g[key] = []T{}
}

func (g DirectedGraph[T]) Insert(key T, value T) error {
	if g[key] != nil && slices.Contains(g[key], value) {
		return fmt.Errorf("value already exists in adjacency list (nodes=%v)", g[key])
	}

	updated := append(g[key], value)
	g[key] = updated

	return nil
}

func (g DirectedGraph[T]) TopologicalSort() []T {
	sorted := []T{}

	visited := make(map[T]bool)
	stack := Stack[T]{}

	for _, vertex := range g.Vertices() {
		if !visited[vertex] {
			var dfs func(v T)

			dfs = func(v T) {
				// mark as visited
				visited[v] = true

				// visit neighbors
				for _, neighbor := range g.Get(v) {
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
