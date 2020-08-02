package Recurison

func tribonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	if n == 2 {
		return 1
	}

	mid := make(map[int]int)

	mid[0] = 0
	mid[1] = 1
	mid[2] = 1

	v3, ok3 := mid[n-3]
	v2, ok2 := mid[n-2]
	v1, ok1 := mid[n-1]

	if ok3 {
		if ok2 {
			if ok1 {
				return v3 + v2 + v1
			} else {
				mid[n-1] = tribonacci(n - 1)
				return v3 + v2 + mid[n-1]
			}
		} else {
			if ok1 {
				mid[n-2] = tribonacci(n - 2)
				return v3 + mid[n-2] + v1
			} else {
				mid[n-2] = tribonacci(n - 2)
				mid[n-1] = tribonacci(n - 1)
				return v3 + mid[n-2] + mid[n-1]
			}
		}
	} else {
		if ok2 {
			if ok1 {
				mid[n-3] = tribonacci(n - 3)
				return mid[n-3] + v2 + v1
			} else {
				mid[n-3] = tribonacci(n - 3)
				mid[n-1] = tribonacci(n - 1)
				return mid[n-3] + v2 + mid[n-1]
			}
		} else {
			if ok1 {
				mid[n-3] = tribonacci(n - 3)
				mid[n-2] = tribonacci(n - 2)
				return mid[n-3] + mid[n-2] + v1
			} else {
				mid[n-3] = tribonacci(n - 3)
				mid[n-2] = tribonacci(n - 2)
				mid[n-1] = tribonacci(n - 1)
				return mid[n-3] + mid[n-2] + mid[n-1]
			}
		}
	}
}
