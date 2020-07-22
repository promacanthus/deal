package thinking

// 计算莱文斯特距离

import "math"

var minDist = math.MaxInt64

func LDBT(a, b string) int {
	if len(a) < 1 {
		return len(a)
	}
	if len(b) < 1 {
		return len(b)
	}

	ldbt([]byte(a), []byte(b), 0, 0, len(a), len(b), 0)
	return minDist
}

func ldbt(aBytes []byte, bBytes []byte, i, j, m, n, edit int) {
	if i == m || j == n {
		if i < m {
			edit += m - i
		}
		if j < n {
			edit += n - j
		}
		if edit < minDist {
			minDist = edit
		}
		return
	}

	if aBytes[i] == bBytes[i] {
		ldbt(aBytes, bBytes, i+1, j+1, m, n, edit)
	} else {
		ldbt(aBytes, bBytes, i+1, j, m, n, edit+1)   // 删除a[i]或者b[j]前添加一个字符
		ldbt(aBytes, bBytes, i, j+1, m, n, edit+1)   // 删除b[j]或者a[i]前添加一个字符
		ldbt(aBytes, bBytes, i+1, j+1, m, n, edit+1) // 将a[i]和b[j]替换为相同字符
	}
}

func LDDP(a, b string) int {
	// 初始化二维状态转移表
	dp := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		dp[i] = make([]int, len(b))
	}

	aBytes := []byte(a)
	bBytes := []byte(b)

	// 第一行特殊处理
	for i := 0; i < len(a); i++ {
		if aBytes[i] == bBytes[0] {
			dp[i][0] = i
		} else {
			if i != 0 {
				dp[i][0] = dp[i-1][0] + 1
			} else {
				dp[i][0] = 1
			}
		}
	}

	// 第一列特殊处理
	for j := 0; j < len(b); j++ {
		if aBytes[0] == bBytes[j] {
			dp[0][j] = j
		} else {
			if j != 0 {
				dp[0][j] = dp[0][j-1] + 1
			} else {
				dp[0][j] = 1
			}
		}
	}

	for i := 1; i < len(a); i++ {
		for j := 1; j < len(b); j++ {
			if aBytes[i] == bBytes[j] {
				dp[i][j] = minimum(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1])
			} else {
				dp[i][j] = minimum(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}

	return dp[len(a)-1][len(b)-1]
}

func minimum(i, j, k int) int {
	minV := math.MaxInt64

	if i < minV {
		minV = i
	}
	if j < minV {
		minV = j
	}
	if k < minV {
		minV = k
	}
	return minV
}
