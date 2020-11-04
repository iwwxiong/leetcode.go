package algorithms29

import (
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend == math.MinInt32 && divisor == 1 {
		return math.MinInt32
	}
	var sign bool
	if (dividend >= 0 && divisor > 0) || (dividend < 0 && divisor < 0) {
		sign = true
	} else {
		sign = false
	}
	if dividend > 0 {
		dividend = opposite(dividend)
	}
	if divisor > 0 {
		divisor = opposite(divisor)
	}
	var result int = negativeDivide(dividend, divisor)
	if !sign {
		result = opposite(result)
	}
	return result
}

func negativeDivide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	var result int = 0
	if dividend > divisor {
		return result
	}
	var oldDivisor int = divisor
	var step int = 1
	result = result + step
	dividend = dividend - divisor

	for dividend < divisor {
		step = step + step
		divisor = divisor + divisor
		if dividend <= divisor {
			dividend = dividend - divisor
			result = result + step
		}
	}
	return result + negativeDivide(dividend, oldDivisor)

}

func opposite(x int) int {
	return ^x + 1
}
