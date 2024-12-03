package main

import (
	"slices"
	"testing"
)

type testCase struct {
	input    string
	expected int
}

func TestPartOne(t *testing.T) {
	var tests = []testCase{
		{input: "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", expected: 161},
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
		{input: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))", expected: 48},
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

func TestFindOperations(t *testing.T) {
	var tests = []struct {
		input         string
		expected      []Operation
		ignoreSignals bool
	}{
		{
			input:         "xmul(2,4)&mul[3,7]!^don't()",
			expected:      []Operation{{instruction: "mul(2,4)", start: 1, end: 9, signal: Do}},
			ignoreSignals: true,
		},
		{
			input: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)",
			expected: []Operation{
				{instruction: "mul(2,4)", start: 1, end: 9, signal: Do},
				{instruction: "mul(5,5)", start: 28, end: 36, signal: Do},
				{instruction: "mul(11,8)", start: 48, end: 57, signal: Do},
			},
			ignoreSignals: true,
		},
		{
			input: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)",
			expected: []Operation{
				{instruction: "mul(2,4)", start: 1, end: 9, signal: Do},
				{instruction: "mul(5,5)", start: 28, end: 36, signal: Dont},
				{instruction: "mul(11,8)", start: 48, end: 57, signal: Dont},
			},
			ignoreSignals: false,
		},
	}

	for _, test := range tests {
		t.Run("FindOperations", func(t *testing.T) {
			actual := findOperations(test.input, test.ignoreSignals)
			for _, expectedOperation := range test.expected {
				if !slices.Contains(actual, expectedOperation) {
					t.Errorf("test failed (expected=%v, actual=%v)\n", test.expected, actual)
				}
			}
		})
	}
}
