package main

import "fmt"

func jump(nums []int) int {
	step := 0
	thisMax := 0
	nextMax := 0
	for i := 0; i < len(nums)-1; i++ {
		nextMax = max(nextMax, i+nums[i])
		if i == thisMax {
			step++
			thisMax = nextMax
		}
	}
	return step
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main()  {
	nums := []int{2,3,1,1,4}
	fmt.Println(jump(nums))
}
