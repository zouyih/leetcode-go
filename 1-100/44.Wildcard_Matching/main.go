package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	n := len(s)
	m := len(p)
	dp := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]bool, m+1)
	}

	dp[0][0] = true
	for j := 1; j < m+1; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			}

			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			}
		}
	}

	return dp[n][m]
}

func main() {
	s := "adceb"
	p := "*a*b"
	fmt.Println(isMatch(s, p))
}
