package main

import (
	"fmt"
)

func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		}

		if nums[left] < nums[mid] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[left] > nums[mid] {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			left++
		}
	}

	return false
}

func main() {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	fmt.Println(search(nums, 0))
	fmt.Println(search(nums, 3))
}
