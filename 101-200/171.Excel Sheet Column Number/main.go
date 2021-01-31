package main

import "fmt"

func titleToNumber(s string) int {
	res := 0
	base := 1

	for i := len(s) - 1; i >= 0; i-- {
		num := int(s[i]) - 64
		res += num * base
		base *= 26
	}
	return res
}

func main() {
	s := "A"
	fmt.Println(titleToNumber(s))

	s = "AB"
	fmt.Println(titleToNumber(s))

	s = "ZY"
	fmt.Println(titleToNumber(s))
}
