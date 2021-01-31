package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	sort.Ints(nums)

	nSum(nums, target, 4, 0, path, &res)
	return res
}

func nSum(nums []int, target int, num int, pos int, path []int, res *[][]int) {
	n := len(nums)
	if n < num || num < 2 {
		return
	}

	if num == 2 {
		left := pos
		right := n - 1
		for left < right {
			num1 := nums[left]
			num2 := nums[right]
			sum := num1 + num2
			if sum == target {
				temp := append(path, num1, num2)
				*res = append(*res, temp)

				for left < right && nums[left] == num1 {
					left++
				}

				for left < right && nums[right] == num2 {
					right--
				}
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
		return
	}

	for i := pos; i < n-2; i++ {
		path = append(path, nums[i])
		nSum(nums, target-nums[i], num-1, i+1, path, res)
		path = path[:len(path)-1]

		for i < n-2 && nums[i] == nums[i+1] {
			i++
		}
	}

}

func main() {
	nums := []int{-3, -2, -1, 0, 0, 1, 2, 3}
	target := 0
	fmt.Println(fourSum(nums, target))
}
