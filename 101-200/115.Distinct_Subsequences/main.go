package main

import "fmt"

func numDistinct(s string, t string) int {
	m := len(s)
	n := len(t)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] += dp[i-1][j-1]
			}
			dp[i][j] += dp[i-1][j]
		}
	}

	return dp[m][n]
}

func numDistinct1(s string, t string) int {
	res := 0
	hash := make(map[string]int)
	dfs(s, t, 0, 0, &res, hash)
	return res
}

func dfs(s string, t string, i, j int, res *int, hash map[string]int) {
	if j == len(t) {
		*res++
		return
	}
	if i == len(s) {
		return
	}

	key := string(i+'0') + "@" + string(j+'0')
	if count, ok := hash[key]; ok {
		*res += count
		return
	}

	preRes := *res
	if s[i] == t[j] {
		dfs(s, t, i+1, j+1, res, hash)
	}
	dfs(s, t, i+1, j, res, hash)

	hash[key] = *res - preRes
}

func main() {
	s := "rabbbit"
	t := "rabbit"
	fmt.Println(numDistinct(s, t))

	s = "babgbag"
	t = "bag"
	fmt.Println(numDistinct(s, t))
}
