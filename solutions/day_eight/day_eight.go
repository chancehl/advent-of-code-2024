package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/chancehl/advent-of-code-2024/utils/input"
	"github.com/chancehl/advent-of-code-2024/utils/timer"
)

type Coordinates struct {
	row int
	col int
}

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
	antennaMap := CreateAntennaMapFromInput(input)
	return len(antennaMap.ComputeAntinodes())
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

func (m AntennaMap) ComputeAntinodes() map[string][]Coordinates {
	coords := make(map[string][]Coordinates)
	antennae := m.FindAntennae()

	for frequency := range antennae {
		for _, coord := range antennae[frequency] {
			fmt.Println(frequency, coord)
		}
	}
	return coords
}

func (m AntennaMap) FindAntennae() map[string][]Coordinates {
	coordinateMap := make(map[string][]Coordinates)
	for row := range m {
		for col := range m {
			if m[row][col] != "." {
				key := m[row][col]
				current := coordinateMap[key]
				current = append(current, Coordinates{row, col})
				coordinateMap[key] = current
			}
		}
	}
	return coordinateMap
}
