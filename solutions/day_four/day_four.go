package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/chancehl/advent-of-code-2024/utils/input"
	"github.com/chancehl/advent-of-code-2024/utils/math"
	"github.com/chancehl/advent-of-code-2024/utils/timer"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
	DiagonalUpLeft
	DiagonalUpRight
	DiagonalDownLeft
	DiagonalDownRight
)

const Xmas = "XMAS"

func main() {
	path, err := filepath.Abs("solutions/day_four/input.txt")
	if err != nil {
		log.Fatalf("failed to construct path to input: %v", err)
	}

	input, err := input.Read(path)
	if err != nil {
		log.Fatalf("failed to read %s: %v", path, err)
	}

	dayFourSolution(input)
}

func dayFourSolution(input string) (int, int) {
	partOneResult, partOneRuntime := timer.ExecuteTimedFunc(PartOne, input)
	partTwoResult, partTwoRuntime := timer.ExecuteTimedFunc(PartTwo, input)

	fmt.Printf("[day four / part one] result=%d, time=%.2fms\n", partOneResult, partOneRuntime)
	fmt.Printf("[day four / part two] result=%d, time=%.2fms\n", partTwoResult, partTwoRuntime)

	return partOneResult, partTwoResult
}

func PartOne(input string) int {
	appearances := 0

	matrix := createMatrixFromInput(input)

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == "X" {
				appearances += search(matrix, i, j)
			}
		}
	}

	return appearances
}

func PartTwo(input string) int {
	return -1
}

func createMatrixFromInput(input string) [][]string {
	lines := strings.Split(input, "\n")

	rows := len(lines)
	cols := len(lines[0])

	matrix := math.Make2DMatrix[string](rows, cols)

	for i, l := range lines {
		for j, c := range strings.Split(l, "") {
			matrix[i][j] = c
		}
	}

	return matrix
}

func search(matrix [][]string, row, col int) int {
	matches := 0

	// up
	if searchUp(matrix, row, col) {
		matches += 1
	}
	// down
	if searchDown(matrix, row, col) {
		matches += 1
	}
	// left
	if searchLeft(matrix, row, col) {
		matches += 1
	}
	// right
	if searchRight(matrix, row, col) {
		matches += 1
	}
	// diagonal up left
	if searchUpLeft(matrix, row, col) {
		matches += 1
	}
	// diagonal up right
	if searchUpRight(matrix, row, col) {
		matches += 1
	}
	// diagonal down left
	if searchDownLeft(matrix, row, col) {
		matches += 1
	}
	// diagonal down right
	if searchDownRight(matrix, row, col) {
		matches += 1
	}

	return matches
}

func searchUp(matrix [][]string, row, col int) bool {
	if row < 3 {
		return false
	}

	letters := []string{
		matrix[row][col],
		matrix[row-1][col],
		matrix[row-2][col],
		matrix[row-3][col],
	}

	return strings.Join(letters, "") == Xmas
}

func searchDown(matrix [][]string, row, col int) bool {
	if (row + 3) > len(matrix)-1 {
		return false
	}

	letters := []string{
		matrix[row][col],
		matrix[row+1][col],
		matrix[row+2][col],
		matrix[row+3][col],
	}

	return strings.Join(letters, "") == Xmas
}

func searchLeft(matrix [][]string, row, col int) bool {
	if col < 3 {
		return false
	}

	letters := []string{
		matrix[row][col],
		matrix[row][col-1],
		matrix[row][col-2],
		matrix[row][col-3],
	}

	return strings.Join(letters, "") == Xmas
}

func searchRight(matrix [][]string, row, col int) bool {
	if (col + 3) > len(matrix[0])-1 {
		return false
	}

	letters := []string{
		matrix[row][col],
		matrix[row][col+1],
		matrix[row][col+2],
		matrix[row][col+3],
	}

	return strings.Join(letters, "") == Xmas
}

func searchUpLeft(matrix [][]string, row, col int) bool {
	if row < 3 || col < 3 {
		return false
	}

	letters := []string{
		matrix[row][col],
		matrix[row-1][col-1],
		matrix[row-2][col-2],
		matrix[row-3][col-3],
	}

	return strings.Join(letters, "") == Xmas
}

func searchUpRight(matrix [][]string, row, col int) bool {
	if row < 3 || (col+3) > len(matrix[0])-1 {
		return false
	}

	letters := []string{
		matrix[row][col],
		matrix[row-1][col+1],
		matrix[row-2][col+2],
		matrix[row-3][col+3],
	}

	return strings.Join(letters, "") == Xmas
}

func searchDownLeft(matrix [][]string, row, col int) bool {
	if (row+3) > len(matrix)-1 || col < 3 {
		return false
	}

	letters := []string{
		matrix[row][col],
		matrix[row+1][col-1],
		matrix[row+2][col-2],
		matrix[row+3][col-3],
	}

	return strings.Join(letters, "") == Xmas
}

func searchDownRight(matrix [][]string, row, col int) bool {
	if (row+3) > len(matrix)-1 || (col+3) > len(matrix[0])-1 {
		return false
	}

	letters := []string{
		matrix[row][col],
		matrix[row+1][col+1],
		matrix[row+2][col+2],
		matrix[row+3][col+3],
	}

	return strings.Join(letters, "") == Xmas
}
