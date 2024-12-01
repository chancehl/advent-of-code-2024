package timer

import "time"

type ExecutableFunc func(input string) int

func ExecuteTimedFunc(f ExecutableFunc, input string) (int, int64) {
	start := time.Now()
	result := f(input)
	end := time.Since(start).Milliseconds()

	return result, end
}
