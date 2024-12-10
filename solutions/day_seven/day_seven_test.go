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
				190: 10 19
				3267: 81 40 27
				83: 17 5
				156: 15 6
				7290: 6 8 6 15
				161011: 16 10 13
				192: 17 8 14
				21037: 9 7 18 13
				292: 11 6 16 20
			`),
			expected: 3749,
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
		{
			input: utils.Dedent(`
				190: 10 19
				3267: 81 40 27
				83: 17 5
				156: 15 6
				7290: 6 8 6 15
				161011: 16 10 13
				192: 17 8 14
				21037: 9 7 18 13
				292: 11 6 16 20
			`),
			expected: 11387,
		},
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

func TestIsValidEquation(t *testing.T) {
	var tests = []struct {
		input      ElfEquation
		expected   bool
		simplified bool
	}{

		{
			input:      ElfEquation{target: 6, operands: []int{1, 2, 3}},
			simplified: true,
			expected:   true,
		},
		{
			input:      ElfEquation{target: 190, operands: []int{19, 10}},
			simplified: true,
			expected:   true,
		},
		{
			input:      ElfEquation{target: 3267, operands: []int{81, 40, 27}},
			simplified: true,
			expected:   true,
		},
		{
			input:      ElfEquation{target: 292, operands: []int{11, 6, 16, 20}},
			simplified: true,
			expected:   true,
		},
		{
			input:      ElfEquation{target: 156, operands: []int{15, 6}},
			simplified: true,
			expected:   false,
		},
		{
			input:      ElfEquation{target: 7290, operands: []int{6, 8, 6, 15}},
			simplified: true,
			expected:   false,
		},
		{
			input:      ElfEquation{target: 156, operands: []int{15, 6}},
			simplified: false,
			expected:   true,
		},
		{
			input:      ElfEquation{target: 7290, operands: []int{6, 8, 6, 15}},
			simplified: false,
			expected:   true,
		},
	}

	for _, test := range tests {
		t.Run("IsValidEquationPart2", func(t *testing.T) {
			actual := IsValidEquation(test.input, test.simplified)
			if actual != test.expected {
				t.Errorf("IsValidEquation(%v) failed (expected=%v, actual=%v)\n", test.input, test.expected, actual)
			}
		})
	}
}
