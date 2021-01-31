package main

import (
	"fmt"
)

func totalNQueens(n int) int {
	res := 0
	queen := make([][]byte, n)
	for i := 0; i < n; i++ {
		queen[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			queen[i][j] = '.'
		}
	}
	dfs(0, queen, &res)
	return res
}

func dfs(row int, queen [][]byte, res *int) {
	n := len(queen)
	if row == n {
		*res++
		return
	}
	for col := 0; col < n; col++ {
		if !isValid(queen, row, col) {
			continue
		}
		queen[row][col] = 'Q'
		dfs(row+1, queen, res)
		queen[row][col] = '.'
	}
}

func isValid(queen [][]byte, row, col int) bool {
	for i := 0; i < row; i++ {
		if queen[i][col] == 'Q' {
			return false
		}
	}

	i := row - 1
	j := col - 1
	for i >= 0 && j >= 0 {
		if queen[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}

	i = row - 1
	j = col + 1
	for i >= 0 && j < len(queen) {
		if queen[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}
	return true
}

func main() {
	n := 4
	fmt.Println(totalNQueens(n))
}
