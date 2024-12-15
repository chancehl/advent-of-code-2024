package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/chancehl/advent-of-code-2024/utils/input"
	"github.com/chancehl/advent-of-code-2024/utils/math"
	"github.com/chancehl/advent-of-code-2024/utils/timer"
)

const NumberOfBlinksP1 = 25
const NumberOfBlinksP2 = 75

func main() {
	path, err := filepath.Abs("solutions/day_eleven/input.txt")
	if err != nil {
		log.Fatalf("failed to construct path to input: %v", err)
	}

	input, err := input.Read(path)
	if err != nil {
		log.Fatalf("failed to read %s: %v", path, err)
	}

	dayElevenSolution(input)
}

func dayElevenSolution(input string) (int64, int64) {
	partOneResult, partOneRuntime := timer.ExecuteTimedFunc(PartOne, input)
	partTwoResult, partTwoRuntime := timer.ExecuteTimedFunc(PartTwo, input)

	fmt.Printf("[day eleven / part one] result=%d, time=%.2fms\n", partOneResult, partOneRuntime)
	fmt.Printf("[day eleven / part two] result=%d, time=%.2fms\n", partTwoResult, partTwoRuntime)

	return partOneResult, partTwoResult
}

func PartOne(input string) int64 {
	var sum int64 = 0
	for _, stone := range ParseStones(input) {
		sum += Blink(stone, NumberOfBlinksP1)
	}
	return sum
}

func PartTwo(input string) int64 {
	var sum int64 = 0
	for _, stone := range ParseStones(input) {
		sum += Blink(stone, NumberOfBlinksP2)
	}
	return sum
}

func ParseStones(input string) []int {
	stones := []int{}
	for _, s := range strings.Split(input, " ") {
		num, _ := strconv.Atoi(s)
		stones = append(stones, num)
	}
	return stones
}

var memo = make(map[string]int64)

func Blink(stone int, blinks int) int64 {
	key := fmt.Sprintf("%d:%d", stone, blinks)

	if count, exists := memo[key]; exists {
		return count
	}

	var result int64

	if blinks == 0 {
		result = 1
	} else if stone == 0 {
		result = Blink(1, blinks-1)
	} else if math.CountDigits(stone)%2 == 0 {
		left, right := math.SplitNumber(stone)
		result = Blink(left, blinks-1) + Blink(right, blinks-1)
	} else {
		result = Blink(stone*2024, blinks-1)
	}

	memo[key] = result
	return result
}
