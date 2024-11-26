package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/chancehl/advent-of-code-2024/utils"
)

func main() {
	path, err := filepath.Abs("solutions/day_xxx/input.txt")
	if err != nil {
		log.Fatalf("failed to construct path to input: %v", err)
	}

	input, err := utils.ReadInput(path)
	if err != nil {
		log.Fatalf("failed to read %s: %v", path, err)
	}

	fmt.Println(dayXXXSolution(input))
}

func dayXXXSolution(input string) (int, int) {
	one := PartOne(input)
	two := PartTwo(input)

	return one, two
}

func PartOne(input string) int {
	return -1
}

func PartTwo(input string) int {
	return -1
}
