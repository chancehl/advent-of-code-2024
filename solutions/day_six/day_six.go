package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/chancehl/advent-of-code-2024/utils/input"
	"github.com/chancehl/advent-of-code-2024/utils/timer"
)

type ElfGuardOrientation int

const (
	Up ElfGuardOrientation = iota
	Down
	Left
	Right
)

type ElfGuardToken = string

const (
	GuardUp    ElfGuardToken = "^"
	GuardDown  ElfGuardToken = "v"
	GuardLeft  ElfGuardToken = "<"
	GuardRight ElfGuardToken = ">"
)

type ElfGuardPatrolPath [][]int

type ElfGuardPatrolMap [][]string

type ElfGuardPatrolState struct {
	guardMap         ElfGuardPatrolMap
	guardRow         int
	guardCol         int
	guardIsPresent   bool
	guardPath        ElfGuardPatrolPath
	guardOrientation ElfGuardOrientation
	guardToken       ElfGuardToken
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
	s := NewElfGuardPatrolState(input)
	for s.guardIsPresent {
		s.MoveGuard()
	}
	return s.guardPath.CountDistinctPositions()
}

func PartTwo(input string) int {
	return -1
}

func NewElfGuardPatrolState(input string) ElfGuardPatrolState {
	guardMap := [][]string{}
	guardRow := -1
	guardCol := -1

	for row, line := range strings.Split(input, "\n") {
		mapRow := []string{}
		for col, char := range strings.Split(line, "") {
			if char == GuardUp { // guard always starts in the up posn
				guardRow = row
				guardCol = col
			}
			mapRow = append(mapRow, char)
		}
		guardMap = append(guardMap, mapRow)
	}

	return ElfGuardPatrolState{
		guardMap:         guardMap,
		guardRow:         guardRow,
		guardCol:         guardCol,
		guardOrientation: Up,
		guardToken:       GuardUp,
		guardIsPresent:   true,
		guardPath:        [][]int{{guardRow, guardCol}}, // always append the starting posn to the path
	}
}

func (s *ElfGuardPatrolState) MoveGuard() {
	nextToken := s.GetTokenInFrontOfGuard()
	if nextToken == nil {
		oldRow, oldCol := s.GetCurrentGuardPosn()
		s.guardMap[oldCol][oldRow] = "."
		s.guardIsPresent = false
	} else if *nextToken == "#" {
		oldRow, oldCol := s.GetCurrentGuardPosn()
		newRow, newCol := s.GetNextGuardTurnedPosn()
		newGuardToken := s.GetRotatedGuardToken()
		newGuardOrientation := s.GetRotatedGuardOrientation()

		s.guardOrientation = newGuardOrientation
		s.guardToken = newGuardToken
		s.guardMap[newRow][newCol] = newGuardToken
		s.guardMap[oldRow][oldCol] = "."
		s.guardRow = newRow
		s.guardCol = newCol
		s.guardPath = append(s.guardPath, []int{newRow, newCol})
	} else {
		oldRow, oldCol := s.GetCurrentGuardPosn()
		newRow, newCol := s.GetNextGuardForwardPosn()

		s.guardMap[newRow][newCol] = s.guardToken
		s.guardMap[oldRow][oldCol] = "."
		s.guardRow = newRow
		s.guardCol = newCol
		s.guardPath = append(s.guardPath, []int{newRow, newCol})
	}
}

func (s *ElfGuardPatrolState) GetCurrentGuardPosn() (int, int) {
	return s.guardRow, s.guardCol
}

func (s *ElfGuardPatrolState) GetNextGuardTurnedPosn() (int, int) {
	orientation := s.guardOrientation

	row := s.guardRow
	col := s.guardCol

	switch orientation {
	case Up:
		return row, col + 1
	case Down:
		return row, col - 1
	case Left:
		return row - 1, col
	case Right:
		return row + 1, col
	default:
		return row, col
	}
}

func (s *ElfGuardPatrolState) GetRotatedGuardToken() ElfGuardToken {
	orientation := s.guardOrientation

	switch orientation {
	case Up:
		return GuardRight
	case Down:
		return GuardLeft
	case Left:
		return GuardUp
	case Right:
		return GuardDown
	default:
		return s.guardToken
	}
}

func (s *ElfGuardPatrolState) GetRotatedGuardOrientation() ElfGuardOrientation {
	orientation := s.guardOrientation

	switch orientation {
	case Up:
		return Right
	case Down:
		return Left
	case Left:
		return Up
	case Right:
		return Down
	default:
		return orientation
	}
}

func (s *ElfGuardPatrolState) GetNextGuardForwardPosn() (int, int) {
	orientation := s.guardOrientation

	row := s.guardRow
	col := s.guardCol

	switch orientation {
	case Up:
		return row - 1, col
	case Down:
		return row + 1, col
	case Left:
		return row, col - 1
	case Right:
		return row, col + 1
	default:
		return row, col
	}
}

func (s *ElfGuardPatrolState) GetTokenInFrontOfGuard() *string {
	newRow, newCol := s.GetNextGuardForwardPosn()
	return s.guardMap.GetTokenAt(newRow, newCol)
}

func (p *ElfGuardPatrolPath) CountDistinctPositions() int {
	visited := make(map[string]bool)
	distinct := 0
	for _, posn := range *p {
		key := fmt.Sprintf("%d,%d", posn[0], posn[1])
		if !visited[key] {
			distinct += 1
		}
		visited[key] = true

	}
	return distinct
}

func (m *ElfGuardPatrolMap) GetTokenAt(row, col int) *string {
	if m.IsPositionInBounds(row, col) {
		return &(*m)[row][col]
	}
	return nil
}

func (m *ElfGuardPatrolMap) IsPositionInBounds(row, col int) bool {
	return row >= 0 && row <= len(*m)-1 && col >= 0 && col <= len((*m)[0])-1
}

func (m *ElfGuardPatrolMap) PrintMap() {
	lines := []string{}
	for _, row := range *m {
		lines = append(lines, strings.Join(row, ""))
	}
	fmt.Println(strings.Join(lines, "\n"))
}

func IsGuardToken(s string) bool {
	return s == GuardUp || s == GuardDown || s == GuardLeft || s == GuardRight
}
