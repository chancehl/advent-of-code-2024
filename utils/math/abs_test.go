package math

import "testing"

func TestAbs(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "positive difference",
			a:        10,
			b:        5,
			expected: 5,
		},
		{
			name:     "negative difference",
			a:        5,
			b:        10,
			expected: 5,
		},
		{
			name:     "zero difference",
			a:        7,
			b:        7,
			expected: 0,
		},
		{
			name:     "large positive difference",
			a:        1000,
			b:        100,
			expected: 900,
		},
		{
			name:     "large negative difference",
			a:        100,
			b:        1000,
			expected: 900,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Abs(test.a, test.b)
			if result != test.expected {
				t.Errorf("Abs(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
			}
		})
	}
}
