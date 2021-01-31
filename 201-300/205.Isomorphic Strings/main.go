package main

import "fmt"

func isIsomorphic(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s2t := make(map[uint8]uint8)
	t2s := make(map[uint8]uint8)
	for i := range s {
		sChar := s[i]
		tChar := t[i]
		if char, ok := s2t[sChar]; ok {
			if char != tChar {
				return false
			}
		}
		if char, ok := t2s[tChar]; ok {
			if char != sChar {
				return false
			}
		}
		s2t[sChar] = tChar
		t2s[tChar] = sChar
	}
	return true
}

func main() {
	s := "egg"
	t := "add"
	fmt.Println(isIsomorphic(s, t))

	s = "foo"
	t = "bar"
	fmt.Println(isIsomorphic(s, t))

	s = "paper"
	t = "title"
	fmt.Println(isIsomorphic(s, t))
}
