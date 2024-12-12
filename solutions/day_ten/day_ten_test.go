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
				0123
				1234
				8765
				9876
			`),
			expected: 1,
		},
		{
			input: utils.Dedent(`
				...0...
				...1...
				...2...
				6543456
				7.....7
				8.....8
				9.....9
			`),
			expected: 2,
		},
		{
			input: utils.Dedent(`
				..90..9
				...1.98
				...2..7
				6543456
				765.987
				876....
				987....
			`),
			expected: 4,
		},
		{
			input: utils.Dedent(`
				10..9..
				2...8..
				3...7..
				4567654
				...8..3
				...9..2
				.....01
			`),
			expected: 3,
		},
		{
			input: utils.Dedent(`
				89010123
				78121874
				87430965
				96549874
				45678903
				32019012
				01329801
				10456732
			`),
			expected: 36,
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
