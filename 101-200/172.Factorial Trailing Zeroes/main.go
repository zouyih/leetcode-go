package main

import "fmt"

func trailingZeroes(n int) int {
	count := 0
	for n > 0 {
		count += n / 5
		n /= 5
	}
	return count
}

func main() {
	n := 3
	fmt.Println(trailingZeroes(n))

	n = 5
	fmt.Println(trailingZeroes(n))
}
