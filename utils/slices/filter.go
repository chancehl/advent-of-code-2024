package slices

func Filter[T any](s []T, filter func(T) bool) []T {
	data := []T{}
	for _, element := range s {
		if filter(element) {
			data = append(data, element)
		}
	}
	return data
}

func Any[T any](s []T, filter func(T) bool) bool {
	filtered := Filter(s, filter)
	return len(filtered) > 0
}
