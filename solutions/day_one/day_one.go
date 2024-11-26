package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/chancehl/advent-of-code-2024/utils"
)

func main() {
	path, err := filepath.Abs("solutions/day_one/input.txt")
	if err != nil {
		log.Fatalf("failed to construct path to input: %v", err)
	}

	input, err := utils.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read %s: %v", path, err)
	}

	fmt.Println(dayOneSolution(input))
}

func dayOneSolution(lines []string) (int, int) {
	one := PartOne(lines)
	two := PartTwo(lines)

	return one, two
}

func PartOne(lines []string) int {
	return -1
}

func PartTwo(lines []string) int {
	return -1
}
