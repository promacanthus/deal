package example

// 最小路径和

import "math"

// 分治
func minPathSum1(grid [][]int) int {
	return int(cal(grid, 0, 0))
}

func cal(grid [][]int, i, j int) float64 {
	// 走到最边上了
	if i == len(grid) || j == len(grid[0]) {
		return math.MaxInt64
	}

	// 走到右下角了
	if i == len(grid)-1 && j == len(grid[0])-1 {
		return float64(grid[i][j])
	}

	return float64(grid[i][j]) + math.Min(cal(grid, i+1, j), cal(grid, i, j+1))
}
