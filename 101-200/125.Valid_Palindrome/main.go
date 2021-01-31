package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if !isValid(s[left]) {
			left++
			continue
		}
		if !isValid(s[right]) {
			right--
			continue
		}
		if strings.ToLower(string(s[left])) == strings.ToLower(string(s[right])) {
			left++
			right--
			continue
		}
		return false
	}
	return true
}

func isValid(a byte) bool {
	if (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9') {
		return true
	}
	return false
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
