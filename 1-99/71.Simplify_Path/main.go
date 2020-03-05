package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	stack := make([]string, 0)

	items := strings.Split(path, "/")
	for _, item := range items {
		if item == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if len(item) > 0 && item != "." {
			stack = append(stack, item)
		}
	}
	res := "/"
	for i, item := range stack {
		res += item
		if i < len(stack)-1 {
			res += "/"
		}
	}
	return res
}

func main() {
	path := "/a//b////c/d//././/.."
	fmt.Println(simplifyPath(path))
}
