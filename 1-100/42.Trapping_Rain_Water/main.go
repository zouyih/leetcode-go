package main

import (
	"fmt"
)

func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	leftMax := make([]int, n)
	leftMax[0] = height[0]
	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]

	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	res := 0
	for i, h := range height {
		res += max(0, min(leftMax[i], rightMax[i])-h)
	}

	return res
}

func trap1(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	left := 0
	right := n - 1

	res := 0
	leftMax := height[left]
	rightMax := height[right]
	for left <= right {
		if leftMax <= rightMax {
			if height[left] < leftMax {
				res += leftMax - height[left]
			}
			if height[left] > leftMax {
				leftMax = height[left]
			}
			left++
		}

		if leftMax > rightMax {
			if height[right] < rightMax {
				res += rightMax - height[right]
			}
			if height[right] > rightMax {
				rightMax = height[right]
			}
			right--
		}
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
	fmt.Println(trap1(height))
}
