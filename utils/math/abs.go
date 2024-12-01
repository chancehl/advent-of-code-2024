package math

func Abs(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
