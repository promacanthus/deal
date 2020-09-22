package sort

func MergeSort(list []int) []int {
	n := len(list)
	if n < 2 {
		return list
	}
	return mergeSort(list, 0, n-1)
}

func mergeSort(list []int, head, tail int) []int {
	if head >= tail {
		return list
	}

	// 找到数组中点进行划分
	mid := head + (tail-head)/2
	mergeSort(list, head, mid)
	mergeSort(list, mid+1, tail)

	// 将内部有序的两个数组合并为一个
	return merge(list, head, mid, tail)
}

func merge(list []int, head, mid, tail int) []int {
	res := make([]int, 0)

	i, j := head, mid+1

	for i <= mid && j <= tail {
		if list[i] <= list[j] {
			res = append(res, list[i])
			i++
		} else {
			res = append(res, list[j])
			j++
		}
	}

	if i <= mid {
		res = append(res, list[i:mid+1]...)
	}
	if j <= tail {
		res = append(res, list[j:tail]...)
	}
	copy(list[head:tail+1], res)
	return res
}

func MergeSort1(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	// 分治法：divide 分为两段,合并两段数据
	return merge1(MergeSort1(nums[:len(nums)/2]), MergeSort1(nums[len(nums)/2:]))
}
func merge1(left, right []int) (result []int) {
	// 两边数组合并游标
	l := 0
	r := 0
	// 注意不能越界
	for l < len(left) && r < len(right) {
		// 谁小合并谁
		if left[l] > right[r] {
			result = append(result, right[r])
			r++
		} else {
			result = append(result, left[l])
			l++
		}
	}
	// 剩余部分合并
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}
