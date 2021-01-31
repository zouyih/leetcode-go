package main

import (
	"fmt"
	"sort"
)

func abs(num int) int {
	if num > 0 {
		return num
	} else {
		return -num
	}
}

func threeSumClosest(nums []int, target int) int {
	res := nums[0] + nums[1] + nums[2]
	n := len(nums)
	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		left := i + 1
		right := n - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == target {
				return target
			}

			if abs(sum-target) < abs(res-target) {
				res = sum
			}

			if sum > target {
				right--
			}
			if sum < target {
				left++
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1, 1, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}
