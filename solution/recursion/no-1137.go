package recursion

var tmp = make(map[int]int, 0)

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	ans, ok := tmp[n]
	if ok {
		return ans
	}
	tmp[n] = tribonacci(n-3) + tribonacci(n-2) + tribonacci(n-1)
	return tmp[n]
}

func tribonacciIteration(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	T := make([]int, n+1)
	T[0] = 0
	T[1] = 1
	T[2] = 1
	for i := 3; i <= n; i++ {
		T[i] = T[i-3] + T[i-2] + T[i-1]
	}
	return T[n]
}
