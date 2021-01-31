package main

import "fmt"

func countAndSay(n int) string {
	if n == 0 {
		return ""
	}

	res := "1"
	n--

	for n > 0 {
		cur := ""
		for i := 0; i < len(res); i++ {
			count := 1
			for i < len(res)-1 && res[i] == res[i+1] {
				count++
				i++
			}
			cur += string(count + '0')
			cur += string(res[i])
		}
		res = cur
		n--
	}
	return res
}

func main() {
	n := 4
	fmt.Println(countAndSay(n))
}
