package main

import "fmt"

func gen(left, right int, cur string, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, cur)
		return
	}

	if left < right {
		gen(left, right-1, cur+")", res)
	}

	if left > 0 {
		gen(left-1, right, cur+"(", res)
	}
}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	gen(n, n, "", &res)
	return res
}

func main() {
	n := 3
	fmt.Println(generateParenthesis(n))
}
