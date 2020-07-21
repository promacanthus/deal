package thinking

func dp05(weight, value []int, n, w int) int {
	// 假设总共物品数5个，背包最大承重9
	state := [5][10]int{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			state[i][j] = -1
		}
	}

	// 第一行特殊处理
	state[0][0] = 0
	if weight[0] < w {
		state[0][weight[0]] = value[0]
	}

	for i := 1; i < n; i++ {
		// 不放入i物品
		for j := 0; j < w; j++ {
			if state[i-1][j] >= 0 {
				state[i][j] = state[i-1][j]
			}
		}
		// 	放入i物品
		for j := 0; j < w-weight[j]; j++ {
			if state[i-1][j] >= 0 {
				v := state[i-1][j] + value[i]
				if v > state[i][j+weight[i]] {
					state[i][j+weight[i]] = v
				}
			}
		}
	}

	maxValue := -1
	for i := 0; i <= w; i++ {
		if state[n-1][i] > maxValue {
			maxValue = state[n-1][i]
		}
	}

	return maxValue
}
