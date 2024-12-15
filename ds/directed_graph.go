package ds

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type DirectedGraphPath[T comparable] []T

type DirectedGraph[T comparable] struct {
	edges      map[T][]T
	comparator func(a, b T) bool
}

func DefaultComparator[T constraints.Ordered](a, b T) bool {
	return a < b
}

func NewDirectedGraph[T comparable](comparator func(a, b T) bool) DirectedGraph[T] {
	return DirectedGraph[T]{
		edges:      make(map[T][]T),
		comparator: comparator,
	}
}

func NewDirectedGraphFromMap[T comparable](edges map[T][]T, comparator func(a, b T) bool) DirectedGraph[T] {
	return DirectedGraph[T]{
		edges:      edges,
		comparator: comparator,
	}
}

func (g DirectedGraph[T]) Size() int {
	return len(g.edges)
}

func (g DirectedGraph[T]) Vertices() []T {
	vertices := make([]T, 0, len(g.edges))
	for key := range g.edges {
		vertices = append(vertices, key)
	}
	sort.Slice(vertices, func(i, j int) bool {
		return g.comparator(vertices[i], vertices[j])
	})
	return vertices
}

func (g DirectedGraph[T]) GetNeighbors(vertex T) []T {
	neighbors := g.edges[vertex]
	sort.Slice(neighbors, func(i, j int) bool {
		return g.comparator(neighbors[i], neighbors[j])
	})
	return neighbors
}

func (g DirectedGraph[T]) AddVertex(vertex T) {
	if _, exists := g.edges[vertex]; !exists {
		g.edges[vertex] = []T{}
	}
}

func (g DirectedGraph[T]) AddEdge(vertex T, neighbor T) {
	if _, exists := g.edges[vertex]; !exists {
		g.AddVertex(vertex)
	}
	if _, exists := g.edges[neighbor]; !exists {
		g.AddVertex(neighbor)
	}

	updated := append(g.edges[vertex], neighbor)
	g.edges[vertex] = updated
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
				for _, neighbor := range g.GetNeighbors(v) {
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

func (g DirectedGraph[T]) FindPath(start T, end T) DirectedGraphPath[T] {
	queue := NewQueue[*DirectedGraphPath[T]]()

	starting := make(DirectedGraphPath[T], 0)
	starting = append(starting, start)

	queue.Enqueue(&starting)

	for !queue.IsEmpty() {
		currentPath := queue.Dequeue()
		lastNode := (*currentPath)[len(*currentPath)-1]

		if lastNode == end {
			return *currentPath
		}

		for _, neighbor := range g.GetNeighbors(lastNode) {
			newPath := DirectedGraphPath[T]{}
			newPath = append(newPath, *currentPath...)
			newPath = append(newPath, neighbor)

			queue.Enqueue(&newPath)
		}
	}

	return nil
}

func (g DirectedGraph[T]) FindDistinctPaths(start T, end T) []DirectedGraphPath[T] {
	paths := []DirectedGraphPath[T]{}

	queue := NewQueue[*DirectedGraphPath[T]]()

	starting := make(DirectedGraphPath[T], 0)
	starting = append(starting, start)

	queue.Enqueue(&starting)

	for !queue.IsEmpty() {
		currentPath := queue.Dequeue()
		lastNode := (*currentPath)[len(*currentPath)-1]

		if lastNode == end {
			paths = append(paths, *currentPath)
		}

		for _, neighbor := range g.GetNeighbors(lastNode) {
			if !currentPath.HasBeenVisited(neighbor) {
				newPath := DirectedGraphPath[T]{}
				newPath = append(newPath, neighbor)

				queue.Enqueue(&newPath)
			}
		}
	}

	return paths
}

func (visited DirectedGraphPath[T]) HasBeenVisited(node T) bool {
	for _, visitedNode := range visited {
		if visitedNode == node {
			return true
		}
	}
	return false
}
