package main

import "fmt"

func longestPalindrome(s string) string {
	n := len(s)

	if n <= 1 {
		return s
	}

	var start, right, left int
	var maxLen, maxLenStart int
	for i := 0; i < len(s); i++ {
		start = i
		for i+1 < n && s[i] == s[i+1] {
			i++
		}
		right = i
		left = start
		for left > 0 && right+1 < n && s[left-1] == s[right+1] {
			left--
			right++
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1
			maxLenStart = left
		}
	}

	res := s[maxLenStart : maxLenStart+maxLen]
	return res
}

func main() {
	s := "aba"
	fmt.Println(longestPalindrome(s))
}
