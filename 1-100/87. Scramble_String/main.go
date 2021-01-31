package main

import "fmt"

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	hash := make(map[int32]int)
	for _, char := range s1 {
		hash[char]++
	}
	for _, char := range s2 {
		hash[char]--
	}
	for _, count := range hash {
		if count != 0 {
			return false
		}
	}

	if s1 == s2 {
		return true
	}
	for i := 1; i < len(s1); i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
		if isScramble(s1[:i], s2[len(s2)-i:]) && isScramble(s1[i:], s2[:len(s2)-i]) {
			return true
		}
	}
	return false
}

func main() {
	s1 := "great"
	s2 := "rgeat"
	fmt.Println(isScramble(s1, s2))
	s1 = "abcde"
	s2 = "caebd"
	fmt.Println(isScramble(s1, s2))
}
