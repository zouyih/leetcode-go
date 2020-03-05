package main

import (
	"fmt"
)

func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, visited, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, visited [][]bool, i, j, index int, word string) bool {
	if index == len(word) {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || visited[i][j] {
		return false
	}

	if board[i][j] != word[index] {
		return false
	}

	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	visited[i][j] = true
	for _, direction := range directions {
		if dfs(board, visited, i+direction[0], j+direction[1], index+1, word) {
			return true
		}
	}
	visited[i][j] = false
	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "ABCCED"))
	fmt.Println(exist(board, "SEE"))
	fmt.Println(exist(board, "ABCB"))
}
