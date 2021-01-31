package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i < n+1; i++ {
		for _, word := range wordDict {
			if len(word) > i {
				continue
			}
			if dp[i-len(word)] && s[i-len(word):i] == word {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

func wordBreak1(s string, wordDict []string) bool {
	mem := make(map[int]bool)
	return dfs(s, 0, wordDict, mem)
}

func dfs(s string, i int, wordDict []string, mem map[int]bool) bool {
	if i == len(s) {
		return true
	}

	if mem[i] {
		return false
	}

	for _, word := range wordDict {
		if i+len(word) > len(s) {
			continue
		}
		if s[i:i+len(word)] != word {
			continue
		}
		if dfs(s, i+len(word), wordDict, mem) {
			return true
		}
	}

	mem[i] = true
	return false
}

func main() {
	s := "applepenapple"
	wordDict := []string{"apple", "pen"}
	fmt.Println(wordBreak(s, wordDict))
}
