package arrays

func twoSum(nums []int, target int) []int {
	n := len(nums)
	if n < 1 {
		return nil
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func twoSumHash(nums []int, target int) []int {
	n := len(nums)
	if n < 1 {
		return nil
	}
	state := make(map[int]int)
	for i := 0; i < n; i++ {
		j, ok := state[target-nums[i]]
		state[nums[i]] = i
		if ok {
			return []int{j, i}
		}
	}
	return nil
}
