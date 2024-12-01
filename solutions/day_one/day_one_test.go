package main

import "testing"

type testCase struct {
	name     string
	input    string
	expected int
}

func TestPartOne(t *testing.T) {
	var tests = []testCase{
		{input: `3   4
4   3
2   5
1   3
3   9
3   3`, expected: 11},
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
		{input: `3   4
4   3
2   5
1   3
3   9
3   3`, expected: 31},
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
