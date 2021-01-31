package main

import (
	"fmt"
	"strconv"
)

func isValid(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '0' && len(s) > 1 {
		return false
	}
	num, _ := strconv.Atoi(s)
	if num > 255 {
		return false
	}
	return true
}

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	for a := 1; a <= 3; a++ {
		for b := 1; b <= 3; b++ {
			for c := 1; c <= 3; c++ {
				if a+b+c >= len(s) {
					continue
				}
				addrs := make([]string, 4)
				addrs[0] = s[:a]
				addrs[1] = s[a : a+b]
				addrs[2] = s[a+b : a+b+c]
				addrs[3] = s[a+b+c:]
				skip := false
				tmp := ""
				for i, addr := range addrs {
					if !isValid(addr) {
						skip = true
						break
					}
					tmp += addr
					if i < len(addrs)-1 {
						tmp += "."
					}
				}
				if skip {
					continue
				}
				res = append(res, tmp)
			}
		}

	}

	return res
}

func restoreIpAddresses1(s string) []string {
	res := make([]string, 0)
	path := ""
	dfs(0, 4, path, s, &res)
	return res
}

func dfs(index, count int, path, s string, res *[]string) {
	if index == len(s) && count == 0 {
		*res = append(*res, path[:len(path)-1])
		return
	}
	if count == 0 {
		return
	}

	for i := 0; i < 3; i++ {
		if index+i >= len(s) {
			break
		}
		if s[index] == '0' {
			dfs(index+1, count-1, path+"0.", s, res)
			break
		}
		num, _ := strconv.Atoi(s[index : index+i+1])
		if num > 255 {
			break
		}
		dfs(index+i+1, count-1, path+s[index:index+i+1]+".", s, res)
	}
}

func main() {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))
	fmt.Println(restoreIpAddresses1(s))
}
