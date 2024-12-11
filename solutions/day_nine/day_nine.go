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

const Empty = -1

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
	return CalculateChecksum(expandedDm)
}

func PartTwo(input string) int {
	return -1
}

func CreateExpandedDiskMap(input string) []int {
	diskmap := []int{}

	for i, c := range strings.Split(input, "") {
		num, _ := strconv.Atoi(c)

		if i%2 == 0 {
			for j := 0; j < num; j++ {
				diskmap = append(diskmap, i/2)
			}
		} else {
			for j := 0; j < num; j++ {
				diskmap = append(diskmap, Empty)
			}
		}
	}

	return diskmap
}

func MoveFiles(dm []int) []int {
	left := 0
	right := len(dm) - 1

	for left < right {
		if dm[left] == Empty && dm[right] != Empty {
			dm[left], dm[right] = dm[right], dm[left]
			left++
			right--
		} else if dm[left] != Empty {
			left++
		} else if dm[right] == Empty {
			right -= 1
		}
	}

	return dm
}

func CalculateChecksum(dm []int) int {
	checksum := 0
	for i, n := range dm {
		if n != Empty {
			checksum += (n * i)
		}
	}
	return checksum
}

func ConvertDiskmapToString(dm []int) string {
	chars := []string{}
	for _, num := range dm {
		if num == -1 {
			chars = append(chars, ".")
		} else {
			chars = append(chars, strconv.Itoa(num))
		}
	}
	return strings.Join(chars, "")
}

func ConvertStringToDiskmap(dm string) []int {
	nums := []int{}
	for _, char := range strings.Split(dm, "") {
		if char == "." {
			nums = append(nums, -1)
		} else {
			num, _ := strconv.Atoi(char)
			nums = append(nums, num)
		}
	}
	return nums
}
