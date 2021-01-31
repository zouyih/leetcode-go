package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	up, down := 0, m-1
	for up <= down {
		mid := (up + down) / 2
		if matrix[mid][0] == target {
			return true
		}
		if matrix[mid][0] < target {
			up = mid + 1
		}
		if matrix[mid][0] > target {
			down = mid - 1
		}
	}

	if down < 0 {
		return false
	}

	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if matrix[down][mid] == target {
			return true
		}
		if matrix[down][mid] < target {
			left = mid + 1
		}
		if matrix[down][mid] > target {
			right = mid - 1
		}
	}

	return false
}

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	target := 3
	fmt.Println(searchMatrix(matrix, target))
	target = 13
	fmt.Println(searchMatrix(matrix, target))
}
