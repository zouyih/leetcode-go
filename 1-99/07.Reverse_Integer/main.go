package main

import (
	"fmt"
)

func reverse(x int) int {
	const (
		MaxInt32 = 1<<31 - 1
		MinInt32 = -1 << 31
	)

	sig := 1
	if x < 0 {
		sig = -1
		x = -x
	}

	res := 0
	for x > 0 {
		res = res*10 + x%10
		x = x / 10
	}

	res = res * sig
	if res > MaxInt32 || res < MinInt32 {
		return 0
	}
	return res
}

func main() {
	fmt.Println(reverse(1534236469))
	fmt.Println(-1 % 10)
}
