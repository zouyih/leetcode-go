package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}

	dp := []int{triangle[0][0]}
	for i := 1; i < n; i++ {
		tmp := make([]int, i+1)
		tmp[0] = dp[0] + triangle[i][0]
		tmp[i] = dp[i-1] + triangle[i][i]
		for j := 1; j < i; j++ {
			tmp[j] = min(dp[j-1], dp[j]) + triangle[i][j]
		}
		dp = tmp
	}
	res := dp[0]
	for i := 1; i < n; i++ {
		res = min(res, dp[i])
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func main() {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(triangle))
}
