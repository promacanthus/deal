package thinking

/* 分治算法
 * 求数组逆序度
 * 在归并排序的合并函数中，统计逆序度
 */

var num int

func DAC(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	dac(nums, 0, n-1)
	return num
}

func dac(nums []int, head, tail int) {
	if head >= tail {
		return
	}

	mid := head + (tail-head)/2

	dac(nums, head, mid)
	dac(nums, mid+1, tail)
	merge(nums, head, mid, tail)
}

func merge(nums []int, head, mid, tail int) {
	i, j, k := head, mid+1, 0
	tmp := make([]int, tail-head+1)

	for i <= mid && j <= tail {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			k++
			i++
		} else {
			num += mid - i + 1
			tmp[k] = nums[j]
			k++
			j++
		}
	}

	for ; i <= mid; i++ {
		tmp[k] = nums[i]
		k++
	}

	for ; j <= tail; j++ {
		tmp[k] = nums[j]
		k++
	}

	for i := 0; i < tail-head; i++ {
		nums[head+i] = tmp[i]
	}
}
