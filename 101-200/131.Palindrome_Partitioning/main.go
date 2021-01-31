package main

import "fmt"

func partition(s string) [][]string {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	for length := 1; length <= n; length++ {
		for i := 0; i+length <= n; i++ {
			j := i + length - 1
			if i == j {
				dp[i][j] = true
				continue
			}

			if s[i] == s[j] && (length == 2 || dp[i+1][j-1]) {
				dp[i][j] = true
			}
		}
	}

	res := make([][]string, 0)
	path := make([]string, 0)
	dfs(s, 0, &dp, path, &res)
	return res
}

func dfs(s string, i int, dp *[][]bool, path []string, res *[][]string) {
	if i == len(s) {
		*res = append(*res, append([]string{}, path...))
	}

	for j := i; j < len(s); j++ {
		if !(*dp)[i][j] {
			continue
		}
		path = append(path, s[i:j+1])
		dfs(s, j+1, dp, path, res)
		path = path[:len(path)-1]
	}
}

func main() {
	s := "aab"
	fmt.Println(partition(s))
}
