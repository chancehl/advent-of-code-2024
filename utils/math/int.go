package math

import "math"

func CountDigits(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func SplitNumber(n int) (int, int) {
	digits := int(math.Log10(float64(n)) + 1)

	divisor := int(math.Pow(10, float64(digits/2)))

	firstHalf := n / divisor
	secondHalf := n % divisor

	return firstHalf, secondHalf
}
