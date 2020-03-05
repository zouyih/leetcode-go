package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	hash := make(map[string][]string)

	for _, str := range strs {
		key := ""

		count := make([]int, 26)
		for _, char := range str {
			count[char-'a'] += 1
		}
		for _, c := range count {
			key += string(c)
		}

		hash[key] = append(hash[key], str)

	}

	res := make([][]string, 0, len(hash))
	for _, v := range hash {
		res = append(res, v)
	}

	return res
}

func groupAnagrams1(strs []string) [][]string {
	hash := make(map[string][]string)

	for _, str := range strs {
		b := []byte(str)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		key := string(b)

		hash[key] = append(hash[key], str)
	}

	res := make([][]string, 0, len(hash))
	for _, v := range hash {
		res = append(res, v)
	}

	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
	fmt.Println(groupAnagrams1(strs))
}
