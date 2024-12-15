package timer

import "time"

type ExecutableFunc[T int | int64] func(input string) T

func ExecuteTimedFunc[T int | int64](f ExecutableFunc[T], input string) (T, float64) {
	start := time.Now()
	result := f(input)
	end := time.Since(start).Seconds() * 1000

	return result, end
}
