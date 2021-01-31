package main

import (
	"fmt"
)

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	pre, cur := 1, 1
	for i := 1; i < n; i++ {
		num := s[i] - '0'
		lastNum := s[i-1] - '0'

		tmp := cur
		if num == 0 {
			if lastNum == 1 || lastNum == 2 {
				cur = pre
			} else {
				return 0
			}
		} else if lastNum == 1 || (lastNum == 2 && num <= 6) {
			cur = pre + cur
		}
		pre = tmp
	}
	return cur
}

func numDecodings1(s string) int {
	if len(s) == 0 {
		return 0
	}
	memorization := make(map[int]int)
	return decodeCounts(0, s, memorization)
}

func decodeCounts(index int, s string, memorization map[int]int) int {
	if index == len(s) {
		return 1
	}

	if s[index] == '0' {
		return 0
	}

	if count, ok := memorization[index]; ok {
		return count
	}
	res := decodeCounts(index+1, s, memorization)
	if index < len(s)-1 {
		num1 := s[index] - '0'
		num2 := s[index+1] - '0'
		if num1 == 1 || (num1 == 2 && num2 <= 6) {
			res += decodeCounts(index+2, s, memorization)
		}
	}
	memorization[index] = res
	return res
}

func main() {
	s := "226"
	fmt.Println(numDecodings(s))
	fmt.Println(numDecodings1(s))
}
