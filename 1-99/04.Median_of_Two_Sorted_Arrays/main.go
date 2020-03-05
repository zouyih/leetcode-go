package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		len1, len2     int
		c1, c2         int
		lenMin, lenMax int
		k              int
	)

	len1 = len(nums1)
	len2 = len(nums2)

	if len1 > len2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	lenMin = 0
	lenMax = len1

	k = (len1 + len2 + 1) / 2
	for lenMin <= lenMax {
		c1 = (lenMin + lenMax) / 2
		c2 = k - c1

		if c1 > 0 && nums1[c1-1] > nums2[c2] {
			lenMax = c1 - 1
		} else if c1 < len1 && nums1[c1] < nums2[c2-1] {
			lenMin = c1 + 1
		} else {
			break
		}
	}

	var maxLeft float64
	if c1 == 0 {
		maxLeft = float64(nums2[c2-1])
	} else if c2 == 0 {
		maxLeft = float64(nums1[c1-1])
	} else {
		maxLeft = math.Max(float64(nums1[c1-1]), float64(nums2[c2-1]))
	}

	if (len1+len2)%2 == 1 {
		return maxLeft
	}

	var minRight float64
	if c1 == len1 {
		minRight = float64(nums2[c2])
	} else if c2 == len2 {
		minRight = float64(nums1[c1])
	} else {
		minRight = math.Min(float64(nums1[c1]), float64(nums2[c2]))
	}
	return (maxLeft + minRight) / 2
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
