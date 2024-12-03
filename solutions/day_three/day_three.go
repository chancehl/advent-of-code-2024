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

type Signal int

const (
	Do Signal = iota
	Dont
)

type Operation struct {
	instruction string
	start       int
	end         int
	signal      Signal
}

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
	operations := findOperations(input, true)
	for _, operation := range operations {
		if operation.signal == Do {
			sum += operation.execute()
		}
	}
	return sum
}

func PartTwo(input string) int {
	sum := 0
	operations := findOperations(input, false)
	for _, operation := range operations {
		if operation.signal == Do {
			sum += operation.execute()
		}
	}
	return sum
}

func (o Operation) execute() int {
	left, right := o.parseOperands()
	return left * right
}

func (o Operation) parseOperands() (int, int) {
	stripped := strings.ReplaceAll(o.instruction, "mul", "")
	stripped = strings.ReplaceAll(stripped, "(", "")
	stripped = strings.ReplaceAll(stripped, ")", "")

	values := strings.Split(stripped, ",")

	left, _ := strconv.Atoi(values[0])
	right, _ := strconv.Atoi(values[1])

	return left, right
}

func findOperations(code string, ignoreSignals bool) []Operation {
	operations := []Operation{}

	r, _ := regexp.Compile(`mul\(\d+,\d+\)`)
	matches := r.FindAllStringIndex(code, -1)

	for _, match := range matches {
		start := match[0]
		end := match[1]

		var signal Signal

		if ignoreSignals {
			signal = Do
		} else {
			signal = findLatestSignal(code[0:start])
		}

		operation := Operation{
			start:       start,
			end:         end,
			signal:      signal,
			instruction: code[start:end],
		}

		operations = append(operations, operation)
	}

	return operations
}

func findLatestSignal(code string) Signal {
	doRegex, _ := regexp.Compile(`do\(\)`)
	dontRegex, _ := regexp.Compile(`don't\(\)`)

	doMatches := doRegex.FindAllStringIndex(code, -1)
	dontMatches := dontRegex.FindAllStringIndex(code, -1)

	if len(doMatches) > 0 && len(dontMatches) > 0 {
		lastDoRange := doMatches[len(doMatches)-1]
		lastDontRange := dontMatches[len(dontMatches)-1]

		if lastDoRange[1] > lastDontRange[1] {
			return Do
		} else {
			return Dont
		}
	}

	if len(dontMatches) > 0 {
		return Dont
	}

	return Do
}
