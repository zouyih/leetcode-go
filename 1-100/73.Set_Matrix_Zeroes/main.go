package main

import (
	"fmt"
)

func setZeroes(matrix [][]int) {
	m := len(matrix)
	if m == 0 {
		return
	}
	n := len(matrix[0])
	row0Flag := false
	col0Flag := false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			col0Flag = true
			break
		}
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			row0Flag = true
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if row0Flag {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if col0Flag {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
	return
}

func setZeroes1(matrix [][]int) {
	m := len(matrix)
	if m == 0 {
		return
	}
	n := len(matrix[0])

	flags := make([]bool, m+n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				flags[i] = true
				flags[m+j] = true
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if flags[i] || flags[m+j] {
				matrix[i][j] = 0
			}
		}
	}

	return
}

func main() {
	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroes(matrix)
	fmt.Println(matrix)
}
