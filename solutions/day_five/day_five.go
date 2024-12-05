package main

import (
	"fmt"
	"log"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/chancehl/advent-of-code-2024/ds"
	"github.com/chancehl/advent-of-code-2024/utils/input"
	"github.com/chancehl/advent-of-code-2024/utils/timer"
)

type ElfSleighLaunchSafetyManual struct {
	rules   ds.AdjacencyList
	updates [][]int
}

func main() {
	path, err := filepath.Abs("solutions/day_five/input.txt")
	if err != nil {
		log.Fatalf("failed to construct path to input: %v", err)
	}

	input, err := input.Read(path)
	if err != nil {
		log.Fatalf("failed to read %s: %v", path, err)
	}

	dayFiveSolution(input)
}

func dayFiveSolution(input string) (int, int) {
	partOneResult, partOneRuntime := timer.ExecuteTimedFunc(PartOne, input)
	partTwoResult, partTwoRuntime := timer.ExecuteTimedFunc(PartTwo, input)

	fmt.Printf("[day five / part one] result=%d, time=%.2fms\n", partOneResult, partOneRuntime)
	fmt.Printf("[day five / part two] result=%d, time=%.2fms\n", partTwoResult, partTwoRuntime)

	return partOneResult, partTwoResult
}

func PartOne(input string) int {
	manual := parseSafetyManual(input)
	sum := 0

	for _, update := range manual.updates {
		valid := true
		for _, value := range update {
			for _, connected := range manual.rules.Get(value) {
				if slices.Contains(update, value) && slices.Contains(update, connected) {
					leftIndex := slices.Index(update, value)
					rightIndex := slices.Index(update, connected)

					if leftIndex > rightIndex {
						valid = false
						break
					}
				}
			}
		}

		if valid {
			sum += update[len(update)/2]
		}
	}

	return sum
}

func PartTwo(input string) int {
	return -1
}

func parseSafetyManual(input string) ElfSleighLaunchSafetyManual {
	rules := parseRules(input)
	updates := parseUpdates(input)

	return ElfSleighLaunchSafetyManual{rules, updates}
}

func parseRules(input string) ds.AdjacencyList {
	rules := make(ds.AdjacencyList)

	parts := strings.Split(input, "\n\n")
	rawRules := strings.Split(parts[0], "\n")

	for _, rule := range rawRules {
		numbers := strings.Split(rule, "|")

		left, _ := strconv.Atoi(numbers[0])
		right, _ := strconv.Atoi(numbers[1])

		if err := rules.Insert(left, right); err != nil {
			panic("invalid input")
		}
	}

	return rules
}

func parseUpdates(input string) [][]int {
	updates := [][]int{}

	parts := strings.Split(input, "\n\n")
	rawUpdates := strings.Split(parts[1], "\n")

	for _, update := range rawUpdates {
		intValues := []int{}

		for _, strValue := range strings.Split(update, ",") {
			value, _ := strconv.Atoi(strValue)
			intValues = append(intValues, value)
		}

		updates = append(updates, intValues)
	}

	return updates
}
