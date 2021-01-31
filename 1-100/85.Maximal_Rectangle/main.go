package main

import (
	"fmt"
)

func maximalRectangle(matrix [][]byte) int {
	res := 0
	m := len(matrix)
	if m == 0 {
		return res
	}

	n := len(matrix[0])
	heights := make([]int, n)
	left := make([]int, n)
	right := make([]int, n)
	for j := 0; j < n; j++ {
		right[j] = n - 1
	}

	for i := 0; i < m; i++ {
		curLeft := 0
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
				left[j] = max(left[j], curLeft)
			} else {
				heights[j] = 0
				left[j] = 0
				curLeft = j + 1
			}
		}

		curRight := n - 1
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				right[j] = min(right[j], curRight)
			} else {
				right[j] = n - 1
				curRight = j - 1
			}
		}

		for j := 0; j < n; j++ {
			area := (right[j] - left[j] + 1) * heights[j]
			res = max(res, area)
		}
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalRectangle(matrix))
}
