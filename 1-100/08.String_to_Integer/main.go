package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	i := 0
	n := len(s)
	sig := 1

	for i < n && s[i] == ' ' {
		i++
	}

	if i < n {
		if s[i] == '-' {
			sig = -1
			i++
		} else if s[i] == '+' {
			i++
		}
	}

	if i < n && (s[i] < '0' || s[i] > '9') {
		return 0
	}

	res := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		extra := s[i] - '0'
		res = res*10 + int(extra)
		if res*sig >= math.MaxInt32 {
			return math.MaxInt32
		}
		if res*sig <= math.MinInt32 {
			return math.MinInt32
		}
		i++
	}
	return res * sig

}

func main() {
	fmt.Println(myAtoi("  -212"))
}
