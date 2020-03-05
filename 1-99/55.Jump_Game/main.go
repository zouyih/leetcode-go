package main

import (
	"fmt"
)

func canJump(nums []int) bool {
	n := len(nums)

	thisMax := 0
	for i := 0; i <= thisMax; i++ {
		if i == n-1 {
			return true
		}
		if i+nums[i] > thisMax {
			thisMax = i + nums[i]
		}
	}
	return false
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums))
}
