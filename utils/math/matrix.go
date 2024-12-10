package math

import "math"

func Make2DMatrix[T any](rows int, cols int) [][]T {
	matrix := make([][]T, rows)
	for row := range matrix {
		matrix[row] = make([]T, cols)
	}
	return matrix
}

func ComputeDistance(x1, y1, x2, y2 int) float64 {
	return math.Abs(float64(x2)-float64(x1)) + math.Abs(float64(y2)-float64(y1))
}
