package main

import (
	"fmt"
)

func largestRectangleArea(heights []int) int {
	stack := make([]int, 0)
	res := 0
	heights = append(heights, 0)

	for i, height := range heights {
		for len(stack) > 0 && height < heights[stack[len(stack)-1]] {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			left := -1
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}
			area := (i - 1 - left) * h
			res = max(res, area)
		}
		stack = append(stack, i)
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(heights))
}
