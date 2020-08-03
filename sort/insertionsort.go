package sort

func InsertionSort(list []int) {
	n := len(list)
	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		// 外层循环控制待插入的元素
		val := list[i] // 待插入元素
		j := i - 1     // 当前有序数组的最后一个索引位置
		for ; j >= 0; j-- {
			// 内层循环从后往前遍历已有序元素，比较并寻找待插入位置
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
