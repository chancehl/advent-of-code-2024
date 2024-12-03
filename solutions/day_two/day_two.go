package main

import (
	"fmt"
	"log"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/chancehl/advent-of-code-2024/utils/input"
	"github.com/chancehl/advent-of-code-2024/utils/math"
	"github.com/chancehl/advent-of-code-2024/utils/timer"
)

type Report []int

func main() {
	path, err := filepath.Abs("solutions/day_two/input.txt")
	if err != nil {
		log.Fatalf("failed to construct path to input: %v", err)
	}

	input, err := input.Read(path)
	if err != nil {
		log.Fatalf("failed to read %s: %v", path, err)
	}

	dayTwoSolution(input)
}

func dayTwoSolution(input string) (int, int) {
	partOneResult, partOneRuntime := timer.ExecuteTimedFunc(PartOne, input)
	partTwoResult, partTwoRuntime := timer.ExecuteTimedFunc(PartTwo, input)

	fmt.Printf("[day two / part one] result=%d, time=%.2fms\n", partOneResult, partOneRuntime)
	fmt.Printf("[day two / part two] result=%d, time=%.2fms\n", partTwoResult, partTwoRuntime)

	return partOneResult, partTwoResult
}

func PartOne(input string) int {
	safe := 0
	for _, report := range ParseReports(input) {
		if report.IsSafe() {
			safe += 1
		}
	}
	return safe
}

func PartTwo(input string) int {
	safe := 0
	for _, report := range ParseReports(input) {
		if report.IsSafeWithDampener() {
			safe += 1
		}
	}
	return safe
}

func ParseReports(input string) []Report {
	reports := []Report{}

	for _, report := range strings.Split(input, "\n") {
		row := Report{}

		for _, level := range strings.Split(report, " ") {
			value, _ := strconv.Atoi(level)
			row = append(row, value)
		}

		reports = append(reports, row)
	}

	return reports
}

func (report Report) IsSafe() bool {
	first := report[0]
	last := report[len(report)-1]

	for i := 1; i < len(report); i += 1 {
		left := report[i-1]
		right := report[i]

		// diff check
		if math.Abs(right, left) > 3 {
			return false
		}

		// order check
		if first < last {
			if !(left < right) {
				return false
			}
		} else {
			if !(left > right) {
				return false
			}
		}
	}

	return true
}

func (report Report) IsSafeWithDampener() bool {
	for index := range report {
		left := report[0:index]
		right := report[index+1:]

		subreport := slices.Concat(left, right)

		if subreport.IsSafe() {
			return true
		}
	}

	return false
}
