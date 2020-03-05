package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	res := 0
	buyPrice := prices[0]
	for _, price := range prices[1:] {
		if price < buyPrice {
			buyPrice = price
		}

		res = max(res, price-buyPrice)
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))

	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}
