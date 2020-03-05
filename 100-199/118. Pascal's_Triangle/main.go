package main

import "fmt"

func generate(numRows int) [][]int {
	res := make([][]int, 0)
	if numRows == 0 {
		return res
	}

	res = append(res, []int{1})

	for n := 2; n <= numRows; n++ {
		row := make([]int, n)
		row[0], row[n-1] = 1, 1

		for i := 1; i < n-1; i++ {
			row[i] = res[n-2][i-1] + res[n-2][i]
		}
		res = append(res, row)
	}
	return res
}

func main() {
	numRows := 5
	rows := generate(numRows)
	for i := 0; i < numRows; i++ {
		fmt.Println(rows[i])
	}
}
