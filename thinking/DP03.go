package thinking

// 0-1背包问题---动态规划（一维状态转移表）
// 只记录放入物品的状态

// n 表示物品个数
// w 表示背包承重
// weight 表示每个物品的重量
func dp03(items []int, n, w int) int {
	state := make([]bool, w+1)

	// 第一行特殊处理，利用哨兵优化
	if items[0] <= w {
		state[items[0]] = true
	}

	// 一维状态转移表
	for i := 1; i < n; i++ {
		// 把第i个物品放入
		for j := w - items[i]; j >= 0; j-- {
			if state[j] {
				state[j+items[i]] = true
			}
		}
	}

	// 输出结果
	for i := w; i >= 0; i-- {
		if state[i] {
			return i
		}
	}
	return 0
}
