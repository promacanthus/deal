package DP

import "math"

// 最小路径和

// 回溯+缓存
func minPathSum(grid [][]int) int {
	return helper(grid, 0, 0)
}

func helper(grid [][]int, i, j int) int {
	if i == len(grid) || j == len(grid[0]) {
		return math.MaxInt64
	}

	if i == len(grid)-1 && j == len(grid[0])-1 {
		return grid[i][j]
	}
	return grid[i][j] + min(helper(grid, i, j+1), helper(grid, i+1, j))
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

// 动态规划
func minPathSumDP(grid [][]int) int {
	iLength := len(grid)
	if iLength < 1 {
		return 0
	}
	jLength := len(grid[0])

	// 初始化状态转移表
	dp := make([][]int, iLength)
	for i, state := range grid {
		dp[i] = make([]int, len(state))
	}

	// 第一行特殊处理
	dp[0][0] = grid[0][0]

	// 	动态规划
	for i := 0; i < iLength; i++ {
		for j := 0; j < jLength; j++ {
			if i == 0 && j != 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if i != 0 && j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else if i != 0 && j != 0 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[iLength-1][jLength-1]
}
