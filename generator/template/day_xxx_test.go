package main

import "testing"

func TestPartOne(t *testing.T) {
	var tests = []struct {
		name     string
		input    []string
		expected int
	}{
		{input: []string{"a", "b", "c", "d"}, expected: -1},
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
	var tests = []struct {
		name     string
		input    []string
		expected int
	}{
		{input: []string{"a", "b", "c", "d"}, expected: -1},
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
