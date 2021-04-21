package algorithms64

func min(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func minPathSum(grid [][]int) int {
	var row int = len(grid)
	var col int = len(grid[0])

	for i := 1; i < row; i++ {
		grid[i][0] = grid[i][0] + grid[i-1][0]
	}

	for j := 1; j < col; j++ {
		grid[0][j] = grid[0][j] + grid[0][j-1]
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
		}
	}

	return grid[row-1][col-1]
}
