package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	path := make([]int, 0)
	dfs(candidates, 0, len(candidates)-1, path, target, &res)

	return res
}

func dfs(candidates []int, start, end int, path []int, target int, res *[][]int) {
	if target == 0 {
		pathCopy := make([]int, 0, len(path))
		pathCopy = append(pathCopy, path...)
		*res = append(*res, pathCopy)
	}

	for i := start; i <= end; i++ {
		residue := target - candidates[i]
		if residue < 0 {
			return
		}

		path = append(path, candidates[i])
		dfs(candidates, i, end, path, residue, res)
		path = path[:len(path)-1]
	}
}

func main() {
	candidates := []int{3, 5, 2}
	target := 8
	fmt.Println(combinationSum(candidates, target))
}
