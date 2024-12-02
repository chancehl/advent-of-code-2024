package math

func Make2DMatrix(rows int, cols int) [][]int {
	matrix := make([][]int, rows)
	for row := range matrix {
		matrix[row] = make([]int, cols)
	}
	return matrix
}
