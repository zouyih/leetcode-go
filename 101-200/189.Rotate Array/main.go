package main

import "fmt"

func rotate1(nums []int, k int) {
	n := len(nums)
	k = k % n

	newNums := append(nums[n-k:], nums[:n-k]...)
	for i, num := range newNums {
		nums[i] = num
	}
	return
}

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	index := 0
	for i := k; i < n || index < k; i++ {
		if i == n {
			n = k
			k = k - index
			index = 0
			i = k
		}
		if k == 0 {
			break
		}

		index = index % k
		nums[i], nums[index] = nums[index], nums[i]
		index++
	}

	return
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate1(nums, k)
	fmt.Println(nums)

	nums = []int{-1, -100, 3, 99}
	k = 2
	rotate(nums, k)
	fmt.Println(nums)
}
