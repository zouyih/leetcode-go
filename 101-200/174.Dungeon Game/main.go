package main

import "fmt"

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	n := len(dungeon[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[m-1][n-1] = max(0, -dungeon[m-1][n-1])

	for i := m - 2; i >= 0; i-- {
		dp[i][n-1] = max(0, dp[i+1][n-1]-dungeon[i][n-1])
	}
	for j := n - 2; j >= 0; j-- {
		dp[m-1][j] = max(0, dp[m-1][j+1]-dungeon[m-1][j])
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			nextMin := min(dp[i+1][j], dp[i][j+1])
			dp[i][j] = max(0, nextMin-dungeon[i][j])
		}
	}

	return dp[0][0] + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	dungeon := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}
	fmt.Println(calculateMinimumHP(dungeon))
}
