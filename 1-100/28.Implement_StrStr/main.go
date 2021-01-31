package main

import "fmt"

func makeFail(needle string) []int {
	fail := make([]int, len(needle))

	for i := 0; i < len(fail); i++ {
		fail[i] = -1
	}

	for i := 1; i < len(needle); i++ {
		j := fail[i-1] + 1
		for j != 0 {
			if needle[i] == needle[j] {
				fail[i] = j
				break
			} else {
				j = fail[j-1] + 1
			}
		}

		if j == 0 && needle[i] == needle[0] {
			fail[i] = 0
		}
	}

	return fail
}

func strStr(haystack string, needle string) int {
	fail := makeFail(needle)

	n := len(haystack)
	m := len(needle)

	i := 0
	j := 0
	for i < n && j < m {
		if haystack[i] == needle[j] {
			i++
			j++
		} else if j != 0 {
			j = fail[j-1] + 1
		} else {
			i++
		}
	}

	if j == m {
		return i - m
	}
	return -1
}

type KMP struct {
	pattern string
	dp      [][256]int
}

func (k KMP) search(s string) int {
	if len(k.pattern) == 0 {
		return 0
	}

	state := 0
	dp := k.dp
	for i := 0; i < len(s); i++ {
		state = dp[state][s[i]]
		if state == len(dp) {
			return i - len(dp) + 1
		}
	}
	return -1
}

func newKmp(pattern string) KMP {
	m := len(pattern)
	if m == 0 {
		return KMP{pattern: pattern}
	}

	dp := make([][256]int, m)
	dp[0][pattern[0]] = 1

	failState := 0
	for i := 1; i < m; i++ {
		for char := 0; char < 256; char++ {
			if char == int(pattern[i]) {
				dp[i][char] = i + 1
			} else {
				dp[i][char] = dp[failState][char]
			}
		}
		failState = dp[failState][pattern[i]]
	}
	return KMP{pattern: pattern, dp: dp}
}

func strStr1(haystack string, needle string) int {
	kmp := newKmp(needle)
	return kmp.search(haystack)
}

func main() {
	haystack := "mississippi"
	needle := "issip"
	fmt.Println(strStr(haystack, needle))
	fmt.Println(strStr1(haystack, needle))
}
