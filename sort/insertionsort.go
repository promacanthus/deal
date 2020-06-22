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
			if val < list[j] { // 数据比较，加上等号算法会变成不稳定排序
				list[j+1] = list[j] // 数据移动
			} else {
				break
			}
		}
		list[j+1] = val // 数据插入
	}
}
