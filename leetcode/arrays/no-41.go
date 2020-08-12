package arrays

import (
	"sort"
)

func firstMissingPositive(nums []int) int {
	n := len(nums)
	// empty
	if n < 1 {
		return 1
	}

	index := -1
	sort.Ints(nums)

	// remove duplicates
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		res = append(res, nums[i])
	}

	for i, num := range res {
		if num > 0 {
			index = i
			break
		}
	}
	n = len(res)
	// all negative or index value is not 1
	if index < 0 || res[index] != 1 {
		return 1
	}

	return dmp(res, n, index, 2)
}

func dmp(res []int, n, index, point int) int {
	// traverse to the end
	if index == n-1 {
		return res[index] + 1
	}
	// somewhere in the middle
	if res[index+1] > point {
		return point
	}
	return dmp(res, n, index+1, point+1)
}

func firstMissingPositive2(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
