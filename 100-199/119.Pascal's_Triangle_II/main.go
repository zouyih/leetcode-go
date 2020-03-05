package main

import "fmt"

func getRow(rowIndex int) []int {
	res := []int{1}

	for n := 2; n <= rowIndex+1; n++ {
		row := make([]int, n)
		row[0], row[n-1] = 1, 1

		for i := 1; i < n-1; i++ {
			row[i] = res[i-1] + res[i]
		}
		fmt.Println(row)
		res = row
	}
	return res
}

func main() {
	numRows := 3
	fmt.Println(getRow(numRows))
}
