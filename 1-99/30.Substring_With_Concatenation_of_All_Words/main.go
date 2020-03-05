package main

import "fmt"

func findSubstring(s string, words []string) []int {
	res := make([]int, 0)
	if len(words) == 0 {
		return res
	}

	wordLen := len(words[0])
	nWords := len(words)

	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word] += 1
	}

	for i := 0; i < wordLen; i++ {
		left, right := i, i
		countMap := make(map[string]int)
		count := 0

		for right+wordLen <= len(s) {
			w := s[right : right+wordLen]
			countMap[w] += 1
			right += wordLen
			count++

			for countMap[w] > wordMap[w] {
				countMap[s[left:left+wordLen]] -= 1
				left += wordLen
				count--
			}

			if count == nWords {
				res = append(res, left)
			}
		}
	}
	return res
}

func main() {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}

	fmt.Println(findSubstring(s, words))
}
