package main

import "fmt"

func wordBreak(s string, wordDict []string) []string {
	hash := make(map[string]bool)
	for _, word := range wordDict {
		hash[word] = true
	}

	mem := make(map[int][]string)
	return wordBreakHelper(s, 0, wordDict, mem)
}

func wordBreakHelper(s string, i int, wordDict []string, mem map[int][]string) []string {
	if i >= len(s) {
		return nil
	}

	if res, ok := mem[i]; ok {
		return res
	}

	res := make([]string, 0)

	for _, word := range wordDict {
		if i+len(word) > len(s) {
			continue
		}

		if s[i:i+len(word)] != word {
			continue
		}

		if i+len(word) == len(s) {
			res = append(res, word)
			continue
		}

		subRes := wordBreakHelper(s, i+len(word), wordDict, mem)
		for _, r := range subRes {
			r = word + " " + r
			res = append(res, r)
		}
	}

	mem[i] = res
	return res
}

func main() {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	res := wordBreak(s, wordDict)
	for _, s := range res {
		fmt.Println(s)
	}
}
