package main

import (
	"fmt"
)

func rotate(matrix [][]int) {
	n := len(matrix)
	trans := make([]int, 4)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-1-i; j++ {
			row := i
			col := j
			for k := 0; k < 4; k++ {
				trans[k] = matrix[row][col]
				temp := row
				row = col
				col = n - 1 - temp
			}
			for k := 0; k < 4; k++ {
				matrix[row][col] = trans[(k+3)%4]
				temp := row
				row = col
				col = n - 1 - temp
			}
		}
	}
}

func main() {
	matrix := [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}
	rotate(matrix)
	fmt.Println(matrix)
}
