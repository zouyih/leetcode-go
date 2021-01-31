package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] < target {
			left++
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return []int{}
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(numbers, target))
}
