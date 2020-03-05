package main

import (
	"fmt"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	newStart, newEnd := newInterval[0], newInterval[1]

	res := make([][]int, 0)
	i := 0
	for i < n && intervals[i][0] < newStart {
		res = append(res, intervals[i])
		i++
	}

	if len(res) == 0 || res[len(res)-1][1] < newStart {
		res = append(res, newInterval)
	} else {
		res[len(res)-1][1] = max(res[len(res)-1][1], newEnd)
	}

	for i < n {
		if newEnd < intervals[i][0] {
			res = append(res, intervals[i])
		}
		res[len(res)-1][1] = max(newEnd, intervals[i][1])
		i++
	}

	return res
}

func insert1(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if newInterval[0] == intervals[mid][0] {
			left = mid + 1
			break
		}
		if newInterval[0] < intervals[mid][0] {
			right = mid - 1
		}
		if newInterval[0] > intervals[mid][0] {
			left = mid + 1
		}
	}

	res := make([][]int, 0)
	res = append(res, intervals[:left]...)
	newStart, newEnd := newInterval[0], newInterval[1]
	if len(res) == 0 || res[len(res)-1][1] < newStart {
		res = append(res, newInterval)
	} else {
		res[len(res)-1][1] = max(res[len(res)-1][1], newEnd)
	}

	for i := left; i < n; i++ {
		if newEnd < intervals[i][0] {
			res = append(res, intervals[i:]...)
			return res
		}
		res[len(res)-1][1] = max(newEnd, intervals[i][1])
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
	intervals := [][]int{
		{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
	}
	newInterval := []int{4, 8}
	fmt.Println(insert(intervals, newInterval))
	fmt.Println(insert1(intervals, newInterval))
}
