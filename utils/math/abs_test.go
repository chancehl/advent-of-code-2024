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
			name:     "Positive difference",
			a:        10,
			b:        5,
			expected: 5,
		},
		{
			name:     "Negative difference",
			a:        5,
			b:        10,
			expected: 5,
		},
		{
			name:     "Zero difference",
			a:        7,
			b:        7,
			expected: 0,
		},
		{
			name:     "Large positive difference",
			a:        1000,
			b:        100,
			expected: 900,
		},
		{
			name:     "Large negative difference",
			a:        100,
			b:        1000,
			expected: 900,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Abs(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Abs(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
