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
