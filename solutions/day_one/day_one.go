package main

import (
	"fmt"
	"log"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/chancehl/advent-of-code-2024/utils/input"
	"github.com/chancehl/advent-of-code-2024/utils/math"
	"github.com/chancehl/advent-of-code-2024/utils/timer"
)

func main() {
	path, err := filepath.Abs("solutions/day_one/input.txt")
	if err != nil {
		log.Fatalf("failed to construct path to input: %v", err)
	}

	input, err := input.Read(path)
	if err != nil {
		log.Fatalf("failed to read %s: %v", path, err)
	}

	dayOneSolution(input)
}

func dayOneSolution(input string) (int, int) {
	partOneResult, partOneRuntime := timer.ExecuteTimedFunc(PartOne, input)
	partTwoResult, partTwoRuntime := timer.ExecuteTimedFunc(PartTwo, input)

	fmt.Printf("[day one / part one] result=%d, time=%d ms\n", partOneResult, partOneRuntime)
	fmt.Printf("[day one / part two] result=%d, time=%d ms\n", partTwoResult, partTwoRuntime)

	return partOneResult, partTwoResult
}

func PartOne(input string) int {
	leftNums, rightNums := parsePairs(input)

	diff := 0

	for index := range leftNums {
		diff += math.Abs(leftNums[index], rightNums[index])
	}

	return diff
}

func PartTwo(input string) int {
	similarity := 0

	leftValues, rightValues := parsePairs(input)

	counts := make(map[int]int)

	for _, value := range rightValues {
		current := counts[value]
		counts[value] = current + 1
	}

	for _, value := range leftValues {
		similarity += (value * counts[value])
	}

	return similarity
}

func parsePairs(input string) ([]int, []int) {
	leftValues := []int{}
	rightValues := []int{}

	for _, line := range strings.Split(input, "\n") {
		values := strings.Split(line, "   ")

		leftInt, _ := strconv.Atoi(values[0])
		leftValues = append(leftValues, leftInt)

		rightInt, _ := strconv.Atoi(values[1])
		rightValues = append(rightValues, rightInt)
	}

	sort.Ints(leftValues)
	sort.Ints(rightValues)

	return leftValues, rightValues
}
