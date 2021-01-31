package main

import (
	"fmt"
)

func minWindow(s string, t string) string {
	need := make(map[int32]int)
	window := make(map[int32]int)

	for _, char := range t {
		need[char]++
	}
	match := 0
	left, right := 0, 0
	d := len(s) + 1
	head := -1
	for right < len(s) {
		char := int32(s[right])
		if need[char] > 0 {
			window[char]++
			if window[char] == need[char] {
				match++
			}
		}
		right++

		for match == len(need) {
			if right-left < d {
				d = right - left
				head = left
			}
			char := int32(s[left])
			if need[char] > 0 {
				window[char]--
				if window[char] < need[char] {
					match--
				}
			}
			left++
		}
	}

	if head == -1 {
		return ""
	}
	return s[head : head+d]
}

func main() {
	S := "ADOBECODEBANC"
	T := "ABC"
	fmt.Println(minWindow(S, T))
}
