package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	carry := 0
	res := ""

	for i >= 0 || j >= 0 || carry > 0 {
		num1 := 0
		if i >= 0 {
			num1 = int(a[i] - '0')
		}
		num2 := 0
		if j >= 0 {
			num2 = int(b[j] - '0')
		}
		sum := num1 + num2 + carry
		res = string(sum%2+'0') + res
		carry = sum / 2

		i--
		j--
	}

	return res
}

func main() {
	a := "1010"
	b := "1011"
	fmt.Println(addBinary(a, b))
}
