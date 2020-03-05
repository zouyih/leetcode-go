package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var res = make([]int, 2)
	var length = len(nums)
	var hash = make(map[int]int)

	for i := 0; i < length; i++ {
		if index, ok := hash[target-nums[i]]; ok {
			res[0] = index
			res[1] = i
		} else {
			hash[nums[i]] = i
		}
	}

	return res
}

func main() {
	nums := []int{1, 2, 3}
	target := 4

	fmt.Println(twoSum(nums, target))
}
