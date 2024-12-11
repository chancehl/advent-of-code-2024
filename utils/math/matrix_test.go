package math

import "testing"

func TestComputeDistance(t *testing.T) {
	var tests = []struct {
		x1       int
		x2       int
		y1       int
		y2       int
		expected float64
	}{
		{
			x1:       2,
			y1:       3,
			x2:       5,
			y2:       7,
			expected: 7,
		},
	}

	for _, test := range tests {
		t.Run("ComputeDistance", func(t *testing.T) {
			actual := ComputeDistance(test.x1, test.y1, test.x2, test.y2)
			if actual != test.expected {
				t.Errorf("test failed (expected=%f, actual=%f)\n", test.expected, actual)
			}
		})
	}
}
