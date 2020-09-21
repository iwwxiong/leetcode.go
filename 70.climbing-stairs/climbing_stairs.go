package algorithms70

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	var temp int
	var n1, n2 int = 1, 2
	for i := 3; i <= n; i++ {
		temp = n2
		n2 = n1 + n2
		n1 = temp
	}
	return n2
}
