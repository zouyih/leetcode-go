package main

import "fmt"

func maxProduct(nums []int) int {
	n := len(nums)
	maxDp := make([]int, n)
	minDp := make([]int, n)

	maxDp[0] = nums[0]
	minDp[0] = nums[0]
	res := nums[0]

	for i := 1; i < n; i++ {
		if nums[i] > 0 {
			maxDp[i] = max(nums[i], maxDp[i-1]*nums[i])
			minDp[i] = min(nums[i], minDp[i-1]*nums[i])
		} else {
			maxDp[i] = max(nums[i], minDp[i-1]*nums[i])
			minDp[i] = min(nums[i], maxDp[i-1]*nums[i])
		}
		res = max(res, maxDp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(nums))
}
