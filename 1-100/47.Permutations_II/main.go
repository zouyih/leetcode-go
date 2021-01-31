package main

import (
	"fmt"
)

func permuteUnique(nums []int) [][]int {
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

	hasSwapped := make(map[int]bool)
	for i := index; i < len(nums); i++ {
		if hasSwapped[nums[i]] {
			continue
		}
		hasSwapped[nums[i]] = true
		nums[index], nums[i] = nums[i], nums[index]
		dfs(nums, index+1, res)
		nums[index], nums[i] = nums[i], nums[index]
	}
}

func main() {
	nums := []int{2, 2, 1, 1}
	fmt.Println(permuteUnique(nums))
}
