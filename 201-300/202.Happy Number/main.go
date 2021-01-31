package main

import "fmt"

func isHappy(n int) bool {
	slow := bitSquareSum(n)
	if slow == 1 {
		return true
	}

	fast := bitSquareSum(slow)
	for slow != fast {
		slow = bitSquareSum(slow)
		fast = bitSquareSum(fast)
		fast = bitSquareSum(fast)

		if slow == 1 {
			return true
		}
	}
	return false
}

func bitSquareSum(n int) int {
	res := 0
	for n != 0 {
		remainder := n % 10
		res += remainder * remainder
		n /= 10
	}
	return res
}

func main() {
	n := 19
	fmt.Println(isHappy(n))
}
