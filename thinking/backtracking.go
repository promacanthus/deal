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

// 0-1背包问题---回溯

var (
	maxW = 0
	mem  = [5][10]bool{}
)

// cw 表示当前重量
// i 表示考察到第i个物品
// w 表示背包承重
// items 表示每个物品的重量
// n 表示物品个数
func packages(i, cw int, items []int, n, w int) {
	// cw==w 表示背包满了
	// i==n 表示物品考察完了
	if cw == w || i == n {
		if cw > maxW {
			maxW = cw
			return
		}
	}
	if mem[i][cw] {
		return
	}
	mem[i][cw] = true
	packages(i+1, cw, items, n, w)
	// 超过背包最大承重，则停止放入
	if cw+items[i] <= w {
		packages(i+1, cw+items[i], items, n, w)
	}
}
