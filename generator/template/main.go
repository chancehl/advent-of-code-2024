package main

import (
	"log"
	"path/filepath"
	"time"

	"github.com/chancehl/advent-of-code-2024/utils"
)

func main() {
	path, err := filepath.Abs("solutions/day_xxx/input.txt")
	if err != nil {
		log.Fatalf("failed to construct path to input: %v", err)
	}

	input, err := utils.ReadInput(path)
	if err != nil {
		log.Fatalf("failed to read %s: %v", path, err)
	}

	dayXXXSolution(input)
}

func dayXXXSolution(input string) (int, int) {
	start := time.Now()
	one := PartOne(input)
	elapsed := time.Since(start)
	resultOne := utils.SolutionResult{Result: one, Time: elapsed.Milliseconds()}

	start = time.Now()
	two := PartTwo(input)
	elapsed = time.Since(start)
	resultTwo := utils.SolutionResult{Result: two, Time: elapsed.Milliseconds()}

	utils.PrintAdventResults(resultOne, resultTwo)

	return one, two
}

func PartOne(input string) int {
	return -1
}

func PartTwo(input string) int {
	return -1
}
