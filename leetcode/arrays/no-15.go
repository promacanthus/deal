package arrays

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for i := 0; i < n; i++ {
		// 需要和上一次枚举的数不相同
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -1 * nums[i]
		// 枚举 b
		// c 对应的指针初始指向数组的最右端
		for j, k := i+1, n-1; j < n; j++ {
			// 需要和上一次枚举的数不相同
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for j < k && nums[j]+nums[k] > target {
				k--
			}
			// 如果指针重合，随着 b 后续的增加，就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if j == k {
				break
			}
			if nums[j]+nums[k] == target {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return ans
}
