package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	index := 0
	for _, num := range nums[1:] {
		if num != nums[index] {
			index++
			nums[index] = num
		}
	}
	return index + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(nums)
	n := removeDuplicates(nums)
	fmt.Println(nums[:n])
}
