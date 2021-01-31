package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	hash := make(map[string]int)
	res := make([]string, 0)
	if len(s) < 10 {
		return res
	}

	for i := range s[:len(s)-9] {
		hash[s[i:i+10]]++
	}

	for key, value := range hash {
		if value > 1 {
			res = append(res, key)
		}
	}
	return res
}

func main() {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	fmt.Println(findRepeatedDnaSequences(s))

	s = "AAAAAAAAAAAAA"
	fmt.Println(findRepeatedDnaSequences(s))

	s = "AAAAAAAAAAA"
	fmt.Println(findRepeatedDnaSequences(s))
}
