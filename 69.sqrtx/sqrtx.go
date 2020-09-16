package algorithms69

import (
	"math"
)

func mySqrt(x int) int {
	var g float64 = float64(x)
	for g*g-float64(x) < -0.1 || g*g-float64(x) >= 0.1 {
		g = (g + float64(x)/g) / 2.0
	}
	var ans int = int(math.Round(g))
	if ans*ans > x {
		ans--
	}
	return ans
}
