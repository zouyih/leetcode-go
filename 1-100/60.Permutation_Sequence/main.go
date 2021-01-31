package main

import (
	"fmt"
)

func getPermutation(n int, k int) string {
	factorial := make([]int, n)
	factorial[0] = 1
	for i := 1; i <= n-1; i++ {
		factorial[i] = factorial[i-1] * i
	}

	used := make([]bool, n)
	res := ""

	dfs(0, n, k, used, factorial, &res)
	return res
}

func dfs(index, n, k int, used []bool, factorial []int, res *string) {
	if index == n {
		return
	}

	count := factorial[n-1-index]
	for i, isUsed := range used {
		if isUsed {
			continue
		}

		if k > count {
			k -= count
			continue
		}

		*res += string(i + 1 + '0')
		used[i] = true
		dfs(index+1, n, k, used, factorial, res)
	}
}

func getPermutation1(n int, k int) string {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	factorial := make([]int, n)
	factorial[0] = 1
	for i := 1; i <= n-1; i++ {
		factorial[i] = factorial[i-1] * i
	}

	res := ""
	k--
	for i := n - 1; i > 0; i-- {
		fac := factorial[i]
		index := k / fac
		res += string(nums[index] + '0')
		k = k % fac
		nums = append(nums[:index], nums[index+1:]...)
	}

	res += string(nums[0] + '0')
	return res
}

func main() {
	n := 4
	k := 9
	fmt.Println(getPermutation(n, k))
	fmt.Println(getPermutation1(n, k))
}
