package arrays

import (
	"math"
	"sort"
)

func majorityElement(nums []int) int {
	n := len(nums)
	mind := n / 2
	state := make(map[int]int, n)

	for i := 0; i < n; i++ {
		state[nums[i]] += 1
		if state[nums[i]] > mind {
			return nums[i]
		}
	}
	return 0
}

func majorityElementSort(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func majorityElementMoore(nums []int) int {
	count := 0
	candidate := math.MaxInt64
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count += 1
		} else {
			count += -1
		}
	}
	return candidate
}
