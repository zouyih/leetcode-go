package main

import "fmt"

func nextPermutation(nums []int) {
	n := len(nums)

	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := n - 1
		for j > i && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	left := i + 1
	right := n - 1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return
}

func main() {
	nums := []int{1, 5, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
