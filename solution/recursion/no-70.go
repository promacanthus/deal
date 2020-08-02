package recursion

// recursion + cache
var res = make(map[int]int, 0)

func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	s, ok := res[n]
	if ok {
		return s
	} else {
		res[n] = climbStairs(n-1) + climbStairs(n-2)
	}
	return res[n]
}

// dynamic programming
func climbStairsDP(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
