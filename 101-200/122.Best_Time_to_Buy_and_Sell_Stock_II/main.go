package main

import "fmt"

func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	res := 0
	start, end := 0, 0
	for end < len(prices) {
		if end == start || prices[end] > prices[end-1] {
			end++
			continue
		}
		if prices[end-1] > prices[start] {
			res += prices[end-1] - prices[start]
		}
		start = end
	}
	if prices[end-1] > prices[start] {
		res += prices[end-1] - prices[start]
	}

	return res
}

func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	res := 0
	buyPrice := prices[0]
	lastPrice := prices[0]
	profit := 0
	for _, price := range prices[1:] {
		if price < lastPrice {
			res += profit
			profit = 0
			buyPrice = -1
		}
		if price < buyPrice || buyPrice == -1 {
			buyPrice = price
		}
		profit = max(profit, price-buyPrice)
		lastPrice = price
	}
	res += profit
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

	prices = []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(prices))
}
