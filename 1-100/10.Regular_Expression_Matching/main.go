package main

import "fmt"

func isMatch(s string, p string) bool {
	n := len(s)
	m := len(p)

	dp := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]bool, m+1)
	}

	dp[0][0] = true
	for j := 1; j < m+1; j++ {
		dp[0][j] = j > 1 && p[j-1] == '*' && dp[0][j-2]
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if p[j-1] == '.' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			}

			if p[j-1] == '*' {
				if dp[i][j-2] {
					dp[i][j] = true
				}

				if (p[j-2] == s[i-1] || p[j-2] == '.') && dp[i-1][j] {
					dp[i][j] = true
				}
			}
		}
	}

	return dp[n][m]
}

func main() {
	s := "aaa"
	p := "ab*a*c*a"
	fmt.Println(isMatch(s, p))
}
