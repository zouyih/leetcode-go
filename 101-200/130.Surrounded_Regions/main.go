package main

import (
	"fmt"
)

func solve(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			isEdge := i == 0 || j == 0 || i == m-1 || j == n-1
			if isEdge && board[i][j] == 'O' {
				dfs(board, i, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return
	}

	if board[i][j] != 'O' {
		return
	}

	board[i][j] = '#'

	dfs(board, i-1, j)
	dfs(board, i, j-1)
	dfs(board, i+1, j)
	dfs(board, i, j+1)

}

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
	for _, row := range board {
		fmt.Printf("%s\n", row)
	}
}
