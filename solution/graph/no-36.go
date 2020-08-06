package graph

// func isValidSudoku(board [][]byte) bool {
// 	for i := 0; i < 9; i++ {
// 		rows := make(map[byte]bool)
// 		columns := make(map[byte]bool)
// 		cube := make(map[byte]bool)
// 		for j := 0; j < 9; j++ {
// 			if board[i][j] != '.' && rows[board[i][j]] {
// 				return false
// 			}
// 			if board[j][i] != '.' && columns[board[j][i]] {
// 				return false
// 			}
// 			rowIndex := 3 * (i / 3)
// 			colIndex := 3 * (i % 3)
// 			if board[rowIndex+j/3][colIndex+j%3] != '.' && cube[board[rowIndex+j/3][colIndex+j%3]] {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		row, col, cube := 0, 0, 0
		for j := 0; j < 9; j++ {
			r := int(board[i][j])
			c := int(board[j][i])
			cu := int(board[3*(i/3)+j/3][3*(i%3)+j%3])
			if r > 0 {
				row = sdk(r, row)
			}
			if c > 0 {
				col = sdk(c, col)
			}
			if cu > 0 {
				cube = sdk(cu, cube)
			}
			if row == -1 || col == -1 || cube == -1 {
				return false
			}
		}
	}
	return true
}

func sdk(n, val int) int {
	if (val>>n)&1 == 1 {
		return -1
	}
	return val ^ (1 << n)
}
