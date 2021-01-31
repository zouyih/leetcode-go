package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	nums1 := strings.Split(version1, ".")
	nums2 := strings.Split(version2, ".")

	for len(nums1) < len(nums2) {
		nums1 = append(nums1, "0")
	}
	for len(nums2) < len(nums1) {
		nums2 = append(nums2, "0")
	}

	for i := 0; i < len(nums1); i++ {
		num1, _ := strconv.Atoi(nums1[i])
		num2, _ := strconv.Atoi(nums2[i])

		if num1 < num2 {
			return -1
		}
		if num1 > num2 {
			return 1
		}
	}
	return 0
}

func main() {
	version1 := "0.1"
	version2 := "1.1"
	fmt.Println(compareVersion(version1, version2))

	version1 = "1.0.1"
	version2 = "1"
	fmt.Println(compareVersion(version1, version2))

	version1 = "7.5.2.4"
	version2 = "7.5.3"
	fmt.Println(compareVersion(version1, version2))

}
