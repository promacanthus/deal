package DP

import "math"

func maxProfit(prices []int) int {
	n := len(prices)
	if prices == nil || n <= 1 {
		return 0
	}
	max := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if prices[i] < prices[j] {
				tmp := prices[j] - prices[i]
				if tmp > max {
					max = tmp
				}
			}
		}
	}
	return max
}

func maxProfitOne(prices []int) int {
	n := len(prices)
	if prices == nil || n <= 1 {
		return 0
	}

	minPrice := math.MaxInt32
	max := 0
	for i := 0; i < n; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > max {
			max = prices[i] - minPrice
		}
	}
	return max
}
