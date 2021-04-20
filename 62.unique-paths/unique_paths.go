package algorithms62

func canMoveLeft(n int, j int) bool {
	return j < n-1
}

func canMoveDown(m int, i int) bool {
	return i < m-1
}

func moveStep(paths [][]int, m, n int) [][]int {
	var p [][]int
	for _, path := range paths {
		if canMoveDown(m, path[0]) {
			p = append(p, []int{path[0] + 1, path[1]})
		}
		if canMoveLeft(n, path[1]) {
			p = append(p, []int{path[0], path[1] + 1})
		}
	}
	return p
}

func uniquePaths(m int, n int) int {
	var paths [][]int = [][]int{{0, 0}}
	for step := 0; step < m+n-2; step++ {
		paths = moveStep(paths, m, n)
	}
	var num int
	for _, path := range paths {
		if path[0] == m-1 && path[1] == n-1 {
			num++
		}
	}
	return num
}

// 动态规划
func uniquePathsV1(m int, n int) int {
	var dp = make([]int, m)
	for i := 0; i < m; i++ {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}
	return dp[m-1]
}
