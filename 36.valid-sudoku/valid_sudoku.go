package algorithms36

func isValidSudoku(board [][]byte) bool {
	var rows, columns, boxes []map[byte]bool
	for i := 0; i < 9; i++ {
		rows = append(rows, make(map[byte]bool))
		columns = append(columns, make(map[byte]bool))
		boxes = append(boxes, make(map[byte]bool))
	}

	for row := range board {
		for index, num := range board[row] {
			if num == '.' {
				continue
			}
			if rows[row][num] || columns[index][num] || boxes[(row/3)*3+index/3][num] {
				return false
			}
			rows[row][num] = true
			columns[index][num] = true
			boxes[(row/3)*3+index/3][num] = true
		}
	}
	return true
}
