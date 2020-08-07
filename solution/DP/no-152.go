package DP

import "math"

func maxProduct(nums []int) int {
	n := len(nums)
	if n < 2 {
		return nums[0]
	}
	max := math.MinInt64
	for window := 1; window <= n; window++ {
		for i := 0; i < n-window+1; i++ {
			start := i
			end := start + window
			sub := nums[start:end]
			tmpMax := 1
			for j := 0; j < len(sub); j++ {
				tmpMax *= sub[j]
			}
			if tmpMax > max {
				max = tmpMax
			}
		}
	}
	return max
}
