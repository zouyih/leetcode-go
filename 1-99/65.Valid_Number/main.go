package main

import (
	"fmt"
)

func isNumber(s string) bool {
	transfer := [][]int{
		{0, 1, 2, 3, -1, -1},
		{-1, -1, 2, 3, -1, -1},
		{8, -1, 2, 5, 4, -1},
		{-1, -1, 5, -1, -1, -1},
		{-1, 6, 7, -1, -1, -1},
		{8, -1, 5, -1, 4, -1},
		{-1, -1, 7, -1, -1, -1},
		{8, -1, 7, -1, -1, -1},
		{8, -1, -1, -1, -1, -1},
	}
	state := 0

	for _, char := range s {
		i := char2int(char)
		state = transfer[state][i]
		if state == -1 {
			return false
		}
	}

	switch state {
	case 2, 5, 7, 8:
		return true
	default:
		return false
	}
}

func char2int(char int32) int {
	switch char {
	case ' ':
		return 0
	case '+', '-':
		return 1
	case '.':
		return 3
	case 'e':
		return 4
	}

	i := char - '0'
	if i >= 0 && i <= 9 {
		return 2
	}

	return 5
}

func main() {
	s := "3."
	fmt.Println(isNumber(s))
}
