package main

import "fmt"

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	res := 0
	for left < right {
		hl := height[left]
		hr := height[right]
		length := right - left

		if hl < hr {
			res = maxInt(hl*length, res)
			left++
		} else {
			res = maxInt(hr*length, res)
			right--
		}
	}

	return res
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
