package math

import (
	"fmt"
	"math"

	"github.com/chancehl/advent-of-code-2024/ds"
)

func Make2DMatrix[T int | string](rows int, cols int) ds.Matrix[T] {
	matrix := make([][]T, rows)
	for row := range matrix {
		matrix[row] = make([]T, cols)
	}
	return matrix
}

func ComputeDistance(x1, y1, x2, y2 int) float64 {
	return math.Abs(float64(x2)-float64(x1)) + math.Abs(float64(y2)-float64(y1))
}

func PrintMatrix[T int | string](matrix ds.Matrix[T]) {
	for row := range matrix {
		fmt.Println(row)
	}
}
