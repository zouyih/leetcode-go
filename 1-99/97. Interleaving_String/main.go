package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	m := len(s1)
	n := len(s2)
	if m+n != len(s3) {
		return false
	}

	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i < m+1; i++ {
		if s1[i-1] == s3[i-1] {
			dp[i][0] = true
		} else {
			break
		}
	}

	for j := 1; j < n+1; j++ {
		if s2[j-1] == s3[j-1] {
			dp[0][j] = true
		} else {
			break
		}
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if dp[i-1][j] && s1[i-1] == s3[i+j-1] {
				dp[i][j] = true
			}

			if dp[i][j-1] && s2[j-1] == s3[i+j-1] {
				dp[i][j] = true
			}
		}
	}

	return dp[m][n]
}

func isInterleave1(s1 string, s2 string, s3 string) bool {
	m := len(s1)
	n := len(s2)
	if m+n != len(s3) {
		return false
	}

	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}

	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
			} else if j == 0 {
				dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
			} else {
				dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}

	return dp[m][n]
}

func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	fmt.Println(isInterleave(s1, s2, s3))

	s1 = "aabcc"
	s2 = "dbbca"
	s3 = "aadbbbaccc"
	fmt.Println(isInterleave(s1, s2, s3))
}
