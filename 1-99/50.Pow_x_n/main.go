package main

import (
	"fmt"
)

func fastPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	half := fastPow(x, n/2)
	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}

func myPow(x float64, n int) float64 {
	res := float64(1)
	if n < 0 {
		x = 1 / x
		n = -(n + 1)
		res = x
	}
	res *= fastPow(x, n)
	return res
}

func myPow1(x float64, n int) float64 {
	res := float64(1)
	if n < 0 {
		x = 1 / x
		n = -(n + 1)
		res = x
	}

	curProduct := x
	for n > 0 {
		if n%2 == 1 {
			res *= curProduct
		}
		curProduct *= curProduct
		n /= 2
	}
	return res
}

func main() {
	x := 2.10000
	n := 3
	fmt.Println(myPow(x, n))
	fmt.Println(myPow1(x, n))
}
