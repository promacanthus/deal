package queue

import (
	"math"
	"sort"
)

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n < 1 || k < 1 {
		return nil
	}

	if k == 1 {
		return nums
	}

	res := make([]int, 0)
	if len(nums) <= k {
		sort.Ints(nums)
		res = append(res, nums[n-1])
		return res
	}

	queue := make([]int, 0)
	for i := 0; i < k; i++ {
		queue = append(queue, nums[i])
	}
	return maxsw(nums, queue, res, k-1, n)
}

func maxsw(nums, queue, res []int, tail, n int) []int {
	if tail == n {
		return res
	}

	ans := findMaxAndIndex(queue)
	res = append(res, ans)

	if tail == n-1 {
		return res
	}

	queue = queue[1:]
	queue = append(queue, nums[tail+1])

	return maxsw(nums, queue, res, tail+1, n)
}

func findMaxAndIndex(queue []int) int {
	maxAns := math.MinInt64
	for i := 0; i < len(queue); i++ {
		if queue[i] > maxAns {
			maxAns = queue[i]
		}
	}
	return maxAns
}
