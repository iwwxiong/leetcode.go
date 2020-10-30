package algorithms29

import (
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	var result int = 0
	if dividend < divisor {
		return result
	}
	var oldDivisor int = divisor
	var step int = 1
	result = result + step
	dividend = dividend - divisor

	for dividend >= divisor {
		step = step + step
		divisor = divisor + divisor
		dividend = dividend - divisor
		result = result + step
	}
	return result + divide(dividend, oldDivisor)
}
