package main

import "fmt"

func isValid(s string) bool {
	mp := map[int32]int32{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]int32, 0)

	for _, char := range s {
		if bracket, ok := mp[char]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != bracket {
				fmt.Println(string(bracket))
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	s := "{[]}"
	fmt.Println(isValid(s))
}
