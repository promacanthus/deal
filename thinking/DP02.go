package thinking

// 0-1背包问题---动态规划（二维状态转移表）
// 通过当前阶段的状态集合推导下一个阶段的状态集合

var states = [5][10]bool{} // 状态转移表

func init() {
	//  初始化状态转移表
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			states[i][j] = false
		}
	}
}

// n 表示物品个数
// w 表示背包承重
// weight 表示每个物品的重量
func dp02(weight []int, n, w int) int {

	// 第一行特殊处理，利用哨兵优化
	{
		// 第一个物品不放入
		states[0][0] = true

		// 第一个物品放入
		if weight[0] <= w {
			states[0][weight[0]] = true
		}
	}

	// 动态规划状态转移
	for i := 1; i < n; i++ {
		// 不放入第i个物品
		for j := 0; j <= w; j++ {
			if states[i-1][j] {
				states[i][j] = states[i-1][j]
			}
		}

		// 放入第i个物品
		for j := 0; j <= w-weight[i]; j++ {
			if states[i-1][j] {
				states[i][j+weight[i]] = true
			}
		}
	}

	// 输出结果，也就是状态转移表的右下角
	for i := w; i >= 0; i-- {
		if states[n-1][i] {
			return i
		}
	}
	return 0
}
