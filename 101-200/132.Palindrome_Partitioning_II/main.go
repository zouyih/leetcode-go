package main

import "fmt"

func minCut(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	for length := 1; length <= n; length++ {
		for i := 0; i+length <= n; i++ {
			j := i + length - 1
			if s[i] == s[j] && (length <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
			}
		}
	}

	num := 0
	memo := make(map[int]int)
	res := minCutHelper(s, 0, num, dp, memo)

	return res
}

func minCutHelper(s string, i, num int, dp [][]bool, memo map[int]int) int {
	n := len(s)
	if dp[i][n-1] {
		return 0
	}

	if memNum, ok := memo[i]; ok {
		return memNum
	}

	minCount := n - i
	for j := i; j < n; j++ {
		if !dp[i][j] {
			continue
		}

		minCount = min(minCount, 1+minCutHelper(s, j+1, num+1, dp, memo))
	}
	memo[i] = minCount
	return minCount
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func main() {
	s := "aab"
	fmt.Println(minCut(s))
}
