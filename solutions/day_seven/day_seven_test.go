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
			expected: -1,
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

func TestIsValidEquation(t *testing.T) {
	var tests = []struct {
		input    ElfEquation
		expected bool
	}{
		{
			input:    ElfEquation{target: 190, operands: []int{19, 10}},
			expected: true,
		},
		// {
		// 	input:    ElfEquation{target: 3267, operands: []int{81, 40, 27}},
		// 	expected: true,
		// },
		// {
		// 	input:    ElfEquation{target: 292, operands: []int{11, 6, 16, 20}},
		// 	expected: true,
		// },
	}

	for _, test := range tests {
		t.Run("IsValidEquation", func(t *testing.T) {
			actual := IsValidEquation(test.input.target, test.input.operands)
			if actual != test.expected {
				t.Errorf("IsValidEquation(%+v) failed (expected=%v, actual=%v)\n", test.input, test.expected, actual)
			}
		})
	}
}

func TestGenerateOperationPermutations(t *testing.T) {
	var tests = []struct {
		n        int
		expected []string
	}{
		{
			n:        1,
			expected: []string{"*", "+"},
		},
		{
			n:        2,
			expected: []string{"**", "++", "+*", "*+"},
		},
	}

	for _, test := range tests {
		t.Run("GenerateOperationPermutations", func(t *testing.T) {
			actual := GenerateOperationPermutations(test.n)
			for _, expected := range test.expected {
				if !slices.Contains(actual, expected) {
					t.Errorf("GenerateOperationPermutations(%d) failed (expected=%v, actual=%v)\n", test.n, expected, actual)
				}
			}
		})
	}
}
