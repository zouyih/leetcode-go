package main

import "fmt"

func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		res += int(num & 1)
		num >>= 1
	}
	return res
}

func main() {
	n := uint32(100)
	fmt.Printf("%b\n", n)
	fmt.Println(hammingWeight(n))
}
