package slices

func Map[T, U any](s []T, f func(T) U) []U {
	updated := make([]U, len(s))
	for _, element := range s {
		updated = append(updated, f(element))
	}
	return updated
}
