package main

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	sig := 1
	if numerator < 0 {
		sig *= -1
		numerator *= -1
	}
	if denominator < 0 {
		sig *= -1
		denominator *= -1
	}

	res := ""
	if sig < 0 {
		res += "-"
	}

	a := numerator / denominator
	b := numerator % denominator

	res += strconv.Itoa(a)
	if b == 0 {
		return res
	}

	res += "."
	indexMap := make(map[int]int)

	for b != 0 {
		if index, ok := indexMap[b]; ok {
			res = res[:index] + "(" + res[index:] + ")"
			return res
		}
		indexMap[b] = len(res)

		b *= 10
		a = b / denominator
		b = b % denominator
		res += strconv.Itoa(a)
	}

	return res
}

func main() {
	fmt.Println(fractionToDecimal(1, 2))
	fmt.Println(fractionToDecimal(6, 2))
	fmt.Println(fractionToDecimal(1, 3))
	fmt.Println(fractionToDecimal(4, 333))
}
