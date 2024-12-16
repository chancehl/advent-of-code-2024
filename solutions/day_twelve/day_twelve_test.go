package main

import (
	"testing"

	"github.com/chancehl/advent-of-code-2024/utils"
)

type testCase struct {
	input    string
	expected int
}

func TestPartOne(t *testing.T) {
	var tests = []testCase{
		{
			input: utils.Dedent(`
				AAAA
				BBCD
				BBCC
				EEEC
			`),
			expected: 140,
		},
		{
			input: utils.Dedent(`
				RRRRIICCFF
				RRRRIICCCF
				VVRRRCCFFF
				VVRCCCJFFF
				VVVVCJJCFE
				VVIVCCJJEE
				VVIIICJJEE
				MIIIIIJJEE
				MIIISIJEEE
				MMMISSJEEE
			`),
			expected: 1930,
		},
	}

	for _, test := range tests {
		t.Run("part one", func(t *testing.T) {
			actual := PartOne(test.input)
			if actual != test.expected {
				t.Errorf("test failed (expected=%d, actual=%d)\n", test.expected, actual)
			}
		})
	}
}

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
