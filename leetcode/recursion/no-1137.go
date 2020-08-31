package recursion

// recursion & cache
var tmp = make(map[int]int)

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	if ans, ok := tmp[n]; ok {
		return ans
	}
	tmp[n] = tribonacci(n-3) + tribonacci(n-2) + tribonacci(n-1)

	return tmp[n]
}

// dynamic programming
func tribonacciDP(n int) int {
	if n == 0 {
		return 0
	}
	// n<3 dp数组会初始化失败
	if n < 3 {
		return 1
	}

	DP := make([]int, n+1)
	DP[0] = 0
	DP[1] = 1
	DP[2] = 1
	for i := 3; i <= n; i++ {
		DP[i] = DP[i-3] + DP[i-2] + DP[i-1]
	}
	return DP[n]
}
