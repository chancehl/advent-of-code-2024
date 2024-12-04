package math

func Make2DMatrix[T any](rows int, cols int) [][]T {
	matrix := make([][]T, rows)
	for row := range matrix {
		matrix[row] = make([]T, cols)
	}
	return matrix
}
