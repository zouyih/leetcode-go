package main

import "fmt"

func minSubArrayLen(s int, nums []int) int {
	res := len(nums) + 1
	i, j := 0, 0

	sum := 0
	for j < len(nums) {
		if sum < s {
			sum += nums[j]
			j++
		}
		for sum >= s {
			res = min(res, j-i)
			sum -= nums[i]
			i++
		}
	}

	if res == len(nums)+1 {
		res = 0
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	s := 7
	nums := []int{2, 3, 1, 2, 4, 3}

	fmt.Println(minSubArrayLen(s, nums))
}
