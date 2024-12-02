package main

import "testing"

type testCase struct {
	name     string
	input    string
	expected int
}

func TestPartOne(t *testing.T) {
	var tests = []testCase{
		{input: `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`, expected: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := PartOne(test.input)
			if actual != test.expected {
				t.Errorf("test %s failed (expected=%d, actual=%d)\n", test.name, test.expected, actual)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	var tests = []testCase{
		{input: "abcd", expected: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := PartTwo(test.input)
			if actual != test.expected {
				t.Errorf("test %s failed (expected=%d, actual=%d)\n", test.name, test.expected, actual)
			}
		})
	}
}
