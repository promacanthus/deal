package sort

func SelectionSort(list []int) {
	n := len(list)
	if n <= 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			// 查找最小的元素
			if list[min] > list[j] {
				min = j // 记录最小值
			}
		}
		list[i], list[min] = list[min], list[i]
	}
}
