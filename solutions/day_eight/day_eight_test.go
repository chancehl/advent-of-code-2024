package main

import (
	"slices"
	"testing"

	"github.com/chancehl/advent-of-code-2024/utils"
)

type testCase struct {
	input    string
	expected int
}

// func TestPartOne(t *testing.T) {
// 	var tests = []testCase{
// 		{
// 			input: utils.Dedent(`
// 				............
// 				........0...
// 				.....0......
// 				.......0....
// 				....0.......
// 				......A.....
// 				............
// 				............
// 				........A...
// 				.........A..
// 				............
// 				............
// 			`),
// 			expected: 14,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run("part one", func(t *testing.T) {
// 			actual := PartOne(test.input)
// 			if actual != test.expected {
// 				t.Errorf("test failed (expected=%d, actual=%d)\n", test.expected, actual)
// 			}
// 		})
// 	}
// }

func TestPartTwo(t *testing.T) {
	var tests = []testCase{
		{input: "abcd", expected: -1},
	}

	for _, test := range tests {
		t.Run("part two", func(t *testing.T) {
			actual := PartTwo(test.input)
			if actual != test.expected {
				t.Errorf("test failed (expected=%d, actual=%d)\n", test.expected, actual)
			}
		})
	}
}

func TestComputeAntinodes(t *testing.T) {
	var tests = []struct {
		input    AntennaMap
		expected map[string][]Coordinates
	}{
		{
			input: CreateAntennaMapFromInput(utils.Dedent(`
				..........
				..........
				..........
				....a.....
				..........
				.....a....
				..........
				..........
				..........
				..........
			`)),
			expected: map[string][]Coordinates{
				"a": {
					{row: 1, col: 3},
					{row: 7, col: 6},
				},
			},
		},
		// {
		// 	input: CreateAntennaMapFromInput(utils.Dedent(`
		// 		............
		// 		........0...
		// 		.....0......
		// 		.......0....
		// 		....0.......
		// 		......A.....
		// 		............
		// 		............
		// 		........A...
		// 		.........A..
		// 		............
		// 		............
		// 	`)),
		// 	expected: []Coordinates{},
		// },
	}

	for _, test := range tests {
		t.Run("ComputeAntinodes", func(t *testing.T) {
			actual := test.input.ComputeAntinodes()
			for key := range test.expected {
				if !slices.Equal(actual[key], test.expected[key]) {
					t.Errorf("test failed (expected=%v, actual=%v)\n", test.expected, actual)
				}
			}
		})
	}
}

func TestFindAntennae(t *testing.T) {
	input := utils.Dedent(`
		............
		........0...
		.....0......
		.......0....
		....0.......
		......A.....
		............
		............
		........A...
		.........A..
		............
		............
	`)

	antennaMap := CreateAntennaMapFromInput(input)

	var tests = []struct {
		input    AntennaMap
		expected map[string][]Coordinates
	}{
		{
			input: antennaMap,
			expected: map[string][]Coordinates{
				"A": {
					{row: 5, col: 6},
					{row: 8, col: 8},
					{row: 9, col: 9},
				},
				"0": {
					{row: 1, col: 8},
					{row: 2, col: 5},
					{row: 3, col: 7},
					{row: 4, col: 4},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run("FindAntennae", func(t *testing.T) {
			actual := test.input.FindAntennae()
			for key := range test.expected {
				if !slices.Equal(actual[key], test.expected[key]) {
					t.Errorf("test failed (expected=%v, actual=%v)\n", test.expected, actual)
				}
			}
		})
	}
}
