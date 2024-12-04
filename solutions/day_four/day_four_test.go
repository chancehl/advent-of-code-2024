package main

import (
	"testing"
)

type testCase struct {
	input    string
	expected int
}

type searchDirectionTestCase struct {
	name     string
	row      int
	col      int
	expected bool
}

const puzzleInput = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestPartOne(t *testing.T) {
	var tests = []testCase{
		{input: puzzleInput, expected: 18},
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

func TestSearchUp(t *testing.T) {
	matrix := createMatrixFromInput(puzzleInput)

	var tests = []searchDirectionTestCase{{
		name:     "out of bounds",
		row:      0,
		col:      0,
		expected: false,
	}, {
		name:     "in bounds and not found",
		row:      3,
		col:      1,
		expected: false,
	}, {
		name:     "in bounds and found",
		row:      9,
		col:      9,
		expected: true,
	}}

	for _, test := range tests {
		t.Run("search up", func(t *testing.T) {
			actual := searchUp(matrix, test.row, test.col)
			if actual != test.expected {
				t.Errorf("test \"%s\" failed (expected=%v, actual=%v)\n", test.name, test.expected, actual)
			}
		})
	}
}

func TestSearchDown(t *testing.T) {
	matrix := createMatrixFromInput(puzzleInput)

	var tests = []searchDirectionTestCase{{
		name:     "out of bounds",
		row:      9,
		col:      0,
		expected: false,
	}, {
		name:     "in bounds and not found",
		row:      3,
		col:      1,
		expected: false,
	}, {
		name:     "in bounds and found",
		row:      3,
		col:      9,
		expected: true,
	}}

	for _, test := range tests {
		t.Run("search down", func(t *testing.T) {
			actual := searchDown(matrix, test.row, test.col)
			if actual != test.expected {
				t.Errorf("test \"%s\" failed (expected=%v, actual=%v)\n", test.name, test.expected, actual)
			}
		})
	}
}

func TestSearchRight(t *testing.T) {
	matrix := createMatrixFromInput(puzzleInput)

	var tests = []searchDirectionTestCase{{
		name:     "out of bounds",
		row:      0,
		col:      9,
		expected: false,
	}, {
		name:     "in bounds and not found",
		row:      0,
		col:      0,
		expected: false,
	}, {
		name:     "in bounds and found",
		row:      0,
		col:      5,
		expected: true,
	}}

	for _, test := range tests {
		t.Run("search right", func(t *testing.T) {
			actual := searchRight(matrix, test.row, test.col)
			if actual != test.expected {
				t.Errorf("test \"%s\" failed (expected=%v, actual=%v)\n", test.name, test.expected, actual)
			}
		})
	}
}

func TestSearchLeft(t *testing.T) {
	matrix := createMatrixFromInput(puzzleInput)

	var tests = []searchDirectionTestCase{{
		name:     "out of bounds",
		row:      0,
		col:      0,
		expected: false,
	}, {
		name:     "in bounds and not found",
		row:      0,
		col:      4,
		expected: false,
	}, {
		name:     "in bounds and found",
		row:      1,
		col:      4,
		expected: true,
	}}

	for _, test := range tests {
		t.Run("search left", func(t *testing.T) {
			actual := searchLeft(matrix, test.row, test.col)
			if actual != test.expected {
				t.Errorf("test \"%s\" failed (expected=%v, actual=%v)\n", test.name, test.expected, actual)
			}
		})
	}
}

func TestSearchDiagonalUpLeft(t *testing.T) {
	matrix := createMatrixFromInput(puzzleInput)

	var tests = []searchDirectionTestCase{{
		name:     "out of bounds",
		row:      0,
		col:      0,
		expected: false,
	}, {
		name:     "in bounds and not found",
		row:      3,
		col:      3,
		expected: false,
	}, {
		name:     "in bounds and found",
		row:      9,
		col:      9,
		expected: true,
	}}

	for _, test := range tests {
		t.Run("search up left", func(t *testing.T) {
			actual := searchUpLeft(matrix, test.row, test.col)
			if actual != test.expected {
				t.Errorf("test \"%s\" failed (expected=%v, actual=%v)\n", test.name, test.expected, actual)
			}
		})
	}
}

func TestSearchDiagonalUpRight(t *testing.T) {
	matrix := createMatrixFromInput(puzzleInput)

	var tests = []searchDirectionTestCase{{
		name:     "out of bounds",
		row:      0,
		col:      9,
		expected: false,
	}, {
		name:     "in bounds and not found",
		row:      3,
		col:      3,
		expected: false,
	}, {
		name:     "in bounds and found",
		row:      5,
		col:      0,
		expected: true,
	}}

	for _, test := range tests {
		t.Run("search up right", func(t *testing.T) {
			actual := searchUpRight(matrix, test.row, test.col)
			if actual != test.expected {
				t.Errorf("test \"%s\" failed (expected=%v, actual=%v)\n", test.name, test.expected, actual)
			}
		})
	}
}

func TestSearchDiagonalDownLeft(t *testing.T) {
	matrix := createMatrixFromInput(puzzleInput)

	var tests = []searchDirectionTestCase{{
		name:     "out of bounds",
		row:      9,
		col:      9,
		expected: false,
	}, {
		name:     "in bounds and not found",
		row:      0,
		col:      4,
		expected: false,
	}, {
		name:     "in bounds and found",
		row:      3,
		col:      9,
		expected: true,
	}}

	for _, test := range tests {
		t.Run("search down left", func(t *testing.T) {
			actual := searchDownLeft(matrix, test.row, test.col)
			if actual != test.expected {
				t.Errorf("test \"%s\" failed (expected=%v, actual=%v)\n", test.name, test.expected, actual)
			}
		})
	}
}

func TestSearchDiagonalDownRight(t *testing.T) {
	matrix := createMatrixFromInput(puzzleInput)

	var tests = []searchDirectionTestCase{{
		name:     "out of bounds",
		row:      9,
		col:      9,
		expected: false,
	}, {
		name:     "in bounds and not found",
		row:      0,
		col:      0,
		expected: false,
	}, {
		name:     "in bounds and found",
		row:      0,
		col:      4,
		expected: true,
	}}

	for _, test := range tests {
		t.Run("search down right", func(t *testing.T) {
			actual := searchDownRight(matrix, test.row, test.col)
			if actual != test.expected {
				t.Errorf("test \"%s\" failed (expected=%v, actual=%v)\n", test.name, test.expected, actual)
			}
		})
	}
}
