package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	res := ""
	s = strings.Trim(s, " ")
	start := len(s) - 1
	end := len(s)

	for start >= 0 {
		if s[start] == ' ' {
			res += s[start+1:end] + " "
			for s[start] == ' ' {
				start--
			}
			end = start + 1
		}
		start--
	}
	res += s[:end]
	return res
}

func main() {
	s := "   the sky is blue   "
	fmt.Println(reverseWords(s))
}
