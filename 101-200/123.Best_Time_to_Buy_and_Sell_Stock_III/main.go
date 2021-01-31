package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	dp := make([][][]int, n+1)
	maxK := 2
	for i := 0; i < n+1; i++ {
		dp[i] = make([][]int, maxK+1)
		for k := 0; k <= maxK; k++ {
			dp[i][k] = make([]int, 2)
		}
	}

	for k := 0; k <= maxK; k++ {
		dp[0][k][1] = math.MinInt32
	}

	for i := 1; i < n+1; i++ {
		for k := 1; k <= maxK; k++ {
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i-1])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i-1])
		}
	}
	return dp[n][2][0]
}

func maxProfit1(prices []int) int {
	dp_i10, dp_i11 := 0, math.MinInt32
	dp_i20, dp_i21 := 0, math.MinInt32

	for i := 0; i < len(prices); i++ {
		dp_i20 = max(dp_i20, dp_i21+prices[i])
		dp_i21 = max(dp_i21, dp_i10-prices[i])
		dp_i10 = max(dp_i10, dp_i11+prices[i])
		dp_i11 = max(dp_i11, -prices[i])
	}
	return dp_i20
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	prices := []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}
	fmt.Println(maxProfit(prices))

	prices = []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(prices))
}
