package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 1
	lastNum := nums[0]
	count := 1
	for cur := 1; cur < len(nums); cur++ {
		if nums[cur] == lastNum {
			count++
		} else {
			lastNum = nums[cur]
			count = 1
		}

		if count <= 2 {
			nums[i] = nums[cur]
			i++
		}
	}
	return i
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	n := removeDuplicates(nums)
	fmt.Println(nums[:n])
}
