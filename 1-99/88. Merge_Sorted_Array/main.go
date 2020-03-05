package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, cur := m-1, n-1, m+n-1

	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[cur] = nums1[i]
			i--
		} else {
			nums1[cur] = nums2[j]
			j--
		}
		cur--
	}

	for j >= 0 {
		nums1[cur] = nums2[j]
		j--
		cur--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)

	fmt.Println(nums1)
}
