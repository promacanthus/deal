package DP

import "math"

func maxProduct(nums []int) int {
	maxV, iMax, iMin := math.MinInt64, 1, 1
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			iMax, iMin = iMin, iMax
		}
		iMax = max(iMax*nums[i], nums[i])
		iMin = min(iMin*nums[i], nums[i])
		maxV = max(maxV, iMax)
	}
	return maxV
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
