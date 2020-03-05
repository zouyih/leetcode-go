package main

import (
	"fmt"
)

func mySqrt(x int) int {
	left := 0
	right := x
	for left <= right {
		mid := (left + right) / 2
		product := mid * mid
		if product == x {
			return mid
		}
		if product > x {
			right = mid - 1
		}
		if product < x {
			left = mid + 1
		}
	}

	return right
}

func main() {
	x := 8
	fmt.Println(mySqrt(x))
}
