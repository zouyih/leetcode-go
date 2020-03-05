package main

import "fmt"

func intToRoman(num int) string {
	bases := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var count int
	var res string
	for i, base := range bases {
		if num == 0 {
			break
		}

		count = num / base
		num = num % base

		for count > 0 {
			res += romans[i]
			count--
		}
	}

	return res
}

func intToRoman1(num int) string {
	bases := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	res := ""
	i := 0
	for num > 0 {
		for num >= bases[i] {
			res += romans[i]
			num -= bases[i]
		}
		i++
	}

	return res
}

func main() {
	num := 58
	fmt.Println(intToRoman(num))
	fmt.Println(intToRoman1(num))
}
