package sort

func InsertionSort(list []int) {
	n := len(list)
	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		val := list[i]
		j := i - 1
		for ; j >= 0; j-- {
			// 数据比较，保证算法稳定相等时不交换
			if val < list[j] {
				// 数据移动
				list[j+1] = list[j]
			} else {
				break
			}
		}
		// 数据插入
		list[j+1] = val
	}
}
