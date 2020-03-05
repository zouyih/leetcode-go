package main

import "fmt"

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	res := grayCode(n - 1)
	for i := len(res) - 1; i >= 0; i-- {
		num := res[i]
		res = append(res, num|(1<<(n-1)))
	}
	return res
}

func grayCode1(n int) []int {
	if n == 0 {
		return []int{0}
	}

	res := []int{0}
	for i := 0; i < n; i++ {
		for j := len(res) - 1; j >= 0; j-- {
			res[j] = res[j] << 1
			res = append(res, res[j]+1)
		}
	}
	return res
}

func main() {
	n := 2
	fmt.Println(grayCode(n))
	fmt.Println(grayCode1(n))
}
