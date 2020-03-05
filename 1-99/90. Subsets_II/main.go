package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	path := make([]int, 0)
	dfs(path, 0, nums, &res)
	return res
}

func dfs(path []int, index int, nums []int, res *[][]int) {
	pathCopy := make([]int, 0, len(path)+1)
	pathCopy = append(pathCopy, path...)
	*res = append(*res, pathCopy)

	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}

		path = append(path, nums[i])
		dfs(path, i+1, nums, res)
		path = path[:len(path)-1]
	}
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}
