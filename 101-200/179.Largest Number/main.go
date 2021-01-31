package main

import (
	"fmt"
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	strNums := make([]string, len(nums))
	for i, num := range nums {
		strNums[i] = strconv.Itoa(num)
	}

	sort.Slice(strNums, func(i, j int) bool {
		a := strNums[i] + strNums[j]
		b := strNums[j] + strNums[i]
		if a > b {
			return true
		}
		return false
	})

	res := ""
	for _, strNum := range strNums {
		res += strNum
	}
	if res[0] == '0' {
		return "0"
	}

	return res
}

func main() {
	nums := []int{10, 2}
	fmt.Println(largestNumber(nums))

	nums = []int{3, 30, 34, 5, 9}
	fmt.Println(largestNumber(nums))
}
