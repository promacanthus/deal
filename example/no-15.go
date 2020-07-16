package example

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 3 {
		return res
	}

	sortB(nums)
	sortI(nums)
	sortS(nums)

	return res
}

// 冒泡
func sortB(nums []int) {
	n := len(nums)
	sortBorder := n - 1
	for i := 0; i < n-1; i++ {
		flag := false
		for j := 0; j < sortBorder; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

// 插入
func sortI(nums []int) {
	for i := 1; i < len(nums); i++ {
		value := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if value < nums[j] {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = value
	}
}

// 选择
func sortS(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
}
