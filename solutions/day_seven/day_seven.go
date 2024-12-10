package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/chancehl/advent-of-code-2024/utils/input"
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

	equations := ParseEquationsFromInput(input)

	for _, equation := range equations {
		if IsValidEquation(equation, true) {
			validEquations = append(validEquations, equation)
		}
	}

	for _, equation := range validEquations {
		result += equation.target
	}

	return result
}

func PartTwo(input string) int {
	validEquations := []ElfEquation{}
	result := 0

	equations := ParseEquationsFromInput(input)

	for _, equation := range equations {
		if IsValidEquation(equation, false) {
			validEquations = append(validEquations, equation)
		}
	}

	for _, equation := range validEquations {
		result += equation.target
	}

	return result
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

func IsValidEquation(e ElfEquation, simplified bool) bool {
	var compute func(int, []int) bool

	compute = func(acc int, rest []int) bool {
		if len(rest) == 0 {
			return acc == e.target
		}

		left := rest[0]

		if !simplified {
			concatted, _ := strconv.Atoi(fmt.Sprintf("%d%d", acc, left))
			return compute(acc*left, rest[1:]) || compute(acc+left, rest[1:]) || compute(concatted, rest[1:])
		}

		return compute(acc*left, rest[1:]) || compute(acc+left, rest[1:])
	}

	return compute(e.operands[0], e.operands[1:])
}
