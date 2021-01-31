package main

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
		}
	}
	return nums[left]
}

func main() {
	nums := []int{1, 3, 5}
	fmt.Println(findMin(nums))

	nums = []int{2, 2, 2, 0, 1}
	fmt.Println(findMin(nums))
}
