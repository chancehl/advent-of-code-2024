package main

import (
	"slices"
	"testing"

	"github.com/chancehl/advent-of-code-2024/ds"
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
		expected map[string][]ds.Coordinates
	}{
		{
			input: antennaMap,
			expected: map[string][]ds.Coordinates{
				"A": {
					{Row: 5, Col: 6},
					{Row: 8, Col: 8},
					{Row: 9, Col: 9},
				},
				"0": {
					{Row: 1, Col: 8},
					{Row: 2, Col: 5},
					{Row: 3, Col: 7},
					{Row: 4, Col: 4},
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
