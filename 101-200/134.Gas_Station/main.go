package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	cur := 0
	total := 0
	index := 0
	for i := range gas {
		change := gas[i] - cost[i]
		cur += change
		total += change
		if cur < 0 {
			cur = 0
			index = i + 1
		}
	}
	if total < 0 {
		return -1
	}
	return index
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
}
