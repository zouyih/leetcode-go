package main

import (
	"fmt"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])

	dp := make([][]int, m)
	num := 1
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if obstacleGrid[i][0] == 1 {
			num = 0
		}
		dp[i][0] = num
	}
	num = 1
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			num = 0
		}
		dp[0][j] = num
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}

	return dp[m-1][n-1]
}

func main() {
	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}
