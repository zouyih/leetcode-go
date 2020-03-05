package main

import "fmt"

func makeHash(n int) [][]int {
	hash := make([][]int, n)
	for i := 0; i < n; i++ {
		hash[i] = make([]int, n)
	}
	return hash
}

func isValidSudoku(board [][]byte) bool {
	rowHash := makeHash(9)
	columnHash := makeHash(9)
	boxHash := makeHash(9)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			n := board[i][j] - '0' - 1
			boxIndex := i/3*3 + j/3

			if rowHash[i][n] > 0 || columnHash[j][n] > 0 || boxHash[boxIndex][n] > 0 {
				return false
			}

			rowHash[i][n]++
			columnHash[j][n]++
			boxHash[boxIndex][n]++
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board))
}
