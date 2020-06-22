package sort

func SelectionSort(list []int) {
	n := len(list)
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		min := list[i]
		index := 0
		flag := false
		for j := i + 1; j < n; j++ {
			if min > list[j] { // 查找最小的元素
				min = list[j] // 记录最小值
				index = j     // 记录最小值的索引
				flag = true   // 不交换数据说明已经有序
			}
		}
		if flag {
			list[index] = list[i] // 移动数据空出位置
			list[i] = min         // 放入最小值
		}
	}
}
