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

type ElfSleighLaunchRule struct {
	left  int
	right int
}

type ElfSleighLaunchSafetyManual struct {
	rules   []ElfSleighLaunchRule
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

// For the record this one really pissed me off... It looks like a directed graph problem to start.
// The idea here is that we're wanting to ensure that every node in the graph is positioned in the
// output so that it's connected nodes are always before it. HOWEVER, the input data is structured
// in such a way that a cycle will never occur. It always stops one short. So for now I just brute
// forced this one. The runtime is sub 5ms so I don't really care if this is less than optimized.
func PartOne(input string) int {
	manual := NewElfSafetyManual(input)
	sum := 0

	for _, update := range manual.updates {
		valid := true

		for _, rule := range manual.rules {
			lIndex := slices.Index(update, rule.left)
			rIndex := slices.Index(update, rule.right)

			if lIndex != -1 && rIndex != -1 && lIndex > rIndex {
				valid = false
				break
			}
		}

		if valid {
			sum += update[len(update)/2]
		}
	}

	return sum
}

func PartTwo(input string) int {
	manual := NewElfSafetyManual(input)
	sum := 0

	for _, update := range manual.updates {
		valid := true

		for _, rule := range manual.rules {
			lIndex := slices.Index(update, rule.left)
			rIndex := slices.Index(update, rule.right)

			if lIndex != -1 && rIndex != -1 && lIndex > rIndex {
				valid = false
				break
			}
		}

		if !valid {
			adjancencyList := make(ds.AdjacencyList)

			for _, rule := range manual.findRelevantRules(update) {
				if slices.Contains(update, rule.left) && slices.Contains(update, rule.right) {
					adjancencyList.Insert(rule.left, rule.right)
				}
			}

			sorted := adjancencyList.TopologicalSort()
			sum += sorted[len(sorted)/2]
		}
	}

	return sum
}

func NewElfSafetyManual(input string) ElfSleighLaunchSafetyManual {
	rules := parseRules(input)
	updates := parseUpdates(input)

	return ElfSleighLaunchSafetyManual{rules, updates}
}

func parseRules(input string) []ElfSleighLaunchRule {
	rules := []ElfSleighLaunchRule{}

	parts := strings.Split(input, "\n\n")
	rawRules := strings.Split(parts[0], "\n")

	for _, rule := range rawRules {
		numbers := strings.Split(rule, "|")

		left, _ := strconv.Atoi(numbers[0])
		right, _ := strconv.Atoi(numbers[1])

		rules = append(rules, ElfSleighLaunchRule{left, right})
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

func (manual ElfSleighLaunchSafetyManual) findRelevantRules(update []int) []ElfSleighLaunchRule {
	rules := []ElfSleighLaunchRule{}
	for _, value := range update {
		for _, rule := range manual.rules {
			if value == rule.left {
				rules = append(rules, rule)
			}
		}
	}
	return rules
}
