package binarysearch

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	i := 0
	for i*i < x {
		i++
	}

	if i*i == x {
		return i
	} else {
		return i - 1
	}
}

func mySqrtBS(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + ((r - l) >> 1)
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
