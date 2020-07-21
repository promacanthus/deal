package thinking

import (
	"fmt"
)

// 八皇后问题

// 数组下标表示行
// 数组值表示queen存储的列
var result = [8]int{}

func cal8queens(row int) {
	if row == 8 {
		printQueens(result)
		return
	}
	// 每一行八种放法
	for column := 0; column < 8; column++ {
		// 去除不符合摆放规则的位置
		if isOK(row, column) {
			// 存储queen位置，第row行，column列
			result[row] = column
			cal8queens(row + 1)
		}
	}
}

func isOK(row int, column int) bool {
	leftUp, rightUp := column-1, column+1
	for i := row - 1; i >= 0; i-- {
		if leftUp >= 0 {
			if result[i] == leftUp {
				return false
			}
		}
		if rightUp < 8 {
			if result[i] == rightUp {
				return false
			}
		}
		leftUp--
		rightUp++
	}
	return true
}

func printQueens(result [8]int) {
	for row := 0; row < 8; row++ {
		for column := 0; column < 8; column++ {
			if result[row] == column {
				fmt.Print("Q ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
