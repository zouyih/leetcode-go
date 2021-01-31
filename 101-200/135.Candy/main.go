package main

import "fmt"

func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1
	}
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	sum := candies[n-1]
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
		sum += candies[i]
	}
	return sum
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	ratings := []int{1, 0, 2}
	fmt.Println(candy(ratings))
}
