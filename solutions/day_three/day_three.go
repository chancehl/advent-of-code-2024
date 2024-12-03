package main

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/chancehl/advent-of-code-2024/utils/input"
	"github.com/chancehl/advent-of-code-2024/utils/timer"
)

func main() {
	path, err := filepath.Abs("solutions/day_three/input.txt")
	if err != nil {
		log.Fatalf("failed to construct path to input: %v", err)
	}

	input, err := input.Read(path)
	if err != nil {
		log.Fatalf("failed to read %s: %v", path, err)
	}

	dayThreeSolution(input)
}

func dayThreeSolution(input string) (int, int) {
	partOneResult, partOneRuntime := timer.ExecuteTimedFunc(PartOne, input)
	partTwoResult, partTwoRuntime := timer.ExecuteTimedFunc(PartTwo, input)

	fmt.Printf("[day three / part one] result=%d, time=%.2fms\n", partOneResult, partOneRuntime)
	fmt.Printf("[day three / part two] result=%d, time=%.2fms\n", partTwoResult, partTwoRuntime)

	return partOneResult, partTwoResult
}

func PartOne(input string) int {
	sum := 0
	for _, instruction := range findInstructions(input) {
		left, right := parseOperands(instruction)
		sum += (left * right)
	}
	return sum
}

func PartTwo(input string) int {
	sum := 0

	return sum
}

func parseOperands(instruction string) (int, int) {
	stripped := strings.ReplaceAll(instruction, "mul", "")
	stripped = strings.ReplaceAll(stripped, "(", "")
	stripped = strings.ReplaceAll(stripped, ")", "")

	values := strings.Split(stripped, ",")

	left, _ := strconv.Atoi(values[0])
	right, _ := strconv.Atoi(values[1])

	return left, right
}

func findInstructions(code string) []string {
	instructions := []string{}
	r, _ := regexp.Compile(`mul\(\d+,\d+\)`)
	matches := r.FindAll([]byte(code), -1)
	for _, match := range matches {
		instructions = append(instructions, string(match))
	}
	return instructions
}

func FindStopExecutionSignal(code string) int {
	r, _ := regexp.Compile(`don't\(\)`)
	posn := r.FindStringIndex(code)
	if posn == nil {
		return -1
	}
	return posn[1]
}

func FindStartExecutionSignal(code string) int {
	r, _ := regexp.Compile(`do\(\)`)
	posn := r.FindStringIndex(code)
	if posn == nil {
		return -1
	}
	return posn[1]
}
