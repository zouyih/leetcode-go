package main

import "fmt"

func longestConsecutive(nums []int) int {
	hash := make(map[int]int)
	res := 0
	for _, num := range nums {
		if _, ok := hash[num]; ok {
			continue
		}
		left := hash[num-1]
		right := hash[num+1]
		sum := left + right + 1
		hash[num] = sum
		hash[num-left] = sum
		hash[num+right] = sum
		res = max(res, sum)
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}
