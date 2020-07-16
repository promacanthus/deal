package sort

func BubbleSort(list []int) {
	n := len(list)
	if n <= 1 {
		return
	}

	// 最后一次交换的索引位置
	// 无序数据的边界，每次只需要比较到这里即可退出
	sortBorder := n - 1
	for i := 0; i < n-1; i++ {
		// 提前退出冒泡循环的标志位
		flag := false
		for j := 0; j < sortBorder; j++ {
			// 比较并交换位置，保证算法稳定相等时不交换
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				// 数据有交换时设置标志位
				flag = true
				sortBorder = j
			}
		}
		if !flag {
			break
		}
	}
}
