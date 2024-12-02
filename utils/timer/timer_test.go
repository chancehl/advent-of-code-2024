// ChatGPT wrote these 🤖
package timer

import (
	"testing"
	"time"
)

func TestExecuteTimedFunc(t *testing.T) {
	tests := []struct {
		name            string
		mockFunc        ExecutableFunc
		input           string
		expectedResult  int
		expectedMinTime int64
	}{
		{
			name: "Instantaneous execution",
			mockFunc: func(input string) int {
				return len(input)
			},
			input:           "hello",
			expectedResult:  5,
			expectedMinTime: 0, // No delay, so time should be very small
		},
		{
			name: "Delayed execution",
			mockFunc: func(input string) int {
				time.Sleep(50 * time.Millisecond)
				return len(input)
			},
			input:           "world",
			expectedResult:  5,
			expectedMinTime: 50, // Ensure the delay is reflected in execution time
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, elapsed := ExecuteTimedFunc(tt.mockFunc, tt.input)

			// Check the returned result
			if result != tt.expectedResult {
				t.Errorf("Expected result: %d, got: %d", tt.expectedResult, result)
			}

			// Check that the elapsed time is at least the expected minimum time
			if elapsed < tt.expectedMinTime {
				t.Errorf("Expected elapsed time >= %dms, got: %dms", tt.expectedMinTime, elapsed)
			}
		})
	}
}
