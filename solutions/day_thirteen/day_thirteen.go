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
	"golang.org/x/exp/slices"
)

type ButtonConfig struct {
	x int
	y int
}

type PrizeLocation struct {
	x int
	y int
}

type MachineConfig struct {
	a     ButtonConfig
	b     ButtonConfig
	prize PrizeLocation
}

func main() {
	path, err := filepath.Abs("solutions/day_thirteen/input.txt")
	if err != nil {
		log.Fatalf("failed to construct path to input: %v", err)
	}

	input, err := input.Read(path)
	if err != nil {
		log.Fatalf("failed to read %s: %v", path, err)
	}

	dayThirteenSolution(input)
}

func dayThirteenSolution(input string) (int, int) {
	partOneResult, partOneRuntime := timer.ExecuteTimedFunc(PartOne, input)
	partTwoResult, partTwoRuntime := timer.ExecuteTimedFunc(PartTwo, input)

	fmt.Printf("[day thirteen / part one] result=%d, time=%.2fms\n", partOneResult, partOneRuntime)
	fmt.Printf("[day thirteen / part two] result=%d, time=%.2fms\n", partTwoResult, partTwoRuntime)

	return partOneResult, partTwoResult
}

func PartOne(input string) int {
	configs := ParseMachineConfigs(input)
	sum := 0

	for _, config := range configs {
		minCost := FindMinimumSolution(config)
		if minCost != -1 {
			sum += minCost
		}
	}

	return sum
}

func PartTwo(input string) int {
	return -1
}

func FindMinimumSolution(config MachineConfig) int {
	buttonCombinations := [][]int{}

	for aPress := 0; aPress <= 100; aPress++ {
		for bPress := 0; bPress <= 100; bPress++ {
			if ((config.a.x*aPress)+(config.b.x*bPress) == config.prize.x) && ((config.a.y*aPress)+(config.b.y*bPress) == config.prize.y) {
				buttonCombinations = append(buttonCombinations, []int{aPress, bPress})
			}
		}
	}

	costs := []int{}

	for _, combination := range buttonCombinations {
		a := combination[0]
		b := combination[1]
		costs = append(costs, (a*3)+(b*1))
	}

	if len(costs) > 0 {
		return slices.Min(costs)
	} else {
		return -1
	}
}

func ParseMachineConfigs(input string) []MachineConfig {
	configs := []MachineConfig{}
	for _, section := range strings.Split(input, "\n\n") {
		config := MachineConfig{}
		for idx, line := range strings.Split(section, "\n") {
			digitRegex, _ := regexp.Compile(`\d+`)
			digitStrings := digitRegex.FindAllString(line, -1)

			x, _ := strconv.Atoi(digitStrings[0])
			y, _ := strconv.Atoi(digitStrings[1])

			if idx == 0 {
				config.a = ButtonConfig{x, y}
			} else if idx == 1 {
				config.b = ButtonConfig{x, y}
			} else {
				config.prize = PrizeLocation{x, y}
			}
		}
		configs = append(configs, config)

	}
	return configs
}
