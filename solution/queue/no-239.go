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

func maxSlidingWindowIteration(nums []int, k int) []int {
	n := len(nums)
	if n*k == 0 {
		return nil
	}
	if k == 1 {
		return nums
	}

	res := make([]int, 0)
	for i := 0; i < n-k+1; i++ {
		maxV := math.MinInt64
		for j := i; j < i+k; j++ {
			maxV = max(maxV, nums[j])
		}
		res = append(res, maxV)
	}
	return res
}

func max(maxV, value int) int {
	if value > maxV {
		return value
	}
	return maxV
}

func maxSlidingWindowDP(nums []int, k int) []int {
	n := len(nums)
	if n*k == 0 {
		return nil
	}
	if k == 1 {
		return nums
	}

	left := make([]int, 0)
	left = append(left, nums[0])

	right := make([]int, n)
	right[n-1] = nums[n-1]

	for i := 1; i < n; i++ {
		if i%k == 0 {
			left = append(left, nums[i])
		} else {
			left = append(left, max(left[i-1], nums[i]))
		}

		j := n - i - 1
		if (j+1)%k == 0 {
			right[j] = nums[j]
		} else {
			right[j] = max(right[j+1], nums[j])
		}
	}

	res := make([]int, 0)
	for i := 0; i < n-k+1; i++ {
		res = append(res, max(left[i+k-1], right[i]))
	}
	return res
}
