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

func dayElevenSolution(input string) (int, int) {
	partOneResult, partOneRuntime := timer.ExecuteTimedFunc(PartOne, input)
	partTwoResult, partTwoRuntime := timer.ExecuteTimedFunc(PartTwo, input)

	fmt.Printf("[day eleven / part one] result=%d, time=%.2fms\n", partOneResult, partOneRuntime)
	fmt.Printf("[day eleven / part two] result=%d, time=%.2fms\n", partTwoResult, partTwoRuntime)

	return partOneResult, partTwoResult
}

func PartOne(input string) int {
	stones := ParseStones(input)
	for i := 0; i < NumberOfBlinksP1; i++ {
		stones = Blink(stones)
	}
	return len(stones)
}

func PartTwo(input string) int {
	return -1
}

func ParseStones(input string) []int {
	stones := []int{}
	for _, s := range strings.Split(input, " ") {
		num, _ := strconv.Atoi(s)
		stones = append(stones, num)
	}
	return stones
}

func Blink(stones []int) []int {
	updated := []int{}
	for _, stone := range stones {
		updated = append(updated, BlinkStone(stone)...)
	}
	return updated
}

func BlinkStone(stone int) []int {
	updated := []int{}
	if stone == 0 {
		updated = append(updated, 1)
	} else if math.CountDigits(stone)%2 == 0 {
		left, right := math.SplitNumber(stone)

		updated = append(updated, left)
		updated = append(updated, right)
	} else {
		updated = append(updated, stone*2024)
	}
	return updated
}
