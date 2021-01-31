package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)

	n := len(nums)
	for i := 0; i < n-2; i++ {
		l := i + 1
		r := n - 1
		target := -nums[i]

		for l < r {
			if nums[l]+nums[r] < target {
				l++
			} else if nums[l]+nums[r] > target {
				r--
			} else {
				triplet := []int{nums[i], nums[l], nums[r]}
				res = append(res, triplet)
				for l < r && nums[l] == triplet[1] {
					l++
				}
				for l < r && nums[r] == triplet[2] {
					r--
				}
			}
		}

		for i < n-2 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}
func main() {
	nums := []int{-2, 0, 0, 2, 2}
	fmt.Println(threeSum(nums))
}
