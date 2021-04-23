package algorithms72

func min3int(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
}

func minDistance(word1 string, word2 string) int {
	var row int = len(word1)
	var col int = len(word2)
	if row == 0 && col == 0 {
		return 0
	}

	if row == 0 {
		return col
	}

	if col == 0 {
		return row
	}
	// 二维数组初始化
	dp := make([][]int, row+1)
	for index := range dp {
		dp[index] = make([]int, col+1)
	}

	for i := 0; i <= row; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= col; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min3int(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[row][col]
}
