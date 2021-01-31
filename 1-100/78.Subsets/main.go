package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)
	dfs(path, nums, 0, &res)
	return res
}

func dfs(path []int, nums []int, start int, res *[][]int) {
	pathCopy := make([]int, 0, len(path)+1)
	pathCopy = append(pathCopy, path...)
	*res = append(*res, pathCopy)

	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		dfs(path, nums, i+1, res)
		path = path[:len(path)-1]
	}
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
