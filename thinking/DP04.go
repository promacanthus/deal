package thinking

import (
	"fmt"
	"math"
)

// 0-1背包问题升级版，增加物品价值维度

var (
	items  = []int{2, 2, 4, 6, 3}
	values = []int{3, 4, 8, 9, 6}
	maxV   = math.MinInt64
	w2     = 9
	n2     = len(items)
	vMem   = [5][10]int{}
)

func init() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			vMem[i][j] = -1
		}
	}
}

// i 表示考察到第i个物品
// cw 表示当前背包重量
// w2 表示背包最大承重
// cv 表示当前背包中物品的总价值
// n 表示物品个数
func dp04(i, cw, cv int) {
	if cw == w2 || i == n2 {
		if cv > maxV {
			maxV = cv
		}
		fmt.Println(maxV)
		return
	}

	//  记录中间状态
	{
		if vMem[i][cw] > cv {
			return
		}
		vMem[i][cw] = cv
	}

	// 不放入物品i
	dp04(i+1, cw, cv)

	// 	放入物品i
	if cw+weight[i] <= w2 {
		dp04(i+1, cw+items[i], cv+values[i])
	}
}
