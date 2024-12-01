package main

import (
	"log"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/chancehl/advent-of-code-2024/utils"
)

func main() {
	path, err := filepath.Abs("solutions/day_one/input.txt")
	if err != nil {
		log.Fatalf("failed to construct path to input: %v", err)
	}

	input, err := utils.ReadInput(path)
	if err != nil {
		log.Fatalf("failed to read %s: %v", path, err)
	}

	dayOneSolution(input)
}

func dayOneSolution(input string) (int, int) {
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
	leftNums, rightNums := parsePairs(input)

	diff := 0
	index := 0

	for index < len(leftNums) {
		left := leftNums[index]
		right := rightNums[index]

		if right > left {
			diff += (right - left)
		} else {
			diff += (left - right)
		}

		index += 1
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
