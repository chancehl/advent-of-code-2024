package main

import (
	"fmt"

	"github.com/chancehl/advent-of-code-2024/utils"
)

func main() {
	lines, _ := utils.ReadInputFile("input.txt")

	for i, line := range lines {
		fmt.Printf("Line %d: %s\n", i, line)
	}
}
