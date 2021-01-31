package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func search1(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if target >= nums[left] && target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] && target <= nums[right] {
			left = mid + 1
		} else if nums[left] > nums[mid] {
			right = mid - 1
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			break
		}

	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println(search(nums, target))
	fmt.Println(search1(nums, target))
}
