package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var res, length = 0, 0
	var start = -1
	var hash = make(map[int32]int)

	for i, char := range s {
		if pos, ok := hash[char]; ok {
			if start < pos {
				start = pos
			}
		}

		hash[char] = i
		length = i - start

		if length > res {
			res = length
		}
	}
	return res
}

func main() {
	s := "dvdf"
	fmt.Println(lengthOfLongestSubstring(s))
}
