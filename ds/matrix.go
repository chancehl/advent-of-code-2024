package ds

import (
	"strconv"
	"strings"
)

type Matrix[T int | string] [][]T

func CreateStringMatrix(input string) Matrix[string] {
	matrix := [][]string{}
	for _, line := range strings.Split(input, "\n") {
		matrix = append(matrix, strings.Split(line, ""))
	}
	return matrix
}

func CreateIntMatrix(input string) Matrix[int] {
	matrix := [][]int{}
	for _, line := range strings.Split(input, "\n") {
		nums := []int{}
		for _, char := range strings.Split(line, "") {
			num, _ := strconv.Atoi(char)
			nums = append(nums, num)
		}
		matrix = append(matrix, nums)
	}
	return matrix
}

func (m *Matrix[T]) IsInBounds(row, col int) bool {
	return row >= 0 && row < len(*m) && col >= 0 && col < len((*m)[0])
}

func (m *Matrix[T]) IsCoordInBounds(c Coordinates) bool {
	return c.Row >= 0 && c.Row < len(*m) && c.Col >= 0 && c.Col < len((*m)[0])
}
