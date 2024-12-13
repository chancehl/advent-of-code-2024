package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/chancehl/advent-of-code-2024/ds"
	"github.com/chancehl/advent-of-code-2024/utils/input"
	"github.com/chancehl/advent-of-code-2024/utils/timer"
)

type GuardOrientation int

type PatrolMap [][]string

type PatrolPath ds.Set[string]

type PathNode struct {
	row int
	col int
	dir int
}

func main() {
	path, err := filepath.Abs("solutions/day_six/input.txt")
	if err != nil {
		log.Fatalf("failed to construct path to input: %v", err)
	}

	input, err := input.Read(path)
	if err != nil {
		log.Fatalf("failed to read %s: %v", path, err)
	}

	daySixSolution(input)
}

func daySixSolution(input string) (int, int) {
	partOneResult, partOneRuntime := timer.ExecuteTimedFunc(PartOne, input)
	partTwoResult, partTwoRuntime := timer.ExecuteTimedFunc(PartTwo, input)

	fmt.Printf("[day six / part one] result=%d, time=%.2fms\n", partOneResult, partOneRuntime)
	fmt.Printf("[day six / part two] result=%d, time=%.2fms\n", partTwoResult, partTwoRuntime)

	return partOneResult, partTwoResult
}

func PartOne(input string) int {
	visited := ds.NewSet[string]()
	patrolMap := Create2DMatrixFromInput(input)

	for patrolMap.IsGuardPresent() {
		// 0. get current guard posn
		guardPosn := patrolMap.GetGuardPosn()

		// 1. add curr posn to visited
		visited.Add(guardPosn.Hash(false))

		// 2. determine next guard posn
		nextGuardPosn := patrolMap.GetNextGuardPosn(*guardPosn)

		// 3. update patrol map
		if nextGuardPosn != nil {
			token := GetTokenForDirection(nextGuardPosn.dir)
			patrolMap[nextGuardPosn.row][nextGuardPosn.col] = token
		}
		patrolMap[guardPosn.row][guardPosn.col] = "."
	}

	return visited.Size()
}

func PartTwo(input string) int {
	return -1
}

func Create2DMatrixFromInput(input string) PatrolMap {
	patrolMap := PatrolMap{}
	for _, line := range strings.Split(input, "\n") {
		patrolMap = append(patrolMap, strings.Split(line, ""))
	}
	return patrolMap
}

func CreateGraphFromInput(input string) ds.DirectedGraph[string] {
	graph := ds.NewDirectedGraph[string](ds.DefaultComparator)

	patrolMap := Create2DMatrixFromInput(input)

	for row := range patrolMap {
		for col := range patrolMap[row] {
			key := fmt.Sprintf("%d,%d", row, col)

			// up neighbor
			if row >= 1 && patrolMap[row-1][col] != "#" {
				value := fmt.Sprintf("%d,%d", row-1, col)
				graph.Insert(key, value)
			}

			// down neighbor
			if row < len(patrolMap)-1 && patrolMap[row+1][col] != "#" {
				value := fmt.Sprintf("%d,%d", row+1, col)
				graph.Insert(key, value)
			}

			// left neighbor
			if col >= 1 && patrolMap[row][col-1] != "#" {
				value := fmt.Sprintf("%d,%d", row, col-1)
				graph.Insert(key, value)
			}

			// right neighbor
			if col < len(patrolMap[0])-1 && patrolMap[row][col+1] != "#" {
				value := fmt.Sprintf("%d,%d", row, col+1)
				graph.Insert(key, value)
			}
		}
	}

	return graph
}

func (m *PatrolMap) GetGuardPosn() *PathNode {
	for row := range *m {
		for col := range (*m)[row] {
			isGuard, dir := IsGuard((*m)[row][col])
			if isGuard {
				return &PathNode{row, col, dir}
			}
		}
	}
	return nil
}

func (m *PatrolMap) IsGuardPresent() bool {
	curr := m.GetGuardPosn()
	return curr != nil
}

func (m *PatrolMap) IsInBounds(node PathNode) bool {
	rows := len(*m) - 1
	cols := len((*m)[0]) - 1
	return node.row >= 0 && node.row <= rows && node.col >= 0 && node.col <= cols
}

func (m *PatrolMap) GetNextGuardPosn(curr PathNode) *PathNode {
	next := m.GetForwardPosn(curr)
	if m.IsInBounds(next) {
		token := (*m)[next.row][next.col]
		if token == "#" {
			var rotated PathNode
			switch curr.dir {
			case 0:
				rotated = PathNode{row: curr.row, col: curr.col + 1, dir: 3}
			case 1:
				rotated = PathNode{row: curr.row, col: curr.col - 1, dir: 2}
			case 2:
				rotated = PathNode{row: curr.row - 1, col: curr.col, dir: 0}
			case 3:
				rotated = PathNode{row: curr.row + 1, col: curr.col, dir: 1}
			}
			return &rotated
		}
		return &next
	}
	return nil
}

func (m *PatrolMap) GetForwardPosn(curr PathNode) PathNode {
	var next PathNode
	switch curr.dir {
	case 0:
		next = PathNode{row: curr.row - 1, col: curr.col, dir: curr.dir}
	case 1:
		next = PathNode{row: curr.row + 1, col: curr.col, dir: curr.dir}
	case 2:
		next = PathNode{row: curr.row, col: curr.col - 1, dir: curr.dir}
	case 3:
		next = PathNode{row: curr.row, col: curr.col + 1, dir: curr.dir}
	}
	return next
}

func (n *PathNode) Hash(includeDir bool) string {
	if !includeDir {
		return fmt.Sprintf("%d,%d", n.row, n.col)
	}
	return fmt.Sprintf("%d,%d,%d", n.row, n.col, n.dir)
}

func IsGuard(c string) (bool, int) {
	if c == "^" {
		return true, 0
	}
	if c == "v" {
		return true, 1
	}
	if c == "<" {
		return true, 2
	}
	if c == ">" {
		return true, 3
	}
	return false, -1
}

func GetTokenForDirection(dir int) string {
	switch dir {
	case 0:
		return "^"
	case 1:
		return "v"
	case 2:
		return "<"
	case 3:
		return ">"
	default:
		return "^"
	}
}
