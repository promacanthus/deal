package sort

// 数组实现的选择排序是不稳定的

// 如待排序[]int{5,5,2}，选择排序的结果是[]int{2,5,5}
// 其中两个5的前后顺序变了

func SelectionSort(list []int) {
	n := len(list)
	if n <= 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		// 外层循环控制待交换的位置
		min := i
		// 从第二个元素开始，查找最小的元素，与i交换
		for j := i + 1; j < n; j++ {
			// 查找最小的元素
			if list[min] > list[j] {
				min = j // 记录最小值
			}
		}
		list[i], list[min] = list[min], list[i]
	}
}
