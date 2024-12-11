package main

import (
	"testing"
)

type testCase struct {
	input    string
	expected int
}

func TestPartOne(t *testing.T) {
	var tests = []testCase{
		{input: "2333133121414131402", expected: 1928},
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

func TestCreateExpandedDiskMap(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{input: "12345", expected: "0..111....22222"},
		{input: "2333133121414131402", expected: "00...111...2...333.44.5555.6666.777.888899"},
	}

	for _, test := range tests {
		t.Run("CreateExpandedDiskMap", func(t *testing.T) {
			actual := CreateExpandedDiskMap(test.input)
			if actual != test.expected {
				t.Errorf("test failed (expected=%s, actual=%s)\n", test.expected, actual)
			}
		})
	}
}

func TestMoveFiles(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{input: "0..111....22222", expected: "022111222......"},
		{input: "00...111...2...333.44.5555.6666.777.888899", expected: "0099811188827773336446555566.............."},
	}

	for _, test := range tests {
		t.Run("MoveFiles", func(t *testing.T) {
			actual := MoveFiles(test.input)
			if actual != test.expected {
				t.Errorf("test failed (expected=%s, actual=%s)\n", test.expected, actual)
			}
		})
	}
}

func TestCalculateChecksum(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{input: "0099811188827773336446555566..............", expected: 1928},
	}

	for _, test := range tests {
		t.Run("CalculateChecksum", func(t *testing.T) {
			actual := CalculateChecksum(test.input)
			if actual != test.expected {
				t.Errorf("test failed (expected=%d, actual=%d)\n", test.expected, actual)
			}
		})
	}
}
