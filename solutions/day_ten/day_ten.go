package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/chancehl/advent-of-code-2024/ds"
	"github.com/chancehl/advent-of-code-2024/utils/input"
	"github.com/chancehl/advent-of-code-2024/utils/timer"
)

type Path []ds.Coordinates

const MaxPathSize = 10 // len(0,1,2,3,4,5,6,7,8,9)
const DotValue = -1
const TrailStart = 0
const TrailEnd = 9

func main() {
	path, err := filepath.Abs("solutions/day_ten/input.txt")
	if err != nil {
		log.Fatalf("failed to construct path to input: %v", err)
	}

	input, err := input.Read(path)
	if err != nil {
		log.Fatalf("failed to read %s: %v", path, err)
	}

	dayTenSolution(input)
}

func dayTenSolution(input string) (int, int) {
	partOneResult, partOneRuntime := timer.ExecuteTimedFunc(PartOne, input)
	partTwoResult, partTwoRuntime := timer.ExecuteTimedFunc(PartTwo, input)

	fmt.Printf("[day ten / part one] result=%d, time=%.2fms\n", partOneResult, partOneRuntime)
	fmt.Printf("[day ten / part two] result=%d, time=%.2fms\n", partTwoResult, partTwoRuntime)

	return partOneResult, partTwoResult
}

func PartOne(input string) int {
	matrix := Create2DMatrix(input)

	graph := CreateGraph(matrix)
	trailheads := FindCoordinates(matrix, TrailStart)
	trailends := FindCoordinates(matrix, TrailEnd)

	score := 0

	for _, start := range trailheads {
		for _, end := range trailends {
			path := graph.FindPath(start, end)
			if path != nil {
				score++
			}
		}
	}

	return score
}

func PartTwo(input string) int {
	matrix := Create2DMatrix(input)

	graph := CreateGraph(matrix)
	trailheads := FindCoordinates(matrix, TrailStart)
	trailends := FindCoordinates(matrix, TrailEnd)

	validPaths := []ds.DirectedGraphPath[ds.Coordinates]{}

	for _, start := range trailheads {
		for _, end := range trailends {
			paths := graph.FindDistinctPaths(start, end)
			if paths != nil {
				validPaths = append(validPaths, paths...)
			}
		}
	}

	return len(validPaths)
}

func Create2DMatrix(input string) [][]int {
	matrix := [][]int{}
	for _, row := range strings.Split(input, "\n") {
		values := []int{}
		for _, col := range strings.Split(row, "") {
			if col == "." {
				values = append(values, DotValue)
			} else {
				num, _ := strconv.Atoi(col)
				values = append(values, num)
			}
		}
		matrix = append(matrix, values)
	}
	return matrix
}

func CreateGraph(matrix [][]int) ds.DirectedGraph[ds.Coordinates] {
	graph := ds.NewDirectedGraph(ds.CoordinateComparator)

	for row := range matrix {
		for col := range matrix[row] {
			coords := ds.Coordinates{Row: row, Col: col}
			current := matrix[row][col]

			// up
			if row > 0 {
				next := matrix[row-1][col]
				if next == current+1 {
					graph.Insert(coords, ds.Coordinates{Row: row - 1, Col: col})
				}
			}

			// down
			if row < len(matrix)-1 {
				next := matrix[row+1][col]
				if next == current+1 {
					graph.Insert(coords, ds.Coordinates{Row: row + 1, Col: col})

				}
			}

			// left
			if col > 0 {
				next := matrix[row][col-1]
				if next == current+1 {
					graph.Insert(coords, ds.Coordinates{Row: row, Col: col - 1})
				}
			}

			// right
			if col < len(matrix[0])-1 {
				next := matrix[row][col+1]
				if next == current+1 {
					graph.Insert(coords, ds.Coordinates{Row: row, Col: col + 1})

				}
			}
		}
	}

	return graph
}

func FindCoordinates(matrix [][]int, marker int) []ds.Coordinates {
	trailheads := []ds.Coordinates{}
	for row := range matrix {
		for col := range matrix[0] {
			if matrix[row][col] == marker {
				trailheads = append(trailheads, ds.Coordinates{Row: row, Col: col})
			}
		}
	}
	return trailheads
}
