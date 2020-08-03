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
	ret := make([]int, tail-head+1)

	i, j, k := head, mid+1, 0

	for ; i <= mid && j <= tail; k++ {
		if list[i] <= list[j] {
			ret[k] = list[i]
			i++
		} else {
			ret[k] = list[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		ret[k] = list[i]
		k++
	}

	for ; j <= tail; j++ {
		ret[k] = list[j]
		k++
	}
	copy(list[head:tail+1], ret)
	return ret
}
