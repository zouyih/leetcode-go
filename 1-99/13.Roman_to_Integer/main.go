package main

import "fmt"

func romanToInt(s string) int {
	romanToBase := map[string]int{"M": 1000, "D": 500, "C": 100, "L": 50, "X": 10, "V": 5, "I": 1}

	var res int
	lastBase := 1000
	for _, char := range s {
		base, _ := romanToBase[string(char)]
		res += base

		if base > lastBase {
			res -= 2 * lastBase
		}
		lastBase = base
	}

	return res
}

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}
