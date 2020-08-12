package DP

import "math"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if triangle == nil || n == 0 {
		return 0
	}

	sum := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			sum[i] = make([]int, len(triangle[i]))
		}
	}

	sum[0][0] = triangle[0][0]

	for i := 1; i < n; i++ {
		sum[i][0] = sum[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			sum[i][j] = min(sum[i-1][j], sum[i-1][j-1]) + triangle[i][j]
		}
		sum[i][i] = sum[i-1][i-1] + triangle[i][i]
	}
	res := math.MaxInt64
	for i := 0; i < n; i++ {
		res = min(res, sum[n-1][i])
	}
	return res
}
