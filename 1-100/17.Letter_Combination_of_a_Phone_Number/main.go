package main

import "fmt"

func letterCombinations(digits string) []string {
	charMap := []string{"0", "1", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	res := make([]string, 0)

	if len(digits) == 0 {
		return res
	}

	chars := charMap[digits[0]-'0']

	if len(digits) == 1 {
		for _, char := range chars {
			res = append(res, string(char))
		}
		return res
	}

	for _, char := range chars {
		for _, s := range letterCombinations(digits[1:]) {
			item := string(char) + s
			res = append(res, item)
		}
	}
	return res
}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}
