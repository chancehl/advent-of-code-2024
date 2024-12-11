package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/chancehl/advent-of-code-2024/ds"
	"github.com/chancehl/advent-of-code-2024/utils/input"
	"github.com/chancehl/advent-of-code-2024/utils/math"
	"github.com/chancehl/advent-of-code-2024/utils/slices"
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

func (m AntennaMap) ComputeAntinodes() map[string][]Coordinates {
	coords := make(map[string][]Coordinates)

	for frequency, antennas := range m.FindAntennae() {
		potentialAntinodes := ds.NewSet[string]()

		for _, pair := range PermuteCoords(antennas) {
			a, b := pair[0], pair[1]

			for row := 0; row < len(m); row++ {
				for col := 0; col < len(m[0]); col++ {
					antinode := Coordinates{row, col}

					distA := int(math.ComputeDistance(a.row, a.col, row, col))
					distB := int(math.ComputeDistance(b.row, b.col, row, col))

					if (distA == 2*distB || distB == 2*distA) && IsInLine(a, b, antinode) {
						key := fmt.Sprintf("%d,%d", row, col)
						potentialAntinodes.Add(key)
					}
				}
			}
		}

		for _, key := range potentialAntinodes.Values() {
			parts := strings.Split(key, ",")
			row, _ := strconv.Atoi(parts[0])
			col, _ := strconv.Atoi(parts[1])
			coords[frequency] = append(coords[frequency], Coordinates{row, col})
		}
	}

	return coords
}

func IsInLine(a, b, c Coordinates) bool {
	// Check same row
	if a.row == b.row && b.row == c.row {
		return true
	}

	// Check same column
	if a.col == b.col && b.col == c.col {
		return true
	}

	// Check forward diagonal
	if (a.row-a.col) == (b.row-b.col) && (b.row-b.col) == (c.row-c.col) {
		return true
	}

	// Check backward diagonal
	if (a.row+a.col) == (b.row+b.col) && (b.row+b.col) == (c.row+c.col) {
		return true
	}

	return false
}

func PermuteCoords(coords []Coordinates) [][]Coordinates {
	permuted := [][]Coordinates{}
	visited := ds.NewSet[string]()
	if len(coords) <= 1 {
		return permuted
	}
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			key := fmt.Sprintf("%d,%d-%d,%d", coords[i].row, coords[i].col, coords[j].row, coords[j].col)
			if !visited.Has(key) {
				visited.Add(key)
				permuted = append(permuted, []Coordinates{coords[i], coords[j]})
			}
		}
	}
	return permuted
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

func IsInSameRow(pair []Coordinates) bool {
	if len(pair) == 0 {
		return false
	}

	if len(pair) == 1 {
		return true
	}

	return slices.All(pair, func(p Coordinates) bool {
		return p.row == pair[0].row
	})
}

func IsInSameCol(pair []Coordinates) bool {
	if len(pair) == 0 {
		return false
	}

	if len(pair) == 1 {
		return true
	}

	return slices.All(pair, func(p Coordinates) bool {
		return p.col == pair[0].col
	})
}

func IsOnSameDiagonal(pair []Coordinates) bool {
	if len(pair) == 0 {
		return false
	}

	if len(pair) == 1 {
		return true
	}

	for i := 0; i < len(pair)-1; i++ {
		a := pair[i]
		b := pair[i+1]

		aSlope := a.col - a.row
		bSlope := b.col - b.row
		if aSlope != bSlope {
			return false
		}
	}
	return true
}

func IsOnSameBackwardsDiagonal(pair []Coordinates) bool {
	if len(pair) == 0 {
		return false
	}

	if len(pair) == 1 {
		return true
	}

	for i := 1; i < len(pair); i++ {
		a := pair[i]
		b := pair[i-1]

		if !(a.row+a.col == b.row+b.col) {
			return false
		}
	}
	return true
}
