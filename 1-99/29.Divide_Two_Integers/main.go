package main

import "fmt"

func divide(dividend int, divisor int) int {
	if divisor == -1 {
		if dividend == -1<<31 {
			return 1<<31 - 1
		}
	}

	sign := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	res := 0
	times := uint(0)

	for dividend >= divisor {
		temp := dividend - divisor<<times
		if temp >= 0 {
			res += 1 << times
			times += 1
			dividend = temp
		} else {
			times -= 1
		}
	}

	if sign == 1 {
		return res
	} else {
		return -res
	}
}

func main() {
	dividend := 1
	divisor := 1
	fmt.Println(divide(dividend, divisor))
}
