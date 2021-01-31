package main

import (
	"fmt"
	"math"
)

func maxProfit(maxK int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([][][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([][]int, maxK+1)
		for j := 0; j <= maxK; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	dp[0][0][1] = math.MinInt32
	for j := 1; j <= maxK; j++ {
		dp[0][j][1] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		for j := 1; j <= maxK; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[len(prices)-1][maxK][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	maxK := 2
	prices := []int{2, 4, 1}
	fmt.Println(maxProfit(maxK, prices))

	prices = []int{3, 2, 6, 5, 0, 3}
	fmt.Println(maxProfit(maxK, prices))
}
