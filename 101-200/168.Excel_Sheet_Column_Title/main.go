package main

import "fmt"

func convertToTitle(n int) string {
	res := ""

	for n != 0 {
		remainder := n % 26
		n = n / 26

		if remainder == 0 {
			remainder = 26
			n--
		}

		res = string('A'+remainder-1) + res
	}

	return res
}

func convertToTitle1(n int) string {
	if n <= 26 {
		return string('A' + n - 1)
	}

	return convertToTitle1((n-1)/26) + convertToTitle1((n-1)%26+1)
}

func main() {
	n := 1
	fmt.Println(convertToTitle(n))

	n = 28
	fmt.Println(convertToTitle(n))

	n = 701
	fmt.Println(convertToTitle(n))
}
