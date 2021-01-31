package main

import "fmt"

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := nums[0]
	count := 0
	for i, num := range nums {
		if num == res {
			count++
		} else {
			count--
		}

		if count == 0 {
			res = nums[i+1]
		}
	}
	return res
}

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))

	nums = []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}
