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

type ElfReport struct {
	levels []int
}

type ElfReportData []ElfReport

type ElfReportSaftetyChecker interface {
	IsSafe() bool
	IsSafeWithDampener() bool
}

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
	valid := 0
	for _, report := range NewElfReportData(input) {
		if report.IsSafe() {
			valid += 1
		}
	}
	return valid
}

func PartTwo(input string) int {
	valid := 0
	for _, report := range NewElfReportData(input) {
		if report.IsSafeWithDampener() {
			valid += 1
		}
	}
	return valid
}

func NewElfReportData(input string) ElfReportData {
	data := ElfReportData{}

	for _, report := range strings.Split(input, "\n") {
		row := []int{}

		for _, level := range strings.Split(report, " ") {
			value, _ := strconv.Atoi(level)
			row = append(row, value)
		}

		data = append(data, ElfReport{levels: row})
	}

	return data
}

func (r ElfReport) IsSafe() bool {
	levels := r.levels

	first := levels[0]
	last := levels[len(levels)-1]

	for i := 1; i < len(levels); i += 1 {
		left := levels[i-1]
		right := levels[i]

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

func (r ElfReport) IsSafeWithDampener() bool {
	for index := range r.levels {
		left := r.levels[0:index]
		right := r.levels[index+1:]

		subreport := ElfReport{levels: slices.Concat(left, right)}

		if subreport.IsSafe() {
			return true
		}
	}

	return false
}
