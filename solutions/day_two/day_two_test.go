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
		{input: `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`, expected: 4},
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

func TestIsSafe(t *testing.T) {
	var tests = []struct {
		name     string
		report   ElfReport
		expected bool
	}{
		{
			name:     "All increasing by stable amount",
			report:   []int{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			name:     "All decreasing by stable amount",
			report:   []int{5, 4, 3, 2, 1},
			expected: true,
		},
		{
			name:     "All increasing, but unstable jumps",
			report:   []int{1, 5, 6, 8, 12},
			expected: false,
		},
		{
			name:     "All decreasing, but unstable jumps",
			report:   []int{15, 12, 9, 4, 0},
			expected: false,
		},
		{
			name:     "Mix of increasing and decreasing values",
			report:   []int{1, 5, 2, 7, 3, 0},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.report.IsSafe()
			if actual != test.expected {
				t.Errorf("test %s failed (expected=%v, actual=%v)\n", test.name, test.expected, actual)
			}
		})
	}
}

func TestIsSafeWithDampener(t *testing.T) {
	var tests = []struct {
		name     string
		report   ElfReport
		expected bool
	}{
		{
			name:     "Safe regardless of what level is removed",
			report:   []int{7, 6, 4, 2, 1},
			expected: true,
		},
		{
			name:     "Unsafe regardless of what level is removed",
			report:   []int{1, 2, 7, 8, 9},
			expected: false,
		},
		{
			name:     "Safe by removing a single level",
			report:   []int{1, 3, 2, 4, 5},
			expected: true,
		},
		{
			name:     "Safe by removing any level",
			report:   []int{1, 3, 6, 7, 9},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.report.IsSafeWithDampener()
			if actual != test.expected {
				t.Errorf("test %s failed (expected=%v, actual=%v)\n", test.name, test.expected, actual)
			}
		})
	}
}
