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

type ElfReport []int

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

	fmt.Printf("[day two / part one] result=%d, time=%d ms\n", partOneResult, partOneRuntime)
	fmt.Printf("[day two / part two] result=%d, time=%d ms\n", partTwoResult, partTwoRuntime)

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

		data = append(data, row)
	}

	return data
}

func (report ElfReport) IsSafe() bool {
	first := report[0]
	last := report[len(report)-1]

	for i := 1; i < len(report); i += 1 {
		left := report[i-1]
		right := report[i]

		if math.Abs(right, left) > 3 { // diff check
			return false
		} else if first < last {
			// we are ascending
			if !(left < right) {
				return false
			}
		} else {
			// we are descending
			if !(left > right) {
				return false
			}
		}
	}

	return true
}

func (report ElfReport) IsSafeWithDampener() bool {
	var currentReport ElfReport

	for index := range report {
		if index == 0 {
			currentReport = report[index+1:]
		} else if index == len(report)-1 {
			currentReport = report[0 : len(report)-1]
		} else {
			left := report[0:index]
			right := report[index+1:]
			currentReport = slices.Concat(left, right)
		}

		if currentReport.IsSafe() {
			return true
		}
	}

	return false
}
