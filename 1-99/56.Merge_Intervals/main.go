package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)

	left, right := intervals[0][0], intervals[0][1]
	for _, interval := range intervals[1:] {
		l, r := interval[0], interval[1]
		if l > right {
			res = append(res, []int{left, right})
			left, right = l, r
		} else {
			right = max(right, r)
		}
	}
	res = append(res, []int{left, right})
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	intervals := [][]int{
		{1, 3}, {2, 6}, {8, 10}, {15, 18},
	}
	fmt.Println(merge(intervals))
}
