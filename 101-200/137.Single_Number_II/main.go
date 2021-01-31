package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0
	for i := 0; i < 64; i++ {
		count := 0
		for _, num := range nums {
			if ((num >> uint(i)) & 1) == 1 {
				count++
			}
		}
		if count%3 != 0 {
			res |= 1 << uint(i)
		}
	}
	return res
}

func main() {
	nums := []int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2}
	fmt.Println(singleNumber(nums))
}
