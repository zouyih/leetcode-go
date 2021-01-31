package main

import (
	"fmt"
	"math"
)

type bucket struct {
	max int
	min int
}

func maximumGap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	minNum := nums[0]
	maxNum := nums[0]
	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
		if num > maxNum {
			maxNum = num
		}
	}

	gap := (maxNum - minNum) / (n - 1)
	if (maxNum-minNum)%(n-1) != 0 {
		gap++
	}
	if gap == 0 {
		return 0
	}

	buckets := make([]bucket, n)
	for i := 0; i < n; i++ {
		buckets[i] = bucket{max: math.MinInt32, min: math.MaxInt32}
	}

	for _, num := range nums {
		index := (num - minNum) / gap
		b := buckets[index]
		if num < b.min {
			b.min = num
		}
		if num > b.max {
			b.max = num
		}
		buckets[index] = b
	}

	res := gap
	lastMax := buckets[0].max

	for _, b := range buckets[1:] {
		if b.max == math.MinInt32 {
			continue
		}
		diff := b.min - lastMax
		if diff > res {
			res = diff
		}
		lastMax = b.max
	}

	return res
}

func main() {
	nums := []int{3, 6, 12, 1}
	fmt.Println(maximumGap(nums))
}
