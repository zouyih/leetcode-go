package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}
	digits = append([]int{1}, digits...)

	return digits
}

func main() {
	digits := []int{4, 3, 2, 1}
	fmt.Println(plusOne(digits))
}
