package sort

func QuickSort(list []int) []int {
	n := len(list)
	if n < 2 {
		return list
	}
	return quickSort(list, 0, n-1)
}

func quickSort(list []int, head, tail int) []int {
	if head >= tail {
		return list
	}

	pivot := partition(list, head, tail)
	quickSort(list, head, pivot-1)
	quickSort(list, pivot+1, tail)
	return list
}

func partition(list []int, head int, tail int) int {
	// pivot的选择是可优化的点
	pivot := list[tail] // pivot的取值
	i := head           // i表示pivot的位置
	for j := head; j < tail; j++ {
		if list[j] < pivot {
			list[i], list[j] = list[j], list[i]
			i++
		}
	}
	// 遍历一遍数组，说明i之前的都是小于pivot
	// i停在哪，pivot就在哪，交换i与pivot的值
	list[i], list[tail] = list[tail], list[i]
	return i
}
