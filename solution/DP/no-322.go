package DP

import "math"

func coinChange(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}
	return coinHelper(coins, amount, make([]int, amount))
}

func coinHelper(coins []int, amount int, dp []int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	if dp[amount-1] != 0 {
		return dp[amount-1]
	}
	min := math.MaxInt64
	for _, coin := range coins {
		res := coinHelper(coins, amount-coin, dp)
		if res >= 0 && res < min {
			min = 1 + res
		}
	}
	if min == math.MaxInt64 {
		dp[amount-1] = -1
	} else {
		dp[amount-1] = min
	}
	return dp[amount-1]
}
