package main

import (
	"fmt"
)

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	if n == 0 {
		return []int{}
	}

	m := len(matrix[0])
	mark := make([][]bool, n)
	for i := 0; i < n; i++ {
		mark[i] = make([]bool, m)
	}

	res := make([]int, n*m)
	directionIndex := 0
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	row := 0
	col := 0
	for i := 0; i < n*m; i++ {
		res[i] = matrix[row][col]
		mark[row][col] = true

		direction := directions[directionIndex]
		i, j := row+direction[0], col+direction[1]
		if !isValid(i, j, mark, n, m) {
			directionIndex = (directionIndex + 1) % 4
			direction = directions[directionIndex]
			i, j = row+direction[0], col+direction[1]
		}
		row, col = i, j
	}
	return res
}

func isValid(i, j int, mark [][]bool, n, m int) bool {
	if i >= n || i < 0 || j >= m || j < 0 {
		return false
	}
	if mark[i][j] {
		return false
	}
	return true
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(matrix))
}
