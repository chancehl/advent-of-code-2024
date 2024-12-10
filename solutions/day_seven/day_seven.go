package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/chancehl/advent-of-code-2024/utils/input"
	"github.com/chancehl/advent-of-code-2024/utils/slices"
	"github.com/chancehl/advent-of-code-2024/utils/timer"
)

type ElfEquation struct {
	target   int
	operands []int
}

func main() {
	path, err := filepath.Abs("solutions/day_seven/input.txt")
	if err != nil {
		log.Fatalf("failed to construct path to input: %v", err)
	}

	input, err := input.Read(path)
	if err != nil {
		log.Fatalf("failed to read %s: %v", path, err)
	}

	daySevenSolution(input)
}

func daySevenSolution(input string) (int, int) {
	partOneResult, partOneRuntime := timer.ExecuteTimedFunc(PartOne, input)
	partTwoResult, partTwoRuntime := timer.ExecuteTimedFunc(PartTwo, input)

	fmt.Printf("[day seven / part one] result=%d, time=%.2fms\n", partOneResult, partOneRuntime)
	fmt.Printf("[day seven / part two] result=%d, time=%.2fms\n", partTwoResult, partTwoRuntime)

	return partOneResult, partTwoResult
}

func PartOne(input string) int {
	validEquations := []ElfEquation{}
	result := 0

	// equations := ParseEquationsFromInput(input)

	// for _, equation := range equations {
	// 	if IsValidEquation(equation.target, equation.operands) {
	// 		validEquations = append(validEquations, equation)
	// 	}
	// }

	for _, equation := range validEquations {
		result += equation.target
	}

	return -1
}

func PartTwo(input string) int {
	return -1
}

func ParseEquationsFromInput(input string) []ElfEquation {
	equations := []ElfEquation{}

	for _, row := range strings.Split(input, "\n") {
		parts := strings.Split(row, ":")

		target, _ := strconv.Atoi(parts[0])
		rest := strings.Trim(parts[1], " ")

		operands := []int{}

		for _, operand := range strings.Split(rest, " ") {
			value, _ := strconv.Atoi(operand)
			operands = append(operands, value)
		}

		equations = append(equations, ElfEquation{target, operands})
	}

	return equations
}

func IsValidEquation(target int, operands []int) bool {
	permutations := GenerateOperationPermutations(len(operands) - 1)

	filteredPermutations := slices.Filter(permutations, func(s string) bool {
		return len(s) == len(operands)-1
	})

	fmt.Println(filteredPermutations)

	return false
}

func GenerateOperationPermutations(n int) []string {
	permutations := []string{"*", "+"}

	var permute func(int, *[]string) *[]string

	permute = func(i int, visited *[]string) *[]string {
		if i == 0 {
			return visited
		}

		for _, existing := range *visited {
			*visited = append(*visited, existing+"*")
			*visited = append(*visited, existing+"+")
		}

		return permute(i-1, visited)
	}

	return *permute(n, &permutations)
}
