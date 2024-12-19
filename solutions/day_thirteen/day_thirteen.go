package main

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/chancehl/advent-of-code-2024/utils/input"
	mathUtils "github.com/chancehl/advent-of-code-2024/utils/math"
	"github.com/chancehl/advent-of-code-2024/utils/timer"
)

type ButtonConfig struct {
	xMovement int64
	yMovement int64
}

type PrizeLocation struct {
	x int64
	y int64
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

func dayThirteenSolution(input string) (int64, int64) {
	partOneResult, partOneRuntime := timer.ExecuteTimedFunc(PartOne, input)
	partTwoResult, partTwoRuntime := timer.ExecuteTimedFunc(PartTwo, input)

	fmt.Printf("[day thirteen / part one] result=%d, time=%.2fms\n", partOneResult, partOneRuntime)
	fmt.Printf("[day thirteen / part two] result=%d, time=%.2fms\n", partTwoResult, partTwoRuntime)

	return partOneResult, partTwoResult
}

func PartOne(input string) int64 {
	sum := int64(0)
	for _, config := range ParseMachineConfigs(input) {
		costs := []int64{}
		for aPress := int64(0); aPress <= 100; aPress++ {
			for bPress := int64(0); bPress <= 100; bPress++ {
				if IsValidCombination(aPress, bPress, config) {
					costs = append(costs, (aPress*3)+bPress)
				}
			}
		}
		if len(costs) > 0 {
			sum += slices.Min(costs)
		}
	}
	return sum
}

func PartTwo(input string) int64 {
	sum := int64(0)
	return sum
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
				config.a = ButtonConfig{xMovement: int64(x), yMovement: int64(y)}
			} else if idx == 1 {
				config.b = ButtonConfig{xMovement: int64(x), yMovement: int64(y)}
			} else {
				config.prize = PrizeLocation{x: int64(x), y: int64(y)}
			}
		}
		configs = append(configs, config)

	}
	return configs
}

func IsValidCombination(aPress, bPress int64, config MachineConfig) bool {
	return (config.a.xMovement*aPress)+(config.b.xMovement*bPress) == config.prize.x && (config.a.yMovement*aPress)+(config.b.yMovement*bPress) == config.prize.y
}

func DoesSolutionExist(config MachineConfig) (bool, int64, int64) {
	// Compute gcd for X and Y movements
	gcdX := mathUtils.GCD(int64(config.a.xMovement), int64(config.b.xMovement))
	gcdY := mathUtils.GCD(int64(config.a.yMovement), int64(config.b.yMovement))

	// Check individual feasibility
	if config.prize.x%gcdX != 0 || config.prize.y%gcdY != 0 {
		return false, -1, -1
	}

	return false, -1, -1
}
