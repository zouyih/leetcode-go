package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		switch token {
		case "*", "/", "+", "-":
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, compute(a, b, token))
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

func compute(a, b int, operate string) int {
	switch operate {
	case "*":
		return a * b
	case "/":
		return a / b
	case "+":
		return a + b
	case "-":
		return a - b
	}
	return -1
}

func main() {
	tokens := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(tokens))
}
