package main

import (
	"fmt"
	"strconv"
	"strings"
)

func maxPoints(points [][]int) int {
	pointsCount := make(map[string]int)
	for _, point := range points {
		pointStr := fmt.Sprintf("%d*%d", point[0], point[1])
		pointsCount[pointStr]++
	}

	if len(pointsCount) == 1 {
		for _, val := range pointsCount {
			return val
		}
	}

	res := 0
	pointStrings := make([]string, 0, len(pointsCount)+1)
	for pointStr := range pointsCount {
		pointStrings = append(pointStrings, pointStr)
	}

	for i, point1 := range pointStrings {
		slopMap := make(map[string]int)
		for _, point2 := range pointStrings[i+1:] {
			slop := getSlop(parsePointStr(point1), parsePointStr(point2))
			slopMap[slop] += pointsCount[point2]
			res = max(res, slopMap[slop]+pointsCount[point1])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func parsePointStr(pointStr string) []int {
	nums := strings.Split(pointStr, "*")
	x, _ := strconv.Atoi(nums[0])
	y, _ := strconv.Atoi(nums[1])
	return []int{x, y}
}

func getSlop(a, b []int) string {
	x1, y1 := a[0], a[1]
	x2, y2 := b[0], b[1]

	dy := y2 - y1
	dx := x2 - x1
	gcdVal := gcd(dy, dx)
	dy /= gcdVal
	dx /= gcdVal

	return fmt.Sprintf("%d/%d", dx, dy)
}

func gcd(x, y int) int {
	for y != 0 {
		remainder := x % y
		x = y
		y = remainder
	}
	return x
}

func main() {
	points := [][]int{{1, 1}, {2, 2}, {3, 3}}
	fmt.Println(maxPoints(points))
}
