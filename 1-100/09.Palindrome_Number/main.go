package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	revertedNum := 0
	for x > revertedNum {
		revertedNum = revertedNum*10 + x%10
		x = x / 10
	}

	if x == revertedNum || x == revertedNum/10 {
		return true
	}

	return false
}

func main() {
	fmt.Println(isPalindrome(10))
}
