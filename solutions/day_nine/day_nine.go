package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/chancehl/advent-of-code-2024/utils/input"
	"github.com/chancehl/advent-of-code-2024/utils/timer"
)

func main() {
	path, err := filepath.Abs("solutions/day_nine/input.txt")
	if err != nil {
		log.Fatalf("failed to construct path to input: %v", err)
	}

	input, err := input.Read(path)
	if err != nil {
		log.Fatalf("failed to read %s: %v", path, err)
	}

	dayNineSolution(input)
}

func dayNineSolution(input string) (int, int) {
	partOneResult, partOneRuntime := timer.ExecuteTimedFunc(PartOne, input)
	partTwoResult, partTwoRuntime := timer.ExecuteTimedFunc(PartTwo, input)

	fmt.Printf("[day nine / part one] result=%d, time=%.2fms\n", partOneResult, partOneRuntime)
	fmt.Printf("[day nine / part two] result=%d, time=%.2fms\n", partTwoResult, partTwoRuntime)

	return partOneResult, partTwoResult
}

func PartOne(input string) int {
	dm := CreateExpandedDiskMap(input)
	expandedDm := MoveFiles(dm)
	checksum := CalculateChecksum(expandedDm)
	return checksum
}

func PartTwo(input string) int {
	return -1
}

func CreateExpandedDiskMap(input string) string {
	fileId := 0
	diskmap := []string{}

	for charIndex, char := range strings.Split(input, "") {
		num, _ := strconv.Atoi(char)

		if charIndex%2 == 0 {
			for i := 0; i < num; i++ {
				diskmap = append(diskmap, strconv.Itoa(fileId))
			}
			fileId += 1
		} else {
			for i := 0; i < num; i++ {
				diskmap = append(diskmap, ".")
			}
		}
	}
	return strings.Join(diskmap, "")
}

func MoveFiles(dm string) string {
	chars := strings.Split(dm, "")

	left := 0
	right := len(chars) - 1

	for left < right {
		if chars[left] == "." && chars[right] != "." {
			chars[left] = chars[right]
			chars[right] = "."
		} else if chars[left] != "." {
			left += 1
		} else if chars[right] == "." {
			right -= 1
		}
	}

	return strings.Join(chars, "")
}

func CalculateChecksum(dm string) int {
	checksum := 0
	chars := strings.Split(dm, "")
	for i := range chars {
		num, _ := strconv.Atoi(chars[i])
		checksum += (num * i)
	}
	return checksum
}
