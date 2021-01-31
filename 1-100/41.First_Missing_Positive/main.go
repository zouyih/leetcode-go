package main

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i, num := range nums {
		if num <= 0 {
			nums[i] = n + 1
		}
	}

	for _, num := range nums {
		if num < 0 {
			num = -num
		}

		if num > n {
			continue
		}

		index := num - 1
		if nums[index] > 0 {
			nums[index] = -nums[index]
		}
	}

	for i, num := range nums {
		if num > 0 {
			return i + 1
		}
	}
	return n + 1
}

func firstMissingPositive1(nums []int) int {
	n := len(nums)
	for i, num := range nums {
		for num > 0 && num < n+1 && i != num-1 && nums[num-1] != num {
			nums[i], nums[num-1] = nums[num-1], nums[i]
			num = nums[i]
		}
	}

	for i, num := range nums {
		if num != i+1 {
			return i + 1
		}
	}

	return n + 1
}

func main() {
	nums := []int{1, 2, 0}
	fmt.Println(firstMissingPositive(nums))
	nums = []int{1, 2, 0}
	fmt.Println(firstMissingPositive1(nums))
}
