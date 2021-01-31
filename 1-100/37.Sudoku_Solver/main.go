package main

import "fmt"

var (
	rowHash, colHash, boxHash [][]int
)

func solveSudoku(board [][]byte) {
	n := len(board)
	rowHash = make([][]int, n)
	colHash = make([][]int, n)
	boxHash = make([][]int, n)

	for i := 0; i < n; i++ {
		rowHash[i] = make([]int, n)
		colHash[i] = make([]int, n)
		boxHash[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] != '.' {
				placeNumber(board, i, j, int(board[i][j]-'0'))
			}
		}
	}
	backTrack(board)
}

func backTrack(board [][]byte) bool {
	n := len(board)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				for num := 1; num <= 9; num++ {
					if couldPlace(i, j, num) {
						placeNumber(board, i, j, num)
						if backTrack(board) {
							return true
						}
						removeNumber(board, i, j, num)
					}
				}
				return false
			}
		}
	}
	return true
}

func couldPlace(i, j, num int) bool {
	boxIndex := i/3*3 + j/3
	return rowHash[i][num-1] == 0 && colHash[j][num-1] == 0 && boxHash[boxIndex][num-1] == 0
}

func placeNumber(board [][]byte, i, j, num int) {
	board[i][j] = byte(num + '0')

	boxIndex := i/3*3 + j/3
	rowHash[i][num-1]++
	colHash[j][num-1]++
	boxHash[boxIndex][num-1]++
}

func removeNumber(board [][]byte, i, j, num int) {
	board[i][j] = '.'

	boxIndex := i/3*3 + j/3
	rowHash[i][num-1]--
	colHash[j][num-1]--
	boxHash[boxIndex][num-1]--
}

func main() {
	board := [][]byte{
		{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
		{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
		{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
		{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
		{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
		{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
	}
	solveSudoku(board)
	for _, row := range board {
		for _, num := range row {
			fmt.Printf("%s ", string(num))
		}
		fmt.Printf("\n")
	}
}
