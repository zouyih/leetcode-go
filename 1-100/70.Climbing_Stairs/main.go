package main

import (
	"fmt"
)

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2

	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n-1]
}

func climbStairs1(n int) int {
	if n == 1 {
		return 1
	}

	first := 1
	second := 2

	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}

	return second
}

func main() {
	n := 3
	fmt.Println(climbStairs(n))
	fmt.Println(climbStairs1(n))
}
