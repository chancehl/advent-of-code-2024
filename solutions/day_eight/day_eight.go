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

type AntennaMap [][]string

func main() {
	path, err := filepath.Abs("solutions/day_eight/input.txt")
	if err != nil {
		log.Fatalf("failed to construct path to input: %v", err)
	}

	input, err := input.Read(path)
	if err != nil {
		log.Fatalf("failed to read %s: %v", path, err)
	}

	dayEightSolution(input)
}

func dayEightSolution(input string) (int, int) {
	partOneResult, partOneRuntime := timer.ExecuteTimedFunc(PartOne, input)
	partTwoResult, partTwoRuntime := timer.ExecuteTimedFunc(PartTwo, input)

	fmt.Printf("[day eight / part one] result=%d, time=%.2fms\n", partOneResult, partOneRuntime)
	fmt.Printf("[day eight / part two] result=%d, time=%.2fms\n", partTwoResult, partTwoRuntime)

	return partOneResult, partTwoResult
}

func PartOne(input string) int {
	return -1
}

func PartTwo(input string) int {
	return -1
}

func CreateAntennaMapFromInput(input string) AntennaMap {
	matrix := AntennaMap{}
	for _, line := range strings.Split(input, "\n") {
		matrix = append(matrix, strings.Split(line, ""))
	}
	return matrix
}

func PermuteCoords(coords []ds.Coordinates) [][]ds.Coordinates {
	permuted := [][]ds.Coordinates{}
	visited := ds.NewSet[string]()
	if len(coords) <= 1 {
		return permuted
	}
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			key := fmt.Sprintf("%d,%d-%d,%d", coords[i].Row, coords[i].Col, coords[j].Row, coords[j].Col)
			if !visited.Has(key) {
				visited.Add(key)
				permuted = append(permuted, []ds.Coordinates{coords[i], coords[j]})
			}
		}
	}
	return permuted
}

func (m AntennaMap) FindAntennae() map[string][]ds.Coordinates {
	coordinateMap := make(map[string][]ds.Coordinates)
	for row := range m {
		for col := range m {
			if m[row][col] != "." {
				key := m[row][col]
				current := coordinateMap[key]
				current = append(current, ds.Coordinates{Row: row, Col: col})
				coordinateMap[key] = current
			}
		}
	}
	return coordinateMap
}
