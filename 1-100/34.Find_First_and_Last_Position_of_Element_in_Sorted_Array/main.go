package main

import "fmt"

func searchRange(nums []int, target int) []int {
	res := make([]int, 2)

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}

	if left < len(nums) && nums[left] == target {
		res[0] = left
	} else {
		res[0] = -1
		res[1] = -1
		return res
	}

	left = 0
	right = len(nums) - 1
	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}

	if right >= 0 && nums[right] == target {
		res[1] = right
	} else {
		res[1] = -1
	}

	return res
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(nums, target))
}
