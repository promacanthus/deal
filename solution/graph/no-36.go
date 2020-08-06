package graph

func isValidSudoku(board [][]byte) bool {
	if board == nil || len(board) == 0 {
		return false
	}

	rowMap := make(map[byte]bool, 9)
	colMap := make(map[byte]bool, 9)
	subMap := make(map[int]bool, 9)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if rowMap[board[i][j]] || colMap[board[i][j]] || subMap[int(board[i][j])] {
					return false
				}
				rowMap[board[i][j]] = true
				colMap[board[i][j]] = true
				subMap[(i/3)*3+j/3] = true
			}
		}
	}
	return true
}
