package algorithms63

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	var m int = len(obstacleGrid)
	var n int = len(obstacleGrid[0])

	var dp = make([]int, m)
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			dp[i] = 0
			break
		} else {
			dp[i] = 1
		}
	}
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[j][i] == 1 {
				dp[j] = 0
			} else {
				if j > 0 {
					dp[j] = dp[j] + dp[j-1]
				}
			}
		}
	}
	return dp[m-1]
}
