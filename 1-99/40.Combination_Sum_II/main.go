package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	path := make([]int, 0)
	dfs(candidates, 0, len(candidates)-1, path, &res, target)

	return res
}

func dfs(candidates []int, start, end int, path []int, res *[][]int, target int) {
	if target == 0 {
		pathCopy := make([]int, 0, len(path))
		pathCopy = append(pathCopy, path...)
		*res = append(*res, pathCopy)
	}

	for i := start; i <= end; i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}

		residue := target - candidates[i]
		if residue < 0 {
			return
		}

		path = append(path, candidates[i])
		dfs(candidates, i+1, end, path, res, residue)
		path = path[:len(path)-1]
	}
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}
