package ds

import (
	"maps"
	"slices"
)

type Set[T comparable] struct {
	data map[T]bool
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{data: make(map[T]bool)}
}

func NewSetFrom[T comparable](vals []T) Set[T] {
	set := NewSet[T]()
	for _, val := range vals {
		set.Add(val)
	}
	return set
}

func (s *Set[T]) Has(val T) bool {
	return s.data[val]
}

func (s *Set[T]) Add(val T) {
	s.data[val] = true
}

func (s *Set[T]) Remove(val T) {
	delete(s.data, val)
}

func (s *Set[T]) Values() []T {
	return slices.Collect(maps.Keys(s.data))
}

func (s *Set[T]) Size() int {
	return len(s.Values())
}
