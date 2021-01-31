package main

import "fmt"

func rangeBitwiseAnd(m, n int) int {
	power := uint(0)
	for m < n {
		m >>= 1
		n >>= 1
		power++
	}
	return m << power
}

func main() {
	m, n := 5, 7
	fmt.Println(rangeBitwiseAnd(m, n))

	m, n = 0, 1
	fmt.Println(rangeBitwiseAnd(m, n))
}
