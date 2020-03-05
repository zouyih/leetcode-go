package main

import (
	"fmt"
)

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	directionIndex := 0
	row, col := 0, 0
	for num := 1; num <= n*n; num++ {
		matrix[row][col] = num

		direction := directions[directionIndex]
		i, j := row+direction[0], col+direction[1]
		if !isValid(i, j, matrix) {
			directionIndex = (directionIndex + 1) % 4
			direction = directions[directionIndex]
			i, j = row+direction[0], col+direction[1]
		}
		row, col = i, j
	}

	return matrix
}

func isValid(i, j int, matrix [][]int) bool {
	n := len(matrix)
	if i < 0 || i >= n || j < 0 || j >= n {
		return false
	}
	if matrix[i][j] > 0 {
		return false
	}
	return true
}

func main() {
	n := 3
	fmt.Println(generateMatrix(n))
}
