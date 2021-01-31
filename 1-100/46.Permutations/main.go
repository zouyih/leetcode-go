package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	dfs(nums, 0, &res)
	return res
}

func dfs(nums []int, index int, res *[][]int) {
	if index == len(nums) {
		pathCopy := make([]int, 0, len(nums))
		pathCopy = append(pathCopy, nums...)
		*res = append(*res, pathCopy)
		return
	}

	for i := index; i < len(nums); i++ {
		nums[index], nums[i] = nums[i], nums[index]
		dfs(nums, index+1, res)
		nums[index], nums[i] = nums[i], nums[index]
	}
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
