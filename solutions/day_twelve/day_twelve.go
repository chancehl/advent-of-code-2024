package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/chancehl/advent-of-code-2024/ds"
	"github.com/chancehl/advent-of-code-2024/utils/input"
	"github.com/chancehl/advent-of-code-2024/utils/timer"
)

func main() {
	path, err := filepath.Abs("solutions/day_twelve/input.txt")
	if err != nil {
		log.Fatalf("failed to construct path to input: %v", err)
	}

	input, err := input.Read(path)
	if err != nil {
		log.Fatalf("failed to read %s: %v", path, err)
	}

	dayTwelveSolution(input)
}

func dayTwelveSolution(input string) (int, int) {
	partOneResult, partOneRuntime := timer.ExecuteTimedFunc(PartOne, input)
	partTwoResult, partTwoRuntime := timer.ExecuteTimedFunc(PartTwo, input)

	fmt.Printf("[day twelve / part one] result=%d, time=%.2fms\n", partOneResult, partOneRuntime)
	fmt.Printf("[day twelve / part two] result=%d, time=%.2fms\n", partTwoResult, partTwoRuntime)

	return partOneResult, partTwoResult
}

func PartOne(input string) int {
	matrix := Create2DMatrix(input)
	ids := GetPlotIdentifiers(matrix)

	for _, id := range ids {
		plots := FindPlots(id, matrix)
		fmt.Println(id, plots)
	}

	return -1
}

func PartTwo(input string) int {
	return -1
}

func Create2DMatrix(input string) [][]string {
	matrix := [][]string{}
	for _, line := range strings.Split(input, "\n") {
		matrix = append(matrix, strings.Split(line, ""))
	}
	return matrix
}

func GetPlotIdentifiers(matrix [][]string) []string {
	identifiers := ds.NewSet[string]()
	for row := range matrix {
		for col := range matrix[0] {
			identifiers.Add(matrix[row][col])
		}
	}
	return identifiers.Values()
}

func FindPlots(id string, matrix [][]string) int {
	visited := ds.NewSet[ds.Coordinates]()
	plots := 0

	var bfs func(ds.Coordinates) = func(c ds.Coordinates) {
		queue := ds.NewQueue[ds.Coordinates]()
		queue.Enqueue(c)
		visited.Add(c)

		for queue.Size() > 0 {
			current := queue.Dequeue()

			directions := []ds.Coordinates{
				{Row: current.Row - 1, Col: current.Col}, // up
				{Row: current.Row + 1, Col: current.Col}, // down
				{Row: current.Row, Col: current.Col - 1}, // left
				{Row: current.Row, Col: current.Col + 1}, // right
			}

			for _, direction := range directions {
				if !visited.Has(direction) && IsInBounds(matrix, direction) && matrix[direction.Row][direction.Col] == id {
					queue.Enqueue(direction)
					visited.Add(direction)
				}
			}
		}
	}

	for row := range matrix {
		for col := range matrix[0] {
			coords := ds.Coordinates{Row: row, Col: col}
			if !visited.Has(coords) && matrix[row][col] == id {
				bfs(coords)
				plots++
			}
		}
	}

	return plots
}

func IsInBounds(matrix [][]string, coords ds.Coordinates) bool {
	return coords.Row > 0 && coords.Row < len(matrix)-1 && coords.Col > 0 && coords.Col < len(matrix[0])-1
}
