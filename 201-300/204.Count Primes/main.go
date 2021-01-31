package main

import "fmt"

func countPrimes(n int) int {
	isPrimes := make([]bool, n)
	for i := range isPrimes {
		isPrimes[i] = true
	}

	res := 0
	for i := 2; i < n; i++ {
		if isPrimes[i] {
			res++
			for j := 2; i*j < n; j++ {
				isPrimes[i*j] = false
			}
		}
	}
	return res
}

func countPrimes1(n int) int {
	res := 0

	for i := 2; i < n; i++ {
		if isPrime(i) {
			res++
		}
	}
	return res
}

func isPrime(num int) bool {
	if num == 2 {
		return true
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	n := 10
	fmt.Println(countPrimes(n))
}
