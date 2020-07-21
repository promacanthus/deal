package thinking

// 0-1背包问题---回溯及缓存中间状态
import (
	"fmt"
	"math"
)

var (
	w      = 9                    // w 表示背包承重
	weight = []int{2, 2, 4, 6, 3} // weight 表示每个物品的重量
	n      = len(weight)          // n 表示物品个数
	mem    [5][10]bool            // 缓存中间结果
)

func init() {
	// 初始化二维数组
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			mem[i][j] = false
		}
	}
}

// i 表示考察到第i个物品
// cw 表示当前背包重量
func dp01(i, cw int) {
	// maxW 表示限制条件下背包能容纳的最大重量
	maxW := math.MinInt64

	if cw == w || i == n {
		if cw > maxW {
			maxW = cw
		}
		// 输出所有中间状态
		fmt.Println(maxW)
		return
	}

	// 增加二维数组，在内存中记录已经执行过的状态
	{
		if mem[i][cw] {
			return
		}
		mem[i][cw] = true
	}

	// 不放入第i个物品
	dp01(i+1, cw)

	// 没有超过最大重量，放入第i个物品
	if cw+weight[i] <= w {
		dp01(i+1, cw+weight[i])
	}
}
