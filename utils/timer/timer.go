package timer

import "time"

type ExecutableFunc func(input string) int

func ExecuteTimedFunc(f ExecutableFunc, input string) (int, float64) {
	start := time.Now()
	result := f(input)
	end := time.Since(start).Seconds() * 1000

	return result, end
}
