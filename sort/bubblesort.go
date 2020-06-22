package sort

func BubbleSort(list []int) {
	n := len(list)
	if n <= 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		flag := false // 提前退出冒泡循环的标志位
		for j := i + 1; j < n; j++ {
			if list[i] > list[j] { // 比较并交换位置，加上等号算法会变成不稳定排序
				tmp := list[i]
				list[i] = list[j]
				list[j] = tmp
				flag = true // 数据有交换时设置标志位
			}
		}
		if !flag {
			break
		}
	}
}
