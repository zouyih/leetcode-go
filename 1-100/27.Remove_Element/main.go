package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	i := -1
	for _, num := range nums {
		if num != val {
			i++
			nums[i] = num
		}
	}
	return i + 1
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(nums)
	val := 2
	n := removeElement(nums, val)
	fmt.Println(nums[:n])
}
