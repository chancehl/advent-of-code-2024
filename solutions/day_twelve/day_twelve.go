package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/chancehl/advent-of-code-2024/ds"
	"github.com/chancehl/advent-of-code-2024/utils/input"
	"github.com/chancehl/advent-of-code-2024/utils/timer"
)

type ElfPlot struct {
	id   string
	land []ds.Coordinates
}

type Direction string

const (
	Up    = "U"
	Down  = "D"
	Left  = "L"
	Right = "R"
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
	matrix := ds.CreateStringMatrix(input)
	ids := GetPlotIdentifiers(matrix)

	totalPrice := 0

	for _, id := range ids {
		plots := FindPlots(id, matrix)
		for _, plot := range plots {
			area := plot.ComputeArea()
			perimiter := plot.ComputePerimeter(matrix)
			totalPrice += area * perimiter
		}
	}

	return totalPrice
}

func PartTwo(input string) int {
	matrix := ds.CreateStringMatrix(input)
	ids := GetPlotIdentifiers(matrix)

	totalPrice := 0

	for _, id := range ids {
		plots := FindPlots(id, matrix)
		for _, plot := range plots {
			area := plot.ComputeArea()
			sides := plot.ComputeSides(matrix)
			fmt.Printf("plot %s has %d sides\n", plot.id, sides)
			totalPrice += area * sides
		}
	}

	return totalPrice
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

func FindPlots(id string, matrix ds.Matrix[string]) []ElfPlot {
	visited := ds.NewSet[ds.Coordinates]()
	plots := []ElfPlot{}

	var bfs func(ds.Coordinates, *ElfPlot) = func(c ds.Coordinates, plot *ElfPlot) {
		queue := ds.NewQueue[ds.Coordinates]()
		queue.Enqueue(c)
		visited.Add(c)

		for !queue.IsEmpty() {
			current := queue.Dequeue()

			for _, direction := range ds.GetInBoundsNeighbors(current, matrix) {
				if !visited.Has(direction) && matrix[direction.Row][direction.Col] == id {
					queue.Enqueue(direction)
					visited.Add(direction)

					plot.land = append(plot.land, direction)
				}
			}
		}
	}

	for row := range matrix {
		for col := range matrix[0] {
			coords := ds.Coordinates{Row: row, Col: col}
			if !visited.Has(coords) && matrix[row][col] == id {
				plot := ElfPlot{id: id, land: []ds.Coordinates{coords}}
				bfs(coords, &plot)
				plots = append(plots, plot)
			}
		}
	}

	return plots
}

func (p ElfPlot) ComputeArea() int {
	return len(p.land)
}

func (p ElfPlot) ComputePerimeter(matrix ds.Matrix[string]) int {
	perimeter := 0
	for _, space := range p.land {
		neighbors := ds.GetNeighbors(space)
		for _, neighbor := range neighbors {
			if matrix.IsCoordInBounds(neighbor) && matrix[neighbor.Row][neighbor.Col] != p.id {
				perimeter++
			} else if !matrix.IsCoordInBounds(neighbor) {
				perimeter++
			}
		}
	}
	return perimeter
}

func (p ElfPlot) ComputeSides(matrix ds.Matrix[string]) int {
	sides := 0
	return sides
}
