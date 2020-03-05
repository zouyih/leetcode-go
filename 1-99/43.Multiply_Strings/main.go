package main

import (
	"fmt"
)

func multiply(num1 string, num2 string) string {
	n1 := len(num1)
	n2 := len(num2)

	pos := make([]int, n1+n2)

	for i := n1 - 1; i >= 0; i-- {
		for j := n2 - 1; j >= 0; j-- {
			p1 := i + j
			p2 := i + j + 1

			d1 := int(num1[i] - '0')
			d2 := int(num2[j] - '0')
			sum := pos[p2] + d1*d2

			pos[p2] = sum % 10
			pos[p1] += sum / 10
		}
	}

	res := ""
	for _, num := range pos {
		if len(res) == 0 && num == 0 {
			continue
		}

		res += string(num + '0')
	}

	if len(res) == 0 {
		return "0"
	}
	return res
}

func main() {
	num1 := "123"
	num2 := "456"
	fmt.Println(multiply(num1, num2))
}
