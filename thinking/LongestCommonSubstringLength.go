package thinking

import "math"

func LCSLBT(a, b string) int {
	aBytes, bBytes := []byte(a), []byte(b)
	return lcslbt(aBytes, bBytes, 0, 0, len(a), len(b), 0)
}

func lcslbt(aBytes, bBytes []byte, i, j, m, n, maxSub int) int {

	// 缓存结果
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	if i == m || j == n {
		return maxSub
	}

	if dp[i][j] > maxSub {
		return dp[i][j]
	}

	dp[i][j] = maxSub

	if aBytes[i] == bBytes[j] {
		return lcslbt(aBytes, bBytes, i+1, j+1, m, n, maxSub+1)
	}
	return twoMaximum(lcslbt(aBytes, bBytes, i+1, j, m, n, maxSub),
		lcslbt(aBytes, bBytes, i, j+1, m, n, maxSub))
}

func twoMaximum(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func LCSLDP(a, b string) int {
	aBytes, bBytes := []byte(a), []byte(b)

	// 初始化二位状态转移表
	dp := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		dp[i] = make([]int, len(b))
	}

	// 	第一行特殊处理
	for i := 0; i < len(a); i++ {
		if aBytes[i] == bBytes[0] {
			dp[i][0] = 1
		} else {
			if i != 0 {
				dp[i][0] = dp[i-1][0]
			} else {
				dp[i][0] = 0
			}
		}
	}
	// 	第一列特殊处理
	for j := 0; j < len(b); j++ {
		if aBytes[0] == bBytes[j] {
			dp[0][j] = 1
		} else {
			if j != 0 {
				dp[0][j] = dp[0][j-1]
			} else {
				dp[0][j] = 0
			}
		}
	}

	for i := 1; i < len(a); i++ {
		for j := 1; j < len(b); j++ {
			if aBytes[i] == bBytes[j] {
				dp[i][j] = threeMaximum(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			} else {
				dp[i][j] = threeMaximum(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}

	return dp[len(a)-1][len(b)-1]
}

func threeMaximum(i, j, k int) int {
	maxV := math.MinInt64
	if i > maxV {
		maxV = i
	}
	if j > maxV {
		maxV = j
	}
	if k > maxV {
		maxV = k
	}
	return maxV
}
