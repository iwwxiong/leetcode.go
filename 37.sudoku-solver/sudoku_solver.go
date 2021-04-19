package algorithms37

func isValid(board [][]byte, i, j int, data byte) bool {
	for m := 0; m < 9; m++ {
		if board[m][j] == data {
			return false
		}
	}
	for n := 0; n < 9; n++ {
		if board[i][n] == data {
			return false
		}
	}

	i1 := i / 3 * 3
	j1 := j / 3 * 3
	for m := 0; m < 3; m++ {
		for n := 0; n < 3; n++ {
			if board[i1+m][j1+n] == data {
				return false
			}
		}
	}
	return true
}

func solveSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				var data byte = '1'
				for data <= '9' {
					if isValid(board, i, j, data) {
						board[i][j] = data
						if solveSudoku(board) {
							return true
						} else {
							board[i][j] = '.'
						}
					}
					data++
				}
				return false
			}
		}
	}
	return true
}
