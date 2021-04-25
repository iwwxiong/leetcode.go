package algorithms85

func max2int(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxAreas(dp []int) int {
	var max int = 0
	if len(dp) == 1 {
		return dp[0]
	}
	for i := 0; i < len(dp); i++ {
		if dp[i] > 0 {
			var sum int = dp[i]
			var j int = 1
			for i-j >= 0 {
				if dp[i-j] == 0 || dp[i] > dp[i-j] {
					break
				}
				sum += dp[i]
				j++
			}
			j = 1
			for i+j < len(dp) {
				if dp[i+j] == 0 || dp[i] > dp[i+j] {
					break
				}
				sum += dp[i]
				j++
			}
			max = max2int(max, sum)
		}
	}
	return max
}

func maximalRectangle(matrix [][]byte) int {
	var row int = len(matrix)
	if row == 0 {
		return 0
	}
	var col int = len(matrix[0])
	var dp []int = make([]int, col)
	for j := 0; j < col; j++ {
		if matrix[0][j] == '1' {
			dp[j] = 1
		} else {
			dp[j] = 0
		}
	}
	var maximal int = maxAreas(dp)
	for i := 1; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == '1' {
				dp[j] = dp[j] + 1
			} else {
				dp[j] = 0
			}
		}
		maximal = max2int(maximal, maxAreas(dp))
	}
	return maximal
}
