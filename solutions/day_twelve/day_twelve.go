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

type ElfPlotBoundary struct {
	coord ds.Coordinates
	dir   Direction
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

			totalPrice += area * len(perimiter)
		}
	}

	return totalPrice
}

func PartTwo(input string) int {
	matrix := ds.CreateStringMatrix(input)
	ids := GetPlotIdentifiers(matrix)

	totalPrice := 0

	for _, id := range ids {
		for _, plot := range FindPlots(id, matrix) {
			area := plot.ComputeArea()
			sides := plot.ComputeSides(matrix)

			totalPrice += area * sides
		}
	}

	return totalPrice
}

func GetPlotIdentifiers(matrix [][]string) []string {
	identifiers := ds.NewSet[string]()
	for row := range matrix {
		for col := range matrix[0] {
			if matrix[row][col] != "." {
				identifiers.Add(matrix[row][col])
			}
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

func (p ElfPlot) ComputePerimeter(matrix ds.Matrix[string]) []ElfPlotBoundary {
	perimeter := []ElfPlotBoundary{}
	for _, space := range p.land {
		neighbors := ds.GetNeighborsInOrder(space, "RDLU")
		for idx, neighbor := range neighbors {
			if matrix.IsCoordInBounds(neighbor) && matrix[neighbor.Row][neighbor.Col] != p.id {
				perimeter = append(perimeter, ElfPlotBoundary{coord: neighbor, dir: GetDirectionFromIndex(idx)})
			} else if !matrix.IsCoordInBounds(neighbor) {
				perimeter = append(perimeter, ElfPlotBoundary{coord: neighbor, dir: GetDirectionFromIndex(idx)})
			}
		}
	}
	return perimeter
}

type BoundaryGraph struct {
	nodes map[ds.Coordinates][]Direction
}

func BuildBoundaryGraph(perimeter []ElfPlotBoundary) BoundaryGraph {
	graph := BoundaryGraph{nodes: make(map[ds.Coordinates][]Direction)}
	for _, boundary := range perimeter {
		graph.nodes[boundary.coord] = append(graph.nodes[boundary.coord], boundary.dir)
	}
	return graph
}

// i give up
// this one does not work for irregular shapes like "C" in the test input. oh well.
func (p ElfPlot) ComputeSides(matrix ds.Matrix[string]) int {
	sides := 0
	perimiter := p.ComputePerimeter(matrix)

	start := perimiter[0]
	curr := perimiter[0]
	visited := ds.NewSet[ElfPlotBoundary]()

	for shouldContinue := true; shouldContinue; shouldContinue = start != curr {
		if curr.dir == "D" {
			// find left neighbor with dir D
			if exists, neighbor := CheckNeighbors(perimiter, ds.NewCoordinate(curr.coord.Row, curr.coord.Col-1), Down); exists && !visited.Has(*neighbor) {
				curr = *neighbor
				visited.Add(*neighbor)
				continue
			}
			// if that doesn't exist, then let's find an up-left neighbor
			if exists, neighbor := CheckNeighbors(perimiter, ds.NewCoordinate(curr.coord.Row-1, curr.coord.Col-1), Left); exists && !visited.Has(*neighbor) {
				curr = *neighbor
				visited.Add(*neighbor)
				sides++
				continue
			}
			// we have no valid moves, so stop
			break
		} else if curr.dir == "L" {
			// find up neighbor with dir L
			if exists, neighbor := CheckNeighbors(perimiter, ds.NewCoordinate(curr.coord.Row, curr.coord.Col-1), Left); exists && !visited.Has(*neighbor) {
				curr = *neighbor
				visited.Add(*neighbor)
				continue
			}
			// if that doesn't exist, then let's find an up-right neighbor
			if exists, neighbor := CheckNeighbors(perimiter, ds.NewCoordinate(curr.coord.Row-1, curr.coord.Col+1), Up); exists && !visited.Has(*neighbor) {
				curr = *neighbor
				visited.Add(*neighbor)
				sides++
				continue
			}
			// we have no valid moves, so stop
			break
		} else if curr.dir == "U" {
			// find right neighbor with dir U
			if exists, neighbor := CheckNeighbors(perimiter, ds.NewCoordinate(curr.coord.Row, curr.coord.Col+1), Up); exists && !visited.Has(*neighbor) {
				curr = *neighbor
				visited.Add(*neighbor)
				continue
			}
			// if that doesn't exist, then let's find a down-right neighbor
			if exists, neighbor := CheckNeighbors(perimiter, ds.NewCoordinate(curr.coord.Row+1, curr.coord.Col+1), Right); exists && !visited.Has(*neighbor) {
				curr = *neighbor
				visited.Add(*neighbor)
				sides++
				continue
			}
			// we have no valid moves, so stop
			break
		} else if curr.dir == "R" {
			// find down neighbor with dir R
			if exists, neighbor := CheckNeighbors(perimiter, ds.NewCoordinate(curr.coord.Row+1, curr.coord.Col), Right); exists && !visited.Has(*neighbor) {
				curr = *neighbor
				visited.Add(*neighbor)
				continue
			}
			// if that doesn't exist, then let's find a down-left nieghbor
			if exists, neighbor := CheckNeighbors(perimiter, ds.NewCoordinate(curr.coord.Row+1, curr.coord.Col-1), Down); exists && !visited.Has(*neighbor) {
				curr = *neighbor
				visited.Add(*neighbor)
				sides++
				continue
			}
			// we have no valid moves, so stop
			break
		}
	}

	return sides
}

func CheckNeighbors(perimiter []ElfPlotBoundary, coord ds.Coordinates, dir Direction) (bool, *ElfPlotBoundary) {
	for _, boundary := range perimiter {
		if boundary.coord.Row == coord.Row && boundary.coord.Col == coord.Col && boundary.dir == dir {
			return true, &boundary
		}
	}
	return false, nil
}

func GetDirectionFromIndex(idx int) Direction {
	if idx < 0 || idx > 4 {
		panic("GetDirectionFromIndex(idx) expects an index between 1-4")
	}
	return []Direction{"R", "D", "L", "U"}[idx]
}
