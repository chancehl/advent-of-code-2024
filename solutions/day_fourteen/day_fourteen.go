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
)

type RobotVelocity struct {
	x int
	y int
}

type RobotPosition struct {
	x int
	y int
}

type Robot struct {
	id       int
	position RobotPosition
	velocity RobotVelocity
}

func main() {
	path, err := filepath.Abs("solutions/day_fourteen/input.txt")
	if err != nil {
		log.Fatalf("failed to construct path to input: %v", err)
	}

	input, err := input.Read(path)
	if err != nil {
		log.Fatalf("failed to read %s: %v", path, err)
	}

	dayFourteenSolution(input)
}

func dayFourteenSolution(input string) (int, int) {
	partOneResult, partOneRuntime := timer.ExecuteTimedFunc(PartOne, input)
	partTwoResult, partTwoRuntime := timer.ExecuteTimedFunc(PartTwo, input)

	fmt.Printf("[day fourteen / part one] result=%d, time=%.2fms\n", partOneResult, partOneRuntime)
	fmt.Printf("[day fourteen / part two] result=%d, time=%.2fms\n", partTwoResult, partTwoRuntime)

	return partOneResult, partTwoResult
}

func PartOne(input string) int {
	return -1
}

func PartTwo(input string) int {
	return -1
}

func ParseRobotInfo(input string) []*Robot {
	trajectories := make([]*Robot, 0)
	for idx, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		posnPart := parts[0]
		velocityPart := parts[1]

		numberRegex, _ := regexp.Compile(`-?\d+`)

		posns := numberRegex.FindAllString(posnPart, -1)
		xPosn, _ := strconv.Atoi(posns[0])
		yPosn, _ := strconv.Atoi(posns[1])
		posn := RobotPosition{x: xPosn, y: yPosn}

		velocities := numberRegex.FindAllString(velocityPart, -1)
		xVelocity, _ := strconv.Atoi(velocities[0])
		yVelocity, _ := strconv.Atoi(velocities[1])
		velocity := RobotVelocity{xVelocity, yVelocity}

		trajectories = append(trajectories, &Robot{id: idx, position: posn, velocity: velocity})
	}
	return trajectories
}
