package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	n := len(nums)
	res := nums[0]
	lastMax := nums[0]

	for i := 1; i < n; i++ {
		if lastMax > 0 {
			lastMax += nums[i]
		} else {
			lastMax = nums[i]
		}

		if lastMax > res {
			res = lastMax
		}
	}
	return res
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
