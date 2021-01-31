package main

import "fmt"

func numTrees(n int) int {
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for v := 1; v <= i; v++ {
			dp[i] += dp[v-1] * dp[i-v]
		}
	}

	return dp[n]
}

func main() {
	n := 3
	fmt.Println(numTrees(n))
}
