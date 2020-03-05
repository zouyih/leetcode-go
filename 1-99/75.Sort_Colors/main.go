package main

import (
	"fmt"
)

func sortColors(nums []int) {
	n := len(nums)
	left, cur, right := 0, 0, n-1

	for cur <= right {
		switch nums[cur] {
		case 2:
			nums[cur], nums[right] = nums[right], nums[cur]
			right--
		case 0:
			nums[cur], nums[left] = nums[left], nums[cur]
			cur++
			left++
		case 1:
			cur++
		}
	}
}

func sortColors1(nums []int) {
	hash := make([]int, 3)
	for _, num := range nums {
		hash[num]++
	}

	index := 0
	for i, count := range hash {
		for count > 0 {
			nums[index] = i
			count--
			index++
		}
	}
}

func main() {
	nums := []int{2, 0, 1}
	sortColors(nums)
	fmt.Println(nums)
}
