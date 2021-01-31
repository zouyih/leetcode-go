package main

import "fmt"

func longestValidParentheses(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}

	res := 0
	dp := make([]int, n)
	if s[1] == ')' && s[0] == '(' {
		dp[1] = 2
		res = 2
	}

	for i := 2; i < n; i++ {
		if s[i] == '(' {
			continue
		}

		if s[i-1] == '(' {
			dp[i] = dp[i-2] + 2
		} else if i-1-dp[i-1] >= 0 && s[i-1-dp[i-1]] == '(' {
			dp[i] = dp[i-1] + 2
			if i-2-dp[i-1] >= 0 {
				dp[i] += dp[i-2-dp[i-1]]
			}
		}

		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}

func longestValidParentheses1(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}

	res := 0
	stack := make([]int, 1, n)
	stack[0] = -1

	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
			continue
		}

		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			curLen := i - stack[len(stack)-1]
			if curLen > res {
				res = curLen
			}
		}

	}

	return res
}

func longestValidParentheses2(s string) int {
	n := len(s)
	stack := make([]int, 0, n)

	for i := 0; i < n; i++ {
		switch s[i] {
		case '(':
			stack = append(stack, i)
		case ')':
			if len(stack) > 0 && s[stack[len(stack)-1]] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, i)
			}
		}
	}

	if len(stack) == 0 {
		return n
	}

	res := 0
	b := n
	for j := len(stack) - 1; j >= 0; j-- {
		a := stack[j]
		if b-a-1 > res {
			res = b - a - 1
		}
		b = a
	}

	if b > res {
		res = b
	}

	return res
}

func main() {
	s := "()(())"
	fmt.Println(longestValidParentheses(s))
	fmt.Println(longestValidParentheses1(s))
	fmt.Println(longestValidParentheses2(s))
}
