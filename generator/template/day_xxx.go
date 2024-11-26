package main

import (
	"fmt"

	"github.com/chancehl/advent-of-code-2024/utils"
)

func main() {
	input, _ := utils.ReadInputFile("./input.txt")

	fmt.Println(dayXXXSolution(input))
}

func dayXXXSolution(lines []string) (int, int) {
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
