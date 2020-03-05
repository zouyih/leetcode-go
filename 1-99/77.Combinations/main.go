package main

import (
	"fmt"
)

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	start := 1
	dfs(path, start, n, k, &res)

	return res
}

func dfs(path []int, start, n, k int, res *[][]int) {
	if len(path) == k {
		pathCopy := make([]int, 0, len(path)+1)
		pathCopy = append(pathCopy, path...)
		*res = append(*res, pathCopy)
		return
	}

	for i := start; i <= n; i++ {
		path = append(path, i)
		if n-i+len(path) < k {
			continue
		}

		dfs(path, i+1, n, k, res)
		path = path[:len(path)-1]
	}
}

func main() {
	n := 4
	k := 2
	for _, nums := range combine(n, k) {
		fmt.Println(nums)
	}
}
